package plocale

import (
	"embed"

	"github.com/actiontech/dms/pkg/dms-common/i18nPkg"
)

//go:embed active.*.toml
var localeFS embed.FS

var Bundle *i18nPkg.Bundle

func init() {
	b, err := i18nPkg.NewBundleFromTomlDir(localeFS, &i18nPkg.StdLogger{})
	if err != nil {
		panic(err)
	}
	Bundle = b
}
