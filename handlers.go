package main

import (
    "net/http"
    "strings"
    "time"
    "github.com/gin-gonic/gin"
)

// Signup Handler
func signupHandler(c *gin.Context) {
    var req User
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    req.Username = strings.TrimSpace(req.Username)
    req.Password = strings.TrimSpace(req.Password)
    if req.Username == "" || req.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password required"})
        return
    }
    // Check if user exists
    if _, err := GetUser(req.Username); err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }
    // Hash password
    hashed, err := HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user := User{Username: req.Username, Password: hashed}
    if err := SaveUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
}

// Login Handler
func loginHandler(c *gin.Context) {
    var req User
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    req.Username = strings.TrimSpace(req.Username)
    req.Password = strings.TrimSpace(req.Password)
    if req.Username == "" || req.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password required"})
        return
    }
    user, err := GetUser(req.Username)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }
    if !CheckPasswordHash(req.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// Send Direct Message
func sendDirectMessageHandler(c *gin.Context) {
    var msg Message
    if err := c.ShouldBindJSON(&msg); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    msg.From = strings.TrimSpace(msg.From)
    msg.To = strings.TrimSpace(msg.To)
    msg.Message = strings.TrimSpace(msg.Message)
    if msg.From == "" || msg.To == "" || msg.Message == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "from, to, and message required"})
        return
    }
    msg.Timestamp = time.Now().Format(time.RFC3339)
    if err := SaveMessage(msg); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store message"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}

// Fetch Chat History
func getChatHistoryHandler(c *gin.Context) {
    user1 := strings.TrimSpace(c.Query("user1"))
    user2 := strings.TrimSpace(c.Query("user2"))
    if user1 == "" || user2 == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "user1 and user2 required"})
        return
    }
    messages, err := GetMessages(user1, user2)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
        return
    }
    c.JSON(http.StatusOK, messages)
}
// Create group
func createGroupHandler(c *gin.Context) {
    var req struct {
        Name    string   `json:"name"`
        Members []string `json:"members"`
    }
    if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" || len(req.Members) == 0 {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }
    group, err := CreateGroup(req.Name, req.Members)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create group"})
        return
    }
    c.JSON(200, group)
}

// Send a message to group
func sendGroupMessageHandler(c *gin.Context) {
    var msg GroupMessage
    if err := c.ShouldBindJSON(&msg); err != nil || msg.GroupID == "" || msg.From == "" || msg.Message == "" {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }
    msg.Timestamp = time.Now().Format(time.RFC3339)
    if err := SaveGroupMessage(msg); err != nil {
        c.JSON(500, gin.H{"error": "Failed to save message"})
        return
    }
    c.JSON(200, gin.H{"message": "Message sent"})
}

// Get group chat history
func getGroupMessagesHandler(c *gin.Context) {
    groupID := c.Query("group_id")
    if groupID == "" {
        c.JSON(400, gin.H{"error": "group_id required"})
        return
    }
    messages, err := GetGroupMessages(groupID)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to get messages"})
        return
    }
    c.JSON(200, messages)
}
// Send a broadcast message
func sendBroadcastMessageHandler(c *gin.Context) {
    var msg BroadcastMessage
    if err := c.ShouldBindJSON(&msg); err != nil || msg.From == "" || msg.Message == "" {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }
    msg.Timestamp = time.Now().Format(time.RFC3339)
    if err := SaveBroadcastMessage(msg); err != nil {
        c.JSON(500, gin.H{"error": "Failed to save broadcast"})
        return
    }
    c.JSON(200, gin.H{"message": "Broadcast sent"})
}

// Get all broadcast messages
func getBroadcastMessagesHandler(c *gin.Context) {
    messages, err := GetBroadcastMessages()
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to get broadcasts"})
        return
    }
    c.JSON(200, messages)
}