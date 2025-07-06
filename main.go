package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
)

var (
    rdb *redis.Client
    ctx context.Context
)

func main() {
    ctx = context.Background()
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    router := gin.Default()

    router.POST("/signup", signupHandler)
    router.POST("/login", loginHandler)
    router.POST("/message", sendDirectMessageHandler)
    router.GET("/messages", getChatHistoryHandler)
    router.POST("/group/create", createGroupHandler)
router.POST("/group/message", sendGroupMessageHandler)
router.GET("/group/messages", getGroupMessagesHandler)
router.POST("/broadcast", sendBroadcastMessageHandler)
router.GET("/broadcasts", getBroadcastMessagesHandler)

    router.Run(":8081")
}