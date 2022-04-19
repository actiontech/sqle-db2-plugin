[English Doc](./README.md)

## SQLE DB2 审核插件

这是 [SQLE](https:github.comactiontechsqle) 的 DB2 SQL 审核插件，SQLE是一个 SQL 审计平台。

### 1.添加规则

您可以在对应文件中定义规则，然后通过 AddRule()
方法将其添加到插件中。更多详情请参考[文档](https:actiontech.github.iosqle-docs-cn3.modules3.7_auditpluginauditplugin_development.html)。

### 2. 编译插件

执行 ```make docker_install``` 进行编译

二进制文件位于 bin 文件夹中，您可以使用它来审核您的数据库。

### 3. 准备插件运行环境

1. 下载驱动 : https://public.dhe.ibm.com/ibmdl/export/pub/software/data/db2/drivers/odbc_cli/linuxx64_odbc_cli.tar.gz
2. 将驱动拷贝到SQLE部署的机器并解压, 解压后的目录如下
   ![](./doc-image/2.png)
3. 执行命令为sqled添加环境变量, **${YOUR_DRIVER_LIB_PATH}需要替换成解压后的驱动目录下的lib目录**

   ```systemctl set-environment "LD_LIBRARY_PATH=${YOUR_DRIVER_LIB_PATH}"```
   ![](./doc-image/3.png)

4. 重新启动sqled ```systemctl restart sqled```
   ![img.png](doc-image/4.png)
   通过日志可以看到, db2插件开始工作

### 4.开始使用

更多详情请参考[docs](https:actiontech.github.iosqle-docs-cn3.modules3.7_auditpluginauditplugin_management.html)。