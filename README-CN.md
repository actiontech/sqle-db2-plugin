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
3. 将sqle-db2-plugin插件拷贝到SQLE的插件目录下
   ![](./doc-image/3.png)
4. 使用ldd检查sqle-db2-plugin的依赖, 应当如下图所示
   ![](./doc-image/4.png)
5. 将第二步驱动目录中的 lib文件夹内的所有内容拷贝到 ldd显示的sqle-db2-plugin 引用的动态链接库 所在目录, 演示环境中动态库在/lib64中, 则拷贝到/lib64中,
   **这一项的目标位置需要根据实际情况进行调整**
   ![](./doc-image/5.png)
6. 再次使用ldd检查sqle-db2-plugin运行环境是否配置正确, 确保没有找不到的依赖
   ![](./doc-image/6.png)

### 4.开始使用

更多详情请参考[docs](https:actiontech.github.iosqle-docs-cn3.modules3.7_auditpluginauditplugin_management.html)。