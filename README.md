## SQLE DB2 Plugin

This is A DB2 SQL audit for  [SQLE](https://github.com/actiontech/sqle), which is an SQL audit platform.

### 1. Add Rule

You can define a rule in coresponding file, and then add it to the plugin by `AddRule()` method. More details please
refer to [docs](https://actiontech.github.io/sqle-docs-cn/3.modules/3.7_auditplugin/auditplugin_development.html).

### 2. Build Binary

```bash
make DATABASE_NAME=${the name of the database you want to connect to with the plugin} docker_install
```

#### The plugin location is in the .bin directory


### 3. Begin to use

1. Prepare the plug-in running environment: [插件部署前环境准备](https://github.com/actiontech/sqle/discussions/360)
2. Using plug-in: [docs](https://actiontech.github.io/sqle-docs-cn/3.modules/3.7_auditplugin/auditplugin_management.html).
