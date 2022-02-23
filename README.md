## SQLE DB2 Plugin

This is A DB2 SQL audit for  [SQLE](https://github.com/actiontech/sqle), which is an SQL audit platform.

### 1. Add Rule

You can define a rule in coresponding file, and then add it to the plugin by `AddRule()` method. More details please
refer to [docs](https://actiontech.github.io/sqle-docs-cn/3.modules/3.7_auditplugin/auditplugin_development.html).

### 2. Build Binary

#### Compile manually

First, you need to install go git tar libpam on your machine. You can execute ```yum install go git tar libpam```

Secondly, you need to pull the go driver of db2, because we need to download db2_cli with the help of the script in the
project. You can execute ```go get -d github.com/ibmdb/go_ibm_db@v0.4.1```

Third, you need to enter the installer directory where go_ibm_db is located, usually
in `$GOPATH/src/github.com/ibmdb/go_ibm_db@v0.4.1/installer`
or `$GOPATH/pkg/mod/github.com/ibmdb/go_ibm_db@v0.4.1/installer`

Fourth, you need to execute in the directory of the third step ```go run setup.go```

Fifth, you need to configure the downloaded driver configuration into the environment variable. You need to configure a
few items. Note that the directory needs to be adjusted according to the actual situation.

```
export DB2HOME=$GOPATH/pkg/mod/github.com/ibmdb/clidriver
export CGO_CFLAGS=-I$DB2HOME/include
export CGO_LDFLAGS=-L$DB2HOME/lib
export LD_LIBRARY_PATH=$DB2HOME/lib
```

Sixth, execute ```make install``` in the sqle-db2-plugin directory, the sqle-db2-plugin file will be generated in the
bin directory of the sqle-db2-plugin directory, this file is the plug-in

Seventh, perform the first four steps on the machine deployed by SQLe and configure the `LD_LIBRARY_PATH` environment
variable in the fifth step

#### Docker compile

First, execute the following command to generate the plugin

```bash
make docker_install
```

Secondly, Configure your plugin environment as described in step 7 of manual compilation

#### Note: Because go_ibm_db connects to the database through the cli of db2, you need to prepare the cli on the SQLE deployment machine
#### The plugin location is in the .bin directory


### 3. Begin to use

More details please refer
to [docs](https://actiontech.github.io/sqle-docs-cn/3.modules/3.7_auditplugin/auditplugin_management.html).