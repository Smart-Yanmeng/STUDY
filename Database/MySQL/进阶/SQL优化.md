# SQL优化

> 插入数据

+ insert优化

  1. 批量插入

     ```mysql
     insert into tb_test values(1,'Tom'),(2, 'Cat'),(3, 'Jerry');
     ```

  2. 手动提交事务

     ```mysql
     start transaction;
     # 命令段
     commit;
     ```

  3. 主键顺序插入

  4. 大批量插入数据
     如果一次性需要插入大批量数据，使用insert语句插入性能较低，此时可以使用MySQL数据库提供的load指令进行插入。操作如下：

     ```mysql
     # 客户端连接服务端时，加上参数 --local-infile
     mysql --local-infile -u root -p
     
     # 设置全局参数local_infile为1，开启从本地加载文件导入数据的开关
     set global local_infile=1;
     
     # 执行load指令将准备好的数据，加载到表结构中
     load data local infile '/root/sql1.log' into table 'tb_user' fields terminated by ',' lines terminated by '\n';
     ```

     主键顺序插入性能高于乱序插入

+ 主键优化
  1. 数据组织方式
     在InnoDB存储引擎中，表数据都是根据主键顺序组织放的，这种存储方式称为索引组织表（index organized table IOT）
     逻辑存储结构：
     表空间Tablespace - 段Segment - 区Extent(1M) - 页Page(16K) - 行Row
  2. 页分裂
     页可以为空，也可以填充一半，也可以填充100%。每个页包含了2-N行数据（如果一行数据多大，会行溢出），根据主键排列
  3. 页合并
     当删除一行记录时，实际上记录并没有被物理删除，只是记录被标记（flaged）为删除并且它的空间变得允许被其他记录声明使用。当页中删除的记录达到MERGE_THRESHOLD（默认为页的50%），InnoDB会开始寻找最靠近的页（前或后）看看是否可以将两个页合并以优化空间使用
     （MERGE_THRESHOLD：合并页的阈值，可以自己设置，在创建表或者创建索引时指定）
  4. 主键设计原则
     + 满足业务需求的情况下，尽量降低主键的长度
     + 插入数据时，尽量选择顺序插入，选择使用AUTO_INCREMENT自增主键
     + 尽量不要使用UUID做主键或者是其他自然主键，如身份证号
     + 业务操作时，避免对主键的修改

+ order by优化

  + Using filesort：通过表的索引或全表扫描，读取满足条件的数据行，然后在排序缓冲区sort buffer中完成排序操作，所有不是通过索引直接返回排序结果的排序都叫FileSort排序

  + Using index：通过有序索引顺序扫描直接返回有序数据，这种情况即为using index，不需要额外排序，效率更高

  1. 根据排序字段建立合适的索引，多字段排序时，也遵循最左前缀法则
  2. 尽量使用覆盖索引
  3. 多字段排序，一个升序一个降序，此时需要注意联合索引在创建时的规则（ASC/DESC）
  4. 如果不可避免的出现filesort，大数据量排序时，可以适当增大排序缓冲区大小sort_buffer_size（默认256k）

+ group by优化
  1. 在分组操作时，可以通过索引来提高效率
  2. 分组操作时，索引的使用也是满足最左前缀法则的

+ limit优化
  一个常见又非常头疼的问题就是limit 2000000, 10，此时需要MySQL排序前2000010记录，仅仅返回2000000 - 2000010的记录，其他记录丢弃，查询排序的代价非常大
  优化思路：一般分页查询时，通过创建覆盖索引能够比较好地提高性能，可以通过覆盖索引加子查询形式进行优化

  ```mysql
  explain select * from tb_sku t, (select id from tb_sku order by id limit 2000000, 10) a where t.id = a.id;
  ```

+ count优化

  + MyISAM引擎把一个表的总行数存在了磁盘上，因此执行count(*)的时候会直接返回这个数，效率很高

  + InnoDB引擎就麻烦了，它执行count(*)的时候，需要把数据一行一行地从引擎里面读出来，然后累积计数
    优化思路：自己计数

  + count()是一个聚合函数，对于返回的结果集，一行行地判断，如果count函数的参数不是NULL，累计值就加1，否则不加，最后返回累计值

  + count的几种用法：

    1. count(主键)
       InnoDB引擎会遍历整张表，把每一行的主键id值都取出来，返回给服务层。服务层拿到主键后，直接按行进行累加（主键不可能为null）

    2. count(字段)
       没有not null约束：InnoDB引擎会遍历整张表把每一行的字段值都取出来，返回给服务层，服务层判断是否为null，不为null，计数累加
       有not null约束：InnoDB引擎会遍历整张表把每一行的字段值都取出来，返回给服务层，直接按行进行累加

    3. count(1)
       InnoDB引擎遍历整张表，但不取值。服务层对于返回的每一行，放一个数字“1”进去，直接按行进行累加

    4. count(*)
       InnoDB引擎并不会把全部字段取出来，而是专门做了优化，不取值，服务层直接按行进行累加

       按照效率排序的话，count(字段) < count(主键id) < counnt(1) ≈ count( * )，所以尽量使用count( * )

+ update优化
  InnoDB的行锁是针对索引加的锁，不是针对记录加的锁，并且该索引不能失效，否则会从行锁升级为表锁，并发性能就会降低