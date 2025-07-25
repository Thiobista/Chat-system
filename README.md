Go Telegram-Style Chat API
A simple, modular chat backend using Go, Gin, and Redis.
Supports user authentication (JWT), direct messaging, groups, and broadcasts.
🚀 Setup
Prerequisites
- Go 1.19+
- Redis (local or Docker)
1. Clone & Install

git clone https://github.com/yourusername/go-telegram-chat.git
cd go-telegram-chat
go mod tidy

2. Start Redis

docker run -p 6379:6379 redis
# OR: redis-server (if installed)

3. Run the Server

go run *.go

Server runs on `:8083` by default.
📚 API Endpoints
Endpoint	Method	Auth	Description
/signup	POST	❌ Public	Register a new user
/login	POST	❌ Public	Login, returns JWT
/message	POST	✅ Bearer	Send direct message
/messages	GET	✅ Bearer	Get chat history
/group/create	POST	✅ Bearer	Create group
/group/message	POST	✅ Bearer	Send group message
/group/messages	GET	✅ Bearer	Get group messages
/broadcast	POST	✅ Bearer	Send broadcast
/broadcasts	GET	✅ Bearer	All broadcasts
Example: Signup

curl -X POST http://localhost:8083/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"mypassword"}'

Example: Login

curl -X POST http://localhost:8083/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"mypassword"}'

Example: Send Direct Message

curl -X POST http://localhost:8083/message \
  -H "Authorization: Bearer {TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{"from":"alice", "to":"bob", "message":"Hi Bob!"}'

Example: Get Chat History

curl -X GET "http://localhost:8083/messages?user1=alice&user2=bob" \
  -H "Authorization: Bearer {TOKEN}"

⚙️ Design Decisions

- Redis as DB: Fast, simple for prototyping chat/message flows.
- JWT Auth: Scalable, stateless session management.
- Gin: Lightweight, easy routing and JSON support.
- Separation: Handlers, storage, models, middleware, and utilities are modular.



📝 License
MIT 