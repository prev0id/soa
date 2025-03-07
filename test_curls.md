*Регистрация:*
```sh
curl -X POST http://localhost:8081/register -H "Content-Type: application/json" -d '{"login": "user1", "password": "pass123", "email": "user1@example.com"}'
```

*Аутентификация:*
```sh
curl -X POST http://localhost:8081/login -H "Content-Type: application/json" -d '{"login": "user1", "password": "pass123"}'
```

*Получение профиля:*
```sh
curl -X GET http://localhost:8081/profile -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDE2MzgxOTIsInVzZXJfaWQiOjF9.uHVO9CmRhqZiAsZNYFsPht7P3yZDj61SgDk1jw9hZtw"
```

*Обновление профиля:*
```sh
curl -X PUT http://localhost:8081/profile -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDE2MzgxOTIsInVzZXJfaWQiOjF9.uHVO9CmRhqZiAsZNYFsPht7P3yZDj61SgDk1jw9hZtw" -H "Content-Type: application/json" -d '{"first_name": "Иван", "last_name": "Иванов", "birth_date": "1990-01-01", "phone": "+79991234567"}'
```sh
