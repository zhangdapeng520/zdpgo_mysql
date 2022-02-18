# zdpgo_mysql
Golang操作MySQL的快捷工具库

项目地址：https://github.com/zhangdapeng520/zdpgo_mysql

## 版本历史
- 2022年2月7日：版本1.4.0
- 2022年2月9日：版本1.5.0
- 2022年2月18日：版本1.6.0 相关bug修复，日志库升级
- 2022年2月18日：版本1.7.0 增加事务相关方法

## 功能列表
核心方法
- New 创建mysql实例
- Close 关闭数据库连接
- Execute 执行SQL语句
- DeleteTable 删除表格
- QueryRow 查询单条数据
- Query 查询多条数据
- TransRowsToMaps 将rows转换为map

常用的增删改查方法
- AddMany 批量添加数据
- Add 添加数据
- DeleteById 根据ID删除数据
- DeleteByIds 根据ID列表删除
- UpdateByIds 根据ID列表修改数据
- FindById 查询单条数据
- FindByIdToStruct 根据id查询数据，并将结果转换为结构体
- FindByIds 根据ID列表查询数据
- FindByIdsToStruct 根据ID列表查询数据并映射到结构体
- FindByPage 分页查询数据
- FindByPageToStruct 执行分页查询并映射到结构体

事务方法
- Begin 开始事务
- Rollback 回滚事务
- Commit 提交事务
