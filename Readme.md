# URL Shortener

>Этот проект представляет собой сервис для сокращения URL-адресов, написанный на языке Go. Он предоставляет HTTP API для создания коротких URL и перенаправления по ним. В качестве хранилища используется PostgreSQL.

```mermaid
graph TD
    A[Начальная страница] -->|Выбрать| B{Уже есть аккаунт?};
    B -- Да --> C(Форма входа);
    B -- Нет --> D(Форма регистрации);
    C --> E[Отправить данные];
    D --> F[Отправить данные];
    E -->|Валидация| G{Данные верны?};
    F -->|Валидация| H{Данные верны?};
    G -- Да --> I[Успешный вход];
    G -- Нет --> J(Показать ошибку);
    H -- Да --> K[Успешная регистрация];
    H -- Нет --> L(Показать ошибку);
    I --> M[Личный кабинет];
    K --> M;
    M --> N(Создать/Управлять URL);
    J --> C;
    L --> D;
```

-----

```mermaid
erDiagram
    CUSTOMER {
        GUID OID PK
        VARCHAR Name
        VARCHAR Address
    }

    ORDER {
        INTEGER OrderID PK
        GUID CustomerOID FK
        DATE OrderDate
        DECIMAL TotalAmount
    }

    PRODUCT {
        INTEGER ProductID PK
        VARCHAR Name
        DECIMAL Price
    }

    ORDERITEM {
        INTEGER OrderItemID PK
        INTEGER OrderID FK
        INTEGER ProductID FK
        INTEGER Quantity
    }

    CUSTOMER ||--o{ ORDER : places
    ORDER ||--o{ ORDERITEM : contains
    PRODUCT ||--o{ ORDERITEM : includes
```

------

```mermaid
sequenceDiagram
    participant Пользователь as U
    participant Клиент as C
    participant Сервер/API as S
    participant База_Данных as BD

    U->>C: Вводит полный URL и нажимает "Сократить"
    C->>S: POST /api/shorten (полный_url, токен_авторизации)
    S->>S: Проверка токена и получение user_id
    S->>S: Генерация короткого URL
    S->>BD: Запрос: INSERT INTO urls (full_url, short_url, user_id)
    BD-->>S: Ответ: запись создана
    S-->>C: Ответ: JSON с short_url
    C-->>U: Отображение сокращенного URL
```

## Запуск

#### **Git clone**

```shell
git clone https://github.com/golkity/Url_Shorter_Go.git
cd Url_Shorter_Go
```

#### Установка зависимостей

```shell
go mod tidy
```

#### Миграция

```shell
make migrat
```

#### Docker

запуск докер-образа

```shell
make dc-up
```

удаление докер-образа

```shell
make dc-down
```

