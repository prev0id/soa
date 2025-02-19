# API Gateway

## Зона ответственности
API Gateway — это центральная точка входа для всех запросов, поступающих от фронта. Он отвечает за аутентификацию, маршрутизацию запросов к соответствующим сервисам и агрегацию ответов.

## Основные задачи
- Прием HTTP-запросов от фронта
- Аутентификация и валидация входящих запросов
- Логирование запросов и ошибок
- Сбор и хранение статистических данных по вызовам API

## Границы сервиса
- API не содержит бизнес-логику приложения
- Не занимается постоянным хранением пользовательских или бизнес-данны

## ER-диаграмма

```mermaid
erDiagram
    REQUEST {
        string id PK
        string user_id
        string endpoint
        string method
        int status_code
        int response_time
        string request_id
        timestamp created_at
    }
    
    ERROR {
        string id PK
        string request_id FK
        string error_message
        string error_type
        string stack_trace
        datetime timestamp
    }

    REQUEST ||--o{ ERROR : ""
```