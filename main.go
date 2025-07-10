package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
)

// Global Redis and context, initialized at startup
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

    // Public endpoints
    router.POST("/signup", signupHandler)
    router.POST("/login", loginHandler)

    // Protected endpoints (require JWT)
    auth := router.Group("/")
    auth.Use(AuthMiddleware())
    {
        auth.POST("/message", sendDirectMessageHandler)
        auth.GET("/messages", getChatHistoryHandler)
        auth.POST("/group/create", createGroupHandler)
        auth.POST("/group/message", sendGroupMessageHandler)
        auth.GET("/group/messages", getGroupMessagesHandler)
        auth.POST("/broadcast", sendBroadcastMessageHandler)
        auth.GET("/broadcasts", getBroadcastMessagesHandler)
    }

    router.Run(":8083")
}