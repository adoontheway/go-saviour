# 懦夫救星服务端

## 进度
### 2021-09-22 14:05 
* 将之前的*mysql*数据库改为mongo
* 将之前的*controller*从原生的*net/http*改为*gin*

## 笔记

### fatih/set.v0
(github链接)[https://github.com/fatih/set]  
基于hash的Set数据结构。同时提供了线程安全和非线程安全的版本。  
此set已经是不再维护的状态，官方建议使用第三方的包来替代。

### mysql设置相关
```sql
mysql -uroot -p
CREATE DATABASE IF NOT EXISTS test_cw DEFAULT CHARACTER SET utf8;
create user 'cw_test' identified by 'test12345';
grant ALL privileges ON test_cw.* to 'cw_test'@'%';
exit;
```

### docker
redis
```sh
docker run -v $PWD/conf/redis.conf:/usr/local/etc/redis/redis.conf -p 7001:6379 --name myredis redis redis-server /usr/local/etc/redis/redis.conf
```

mysql
```sh
docker run --name cs_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -d mysql
```