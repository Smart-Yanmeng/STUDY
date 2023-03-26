DCL - 数据控制语言，用来管理数据库用户、控制数据库的访问权限

> 管理用户

1. 查询用户
   ```mysql
   USE mysql;
   SELECT * FROM user
   ```

2. 创建用户
   ```mysql
   CREATE USER '用户名'@'主机名' IDENTIFIED BY '密码';
   ```

3. 修改用户密码
   ```mysql
   ALTER USER '用户名'@'主机名' IDENTIFIED WITH mysql_native_password BY '新密码'
   ```

4. 删除用户
   ```mysql
   DROP USER '用户名'@'主机名';
   ```

> 常用权限

ALL, ALL PRIVILEGES - 所有权限

SELECT - 查询数据

INSERT - 插入数据

UPDATE - 修改数据

DELETE - 删除数据

ALTER - 修改表

DROP - 删除数据库/表/视图

CREATE - 创建数据库/表

> 权限控制

1. 查询权限
   ```mysql
   SHOW GRANTS FOR '用户名'@'主机名';
   ```

2. 授予权限
   ```mysql
   GRANT 权限列表 ON 数据库名.表名 TO '用户名'@'主机名';
   ```

3. 撤销权限
   ```mysql
   REVOKE 权限列表 ON 数据库名.表名 FROM '用户名'@'主机名';
   ```

   