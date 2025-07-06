package main

import (
    // "context"
    "encoding/json"
    "fmt"
    "sort"
    "github.com/go-redis/redis/v8"
    "github.com/google/uuid"
     
)

func SaveUser(user User) error {
    return rdb.HSet(ctx, "user:"+user.Username, "password", user.Password).Err()
}

func GetUser(username string) (User, error) {
    pw, err := rdb.HGet(ctx, "user:"+username, "password").Result()
    if err == redis.Nil {
        return User{}, fmt.Errorf("user not found")
    } else if err != nil {
        return User{}, err
    }
    return User{Username: username, Password: pw}, nil
}

func SaveMessage(msg Message) error {
    users := []string{msg.From, msg.To}
    sort.Strings(users)
    chatKey := fmt.Sprintf("messages:%s:%s", users[0], users[1])

    msgJSON, err := json.Marshal(msg)
    if err != nil {
        return err
    }
    return rdb.RPush(ctx, chatKey, msgJSON).Err()
}

func GetMessages(user1, user2 string) ([]Message, error) {
    users := []string{user1, user2}
    sort.Strings(users)
    chatKey := fmt.Sprintf("messages:%s:%s", users[0], users[1])

    vals, err := rdb.LRange(ctx, chatKey, 0, -1).Result()
    if err != nil {
        return nil, err
    }
    messages := make([]Message, 0, len(vals))
    for _, v := range vals {
        var m Message
        if err := json.Unmarshal([]byte(v), &m); err == nil {
            messages = append(messages, m)
        }
    }
    return messages, nil
}

// Create a new group
func CreateGroup(name string, members []string) (Group, error) {
    id := uuid.New().String()
    group := Group{ID: id, Name: name, Members: members}
    // Store group info as hash
    if err := rdb.HSet(ctx, "group:"+id, "name", name).Err(); err != nil {
        return Group{}, err
    }
    // Store members as set
    for _, user := range members {
        if err := rdb.SAdd(ctx, "group:"+id+":members", user).Err(); err != nil {
            return Group{}, err
        }
    }
    return group, nil
}

// Add user to group
func AddUserToGroup(groupID, username string) error {
    return rdb.SAdd(ctx, "group:"+groupID+":members", username).Err()
}

// Save a message to group chat
func SaveGroupMessage(msg GroupMessage) error {
    msgJSON, err := json.Marshal(msg)
    if err != nil {
        return err
    }
    return rdb.RPush(ctx, "group:"+msg.GroupID+":messages", msgJSON).Err()
}

// Get group messages
func GetGroupMessages(groupID string) ([]GroupMessage, error) {
    vals, err := rdb.LRange(ctx, "group:"+groupID+":messages", 0, -1).Result()
    if err != nil {
        return nil, err
    }
    messages := make([]GroupMessage, 0, len(vals))
    for _, v := range vals {
        var m GroupMessage
        if err := json.Unmarshal([]byte(v), &m); err == nil {
            messages = append(messages, m)
        }
    }
    return messages, nil
}

// Save a broadcast message
func SaveBroadcastMessage(msg BroadcastMessage) error {
    msgJSON, err := json.Marshal(msg)
    if err != nil {
        return err
    }
    return rdb.RPush(ctx, "broadcast:messages", msgJSON).Err()
}

// Get all broadcast messages
func GetBroadcastMessages() ([]BroadcastMessage, error) {
    vals, err := rdb.LRange(ctx, "broadcast:messages", 0, -1).Result()
    if err != nil {
        return nil, err
    }
    messages := make([]BroadcastMessage, 0, len(vals))
    for _, v := range vals {
        var m BroadcastMessage
        if err := json.Unmarshal([]byte(v), &m); err == nil {
            messages = append(messages, m)
        }
    }
    return messages, nil
}