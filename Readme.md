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
classDiagram
    class Customer {
        +OID : GUID
        +Name : VARCHAR
        +Address : VARCHAR
    }

    class Order {
        +OrderID : INTEGER
        +CustomerOID : GUID
        +OrderDate : DATE
        +TotalAmount : DECIMAL
    }

    class Product {
        +ProductID : INTEGER
        +Name : VARCHAR
        +Price : DECIMAL
    }

    class OrderItem {
        +OrderItemID : INTEGER
        +OrderID : INTEGER
        +ProductID : INTEGER
        +Quantity : INTEGER
    }

    Customer --> Order : places
    Order --> OrderItem : contains
    Product --> OrderItem : includes

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

