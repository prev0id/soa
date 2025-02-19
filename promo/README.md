# Сервис промо

## Зона ответственности
Сервис постов отвечает за создание, хранение и управление промокодами и комментариями

## Основные задачи
- Создание и обновление промокодов/комментариев
- Управление комментариями, включая вложенные ответы
- Обеспечение фильтрации и поиска предложений

## Границы сервиса
- Не занимается аутентификацией и управлением учетными записями пользователей
- Не обрабатывает статистику просмотров, лайков или комментариев — это зона сервиса аналитики

## ER-диаграмма

```mermaid
erDiagram
    PROMO {
        string id PK
        string user_id FK
        string title
        string description
        string promo_code
        timestamp created_at
        timestamp updated_at
    }

    COMMENT {
        string id PK
        string post_id FK
        string user_id FK
        string parant_id
        string text
        timestamp created_at
        timestamp updated_at
    }

    LIKE_PROMO {
        string promo_id FK
        string like_id FK
    }

    LIKE_COMMENT {
        string comment_id FK
        string like_id FK
    }
    
    LIKE {
        string id PK
        string user_id FK
        timestamp created_at
    }

    PROMO ||--o{ COMMENT : ""
    PROMO ||--o{ LIKE_PROMO : ""
    COMMENT ||--o{ LIKE_COMMENT : ""
    LIKE_PROMO ||--o{ LIKE : ""
    LIKE_COMMENT ||--o{ LIKE : ""

```