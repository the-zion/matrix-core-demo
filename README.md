# Demo of cube base in kratos

## demo构建
```
make build
```

## mysql数据库
```
docker pull mysql:8.0
docker run -itd --name mysql -v </path>:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 mysql:8.0
```

# 启动
端口：http(8000) grpc(9000)
```
./bin/cube-core
```
# 接口测试
```
人员注册 
POST 127.0.0.1:8000/users
body：{"user":{"email":"12345@qq.com", "password":"123456", "username":"eric"}}

人员登录
POST 127.0.0.1:8000/users/login
body：{"user":{"email":"aaa", "password":"123"}}
```



