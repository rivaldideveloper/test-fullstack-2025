package main

import(
"content"
"github.com/redis/ho-redis/v9"
)

var ctx = content.Background()
var rdb *redis.client

func initredis(){
rdb = redis.newclient(&redis.options){
addr: "localhost:6379"
DB: 0,
})
}
