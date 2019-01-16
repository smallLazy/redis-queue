# redis-queue

该项目主要是redis的一个小案例，Go 操作redis案例
###Ubuntu16.04 安装redis服务端

```
redis sudo apt-get install redis-server #安装redis
ps -ef | grep redis # 查看进程
redissudo service redis-server start # 启动redis
ps -ef | grep reids # 再次查看进程
redis-cli # 进入redis
```

### golang操作redis
1. 如若没有 Redis 包，则先执行 `go get github.com/go-redis/redis`
2. 在主文件 main.go 中导入`redis` 包 `import github.com/go-redis/redis`
3. go连接 redis服务端
```
func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	pong, err := client.Ping().Result()
	return client
}
```
4. 在main函数中获取reids对象，执行redis操作
```
client := GetRedisClient()
err := client.Set("name", "lazy", 0).Err()
if err != nil {
  log.Println(err)
}
value, err := client.Get("name").Result()
if err != nil {
  log.Println(err)
}
log.Println("name=:", value)
```
