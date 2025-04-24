*Регистрация:*
```sh
curl -X POST http://localhost:8081/register -H "Content-Type: application/json" -d '{"login": "user1", "password": "pass123", "email": "user1@example.com"}'
curl -X POST http://localhost:8080/user/register -H "Content-Type: application/json" -d '{"login": "user1", "password": "pass123", "email": "user1@example.com"}' -i
```

*Аутентификация:*
```sh
curl -X POST http://localhost:8081/login -H "Content-Type: application/json" -d '{"login": "user1", "password": "pass123"}'
curl -X POST http://localhost:8080/user/login -H "Content-Type: application/json" -d '{"login": "user1", "password": "pass123"}' -i
```

*Получение профиля:*
```sh
curl -X GET http://localhost:8081/profile -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDE2MzgxOTIsInVzZXJfaWQiOjF9.uHVO9CmRhqZiAsZNYFsPht7P3yZDj61SgDk1jw9hZtw"
curl -X GET http://localhost:8080/user/profile -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI4OTk5NzEsInVzZXJfaWQiOjF9.s8Ci2-sEeyT3LLVd6j4gVfkRxevVVGRPeJtcecr9a9Q" -i
```

*Обновление профиля:*
```sh
curl -X PUT http://localhost:8081/profile -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDE2MzgxOTIsInVzZXJfaWQiOjF9.uHVO9CmRhqZiAsZNYFsPht7P3yZDj61SgDk1jw9hZtw" -H "Content-Type: application/json" -d '{"first_name": "Иван", "last_name": "Иванов", "birth_date": "1990-01-01", "phone": "+79991234567"}'
curl -X PUT http://localhost:8080/user/profile -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI4OTk5NzEsInVzZXJfaWQiOjF9.s8Ci2-sEeyT3LLVd6j4gVfkRxevVVGRPeJtcecr9a9Q" -H "Content-Type: application/json" -d '{"first_name": "Иван", "last_name": "Иванов", "birth_date": "2001-01-01", "phone": "+79991234567"}' -i
```sh
