module github.com/actiontech/sqle-db2-plugin

go 1.19

require (
	github.com/actiontech/sqle v1.2210.1-0.20260507092346-4ce0c2946163
	github.com/ibmdb/go_ibm_db v0.4.1
)

replace (
	github.com/actiontech/sqle => ../sqle-ee
	github.com/labstack/echo/v4 => github.com/labstack/echo/v4 v4.6.1
	github.com/pingcap/log => github.com/pingcap/log v0.0.0-20191012051959-b742a5d432e9
	github.com/pingcap/parser => github.com/sjjian/parser v0.0.0-20240704052347-b6199b7bccae
	google.golang.org/grpc => google.golang.org/grpc v1.29.0
)

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/actiontech/dms v0.0.0-20251027081421-309bc24335ca // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/denisenkom/go-mssqldb v0.9.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/go-hclog v0.14.1 // indirect
	github.com/hashicorp/go-plugin v1.4.2 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20180604194846-3520598351bb // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/labstack/echo/v4 v4.10.2 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/go-testing-interface v1.14.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.4.0 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pingcap/errors v0.11.5-0.20201126102027-b0a155152ca3 // indirect
	github.com/pingcap/log v0.0.0-20210317133921-96f4fcab92a4 // indirect
	github.com/pingcap/parser v3.0.12+incompatible // indirect
	github.com/pingcap/tidb v1.1.0-beta.0.20200630082100-328b6d0a955c // indirect
	github.com/pingcap/tipb v0.0.0-20200522051215-f31a15d98fce // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20190728182440-6a916e37a237 // indirect
	github.com/shirou/gopsutil v3.21.10+incompatible // indirect
	github.com/sijms/go-ora/v2 v2.2.15 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4 // indirect
	google.golang.org/grpc v1.50.1 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/gorm v1.24.3 // indirect
	vitess.io/vitess v0.12.0 // indirect
)
