//go:build test_without_cgo
// +build test_without_cgo

package main

import (
	"context"
	"strings"
	"testing"

	"github.com/actiontech/dms/pkg/dms-common/i18nPkg"
	driverV2 "github.com/actiontech/sqle/sqle/driver/v2"
	"golang.org/x/text/language"

	"github.com/actiontech/sqle-db2-plugin/plocale"
)

func TestGenI18nRuleInfo(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T, result driverV2.I18nRuleInfo)
	}{
		{
			name: "contains Chinese RuleInfo",
			check: func(t *testing.T, result driverV2.I18nRuleInfo) {
				info, ok := result[language.Chinese]
				if !ok || info == nil {
					t.Fatal("expected Chinese RuleInfo to exist")
				}
				if info.Desc == "" {
					t.Error("expected Chinese Desc to be non-empty")
				}
			},
		},
		{
			name: "contains English RuleInfo",
			check: func(t *testing.T, result driverV2.I18nRuleInfo) {
				info, ok := result[language.English]
				if !ok || info == nil {
					t.Fatal("expected English RuleInfo to exist")
				}
				if info.Desc == "" {
					t.Error("expected English Desc to be non-empty")
				}
			},
		},
		{
			name: "annotation is empty when nil",
			check: func(t *testing.T, result driverV2.I18nRuleInfo) {
				info := result[language.Chinese]
				if info.Annotation != "" {
					t.Errorf("expected Annotation to be empty, got %q", info.Annotation)
				}
			},
		},
		{
			name: "language tag count matches bundle",
			check: func(t *testing.T, result driverV2.I18nRuleInfo) {
				expected := len(plocale.Bundle.LanguageTags())
				if len(result) != expected {
					t.Errorf("expected %d language tags, got %d", expected, len(result))
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := genI18nRuleInfo(plocale.Bundle, plocale.DB2DQL1Desc, nil, plocale.DB2RuleCategorySQLConvention)
			tc.check(t, result)
		})
	}
}

// createDQL1Handler creates the audit handler for db2_dql_1 rule.
// This replicates the handler logic from main() for testability.
func createDQL1Handler() func(ctx context.Context, rule *driverV2.Rule, rawSQL string, nextSQL []string) (i18nPkg.I18nStr, error) {
	return func(ctx context.Context, rule *driverV2.Rule, rawSQL string, nextSQL []string) (i18nPkg.I18nStr, error) {
		lowerSQL := strings.ToLower(rawSQL)
		if strings.Contains(lowerSQL, "select") &&
			!strings.Contains(lowerSQL, "where") {
			return plocale.Bundle.LocalizeAll(plocale.DB2DQL1Message), nil
		}
		return nil, nil
	}
}

func TestRuleDQL1Handler(t *testing.T) {
	handler := createDQL1Handler()
	rule := &driverV2.Rule{
		Name:  "db2_dql_1",
		Level: driverV2.RuleLevelError,
	}

	cases := []struct {
		name        string
		sql         string
		expectNil   bool
		checkResult func(t *testing.T, result i18nPkg.I18nStr)
	}{
		{
			name:      "triggers on SELECT without WHERE",
			sql:       "SELECT * FROM t",
			expectNil: false,
			checkResult: func(t *testing.T, result i18nPkg.I18nStr) {
				if result == nil {
					t.Fatal("expected non-nil result")
				}
				if _, ok := result[language.Chinese]; !ok {
					t.Error("expected Chinese message in result")
				}
				if _, ok := result[language.English]; !ok {
					t.Error("expected English message in result")
				}
			},
		},
		{
			name:      "does not trigger on SELECT with WHERE",
			sql:       "SELECT * FROM t WHERE id=1",
			expectNil: true,
		},
		{
			name:      "does not trigger on non-SELECT",
			sql:       "INSERT INTO t VALUES(1)",
			expectNil: true,
		},
		{
			name:      "checks DefaultLang Chinese in result",
			sql:       "SELECT * FROM t",
			expectNil: false,
			checkResult: func(t *testing.T, result i18nPkg.I18nStr) {
				zhMsg, ok := result[language.Chinese]
				if !ok || zhMsg == "" {
					t.Error("expected non-empty Chinese message")
				}
			},
		},
		{
			name:      "case insensitive lowercase select",
			sql:       "select * from t",
			expectNil: false,
			checkResult: func(t *testing.T, result i18nPkg.I18nStr) {
				if result == nil {
					t.Fatal("expected non-nil result for lowercase select")
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := handler(context.Background(), rule, tc.sql, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.expectNil {
				if result != nil {
					t.Errorf("expected nil result, got %v", result)
				}
			} else if tc.checkResult != nil {
				tc.checkResult(t, result)
			}
		})
	}
}

func TestDB2Dialector(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T, d *DB2Dialector)
	}{
		{
			name: "String returns DB2",
			check: func(t *testing.T, d *DB2Dialector) {
				if got := d.String(); got != "DB2" {
					t.Errorf("expected %q, got %q", "DB2", got)
				}
			},
		},
		{
			name: "DatabaseAdditionalParam contains database_name",
			check: func(t *testing.T, d *DB2Dialector) {
				ps := d.DatabaseAdditionalParam()
				found := false
				for _, p := range ps {
					if p.Key == "database_name" {
						found = true
						break
					}
				}
				if !found {
					t.Error("expected DatabaseAdditionalParam to contain 'database_name'")
				}
			},
		},
		{
			name: "ShowDatabaseSQL contains SYSIBM.SYSDUMMY1",
			check: func(t *testing.T, d *DB2Dialector) {
				sql := d.ShowDatabaseSQL()
				if !strings.Contains(sql, "SYSIBM.SYSDUMMY1") {
					t.Errorf("expected ShowDatabaseSQL to contain 'SYSIBM.SYSDUMMY1', got %q", sql)
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := &DB2Dialector{}
			tc.check(t, d)
		})
	}
}
