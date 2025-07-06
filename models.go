package main

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Message struct {
    From      string `json:"from"`
    To        string `json:"to"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type Group struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    Members []string `json:"members"`
}

type GroupMessage struct {
    GroupID   string `json:"group_id"`
    From      string `json:"from"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}
type BroadcastMessage struct {
    From      string `json:"from"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}