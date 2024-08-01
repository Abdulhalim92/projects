# Викторина
- Основной функционал: Учет вопросов, варинтов ответов и правильных ответов
- Расширение: Учет попыток, подсказок и команда для подсказок

````html
quizgo/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── quiz/
│   │   ├── quiz.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── user/
│   │   ├── user.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── hint/
│   │   ├── hint.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── http/
│   │   ├── handler/
│   │   │   ├── hint_handler.go
│   │   │   ├── user_handler.go
│   │   │   └── borrow_handler.go
│   │   ├── middleware/
│   │   │   └── auth_middleware.go
│   │   └── router.go
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── database.go
│   │   └── migrations/
│   └── util/
│       └── util.go
├── pkg/
│   └── model/
│       ├── quiz.go
│       ├── user.go
│       └── hint.go
├── docs/
│   ├── api/
│   └── architecture/
├── scripts/
│   └── setup.sh
├── .env
├── .gitignore
├── go.mod
└── go.sum
````
