# Сервис пользователей

## Зона ответственности
Сервис пользователей отвечает за регистрацию, аутентификацию и управление учетными записями. Он хранит информацию о пользователях, их ролях, а также управляет сессиями

## Основные задачи
- Регистрация новых пользователей
- Аутентификация и авторизация
- Управление профилем и ролями пользователей
- Обработка сессий

## Границы сервиса
- Не отвечает за посты, комментарии или статистику

## ER-диаграмма

```mermaid
erDiagram
    USER {
        string id PK
        string name
        string login
        string password(hashed)
        int role_id FK
        timestamp created_at
        timestamp updated_at
    }

    SESSION {
        string id PK
        string user_id FK
        string token
        timestamp created_at
        timestamp expires_at
    }

    ROLE {
        int id PK
        string name
        timestamp created_at
        timestamp updated_at
    }

    USER ||--o{ SESSION : ""
    USER ||--o{ ROLE : ""
```
