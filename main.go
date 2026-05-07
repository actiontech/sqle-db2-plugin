package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"strings"

	"github.com/actiontech/dms/pkg/dms-common/i18nPkg"
	driverV2 "github.com/actiontech/sqle/sqle/driver/v2"
	driver "github.com/actiontech/sqle/sqle/pkg/driver"
	"github.com/actiontech/sqle/sqle/pkg/params"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/actiontech/sqle-db2-plugin/plocale"
)

// DB2Dialector implements the driver.Dialector interface for IBM DB2.
type DB2Dialector struct {
	driver.BaseDialector
	databaseName string
}

func (d *DB2Dialector) String() string {
	return "DB2"
}

func (d *DB2Dialector) DatabaseAdditionalParam() params.Params {
	return params.Params{
		&params.Param{
			Key:   "database_name",
			Value: "",
			Desc:  "数据库名",
			Type:  params.ParamTypeString,
		},
	}
}

func (d *DB2Dialector) Open(dsn *driverV2.DSN) (*sql.DB, *sql.Conn, error) {
	dbName := dsn.AdditionalParams.GetParam("database_name").String()
	d.databaseName = dbName
	con := fmt.Sprintf("HOSTNAME=%v;DATABASE=%v;PORT=%v;UID=%v;PWD=%v",
		dsn.Host, dbName, dsn.Port, dsn.User, dsn.Password)
	return d.GetConn("go_ibm_db", con)
}

func (d *DB2Dialector) ShowDatabaseSQL() string {
	return fmt.Sprintf("select '%v' as name from SYSIBM.SYSDUMMY1", d.databaseName)
}

// genI18nRuleInfo generates I18nRuleInfo for all supported languages from the Bundle.
func genI18nRuleInfo(bundle *i18nPkg.Bundle, desc, annotation, category *i18n.Message) driverV2.I18nRuleInfo {
	result := make(driverV2.I18nRuleInfo, len(bundle.LanguageTags()))
	for _, langTag := range bundle.LanguageTags() {
		result[langTag] = &driverV2.RuleInfo{
			Desc:       bundle.LocalizeMsgByLang(langTag, desc),
			Annotation: bundle.LocalizeMsgByLang(langTag, annotation),
			Category:   bundle.LocalizeMsgByLang(langTag, category),
		}
	}
	return result
}

var version string
var printVersion = flag.Bool("version", false, "Print version & exit")

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		return
	}

	b := driver.NewDriverBuilder(&DB2Dialector{})

	// Define rule with i18n support
	ruleDQL1 := &driverV2.Rule{
		Name:         "db2_dql_1",
		Level:        driverV2.RuleLevelError,
		I18nRuleInfo: genI18nRuleInfo(plocale.Bundle, plocale.DB2DQL1Desc, nil, plocale.DB2RuleCategorySQLConvention),
	}

	// Define audit handler returning I18nStr
	ruleDQL1Handler := func(ctx context.Context, rule *driverV2.Rule, rawSQL string, nextSQL []string) (i18nPkg.I18nStr, error) {
		lowerSQL := strings.ToLower(rawSQL)
		if strings.Contains(lowerSQL, "select") &&
			!strings.Contains(lowerSQL, "where") {
			return plocale.Bundle.LocalizeAll(plocale.DB2DQL1Message), nil
		}
		return nil, nil
	}

	b.AddRule(ruleDQL1, ruleDQL1Handler)
	b.SetEnableOptionalModule(
		driverV2.OptionalModuleQuery,
		driverV2.OptionalModuleI18n,
	)
	b.Serve(driver.NewDriverImpl)
}
