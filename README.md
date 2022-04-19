[中文文档](./README-CN.md)

## SQLE DB2 Plugin

This is A DB2 SQL audit for  [SQLE](https://github.com/actiontech/sqle), which is an SQL audit platform.

### 1. Add Rule

You can define a rule in coresponding file, and then add it to the plugin by `AddRule()` method. More details please
refer to [docs](https://actiontech.github.io/sqle-docs-cn/3.modules/3.7_auditplugin/auditplugin_development.html).

### 2. Build Binary

#### Docker compile

execute ```make docker_install```

The binary file is located in bin/ folder, and you can use it to audit your database.

### 3. Prepare the plugin runtime environment

1. Download the driver on the machine deployed by SQLE
   ```bash
   curl -L https://public.dhe.ibm.com/ibmdl/export/pub/software/data/db2/drivers/odbc_cli/linuxx64_odbc_cli.tar.gz -o /opt/sqle/linuxx64_odbc_cli.tar.gz
   ```
2. Unzip the driver
   ```bash
   tar zxvf /opt/sqle/linuxx64_odbc_cli.tar.gz -C /opt/sqle/
   ```
3. Add environment variables for sqled
   ``` bash
   systemctl set-environment "LD_LIBRARY_PATH=/opt/sqle/clidriver/lib"
   ```
4. Restart sqled
   ``` bash
   systemctl restart sqled
   ```

### 4. Begin to use

More details please refer
to [docs](https://actiontech.github.io/sqle-docs-cn/3.modules/3.7_auditplugin/auditplugin_management.html).


