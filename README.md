
# Go REST API example
Простой REST API для телефонной книги

## Функционал:
- возможность создавать записи в телефонной книге;
- просматривать записи;
- изменять записи;
- удалять записи.

## Порядок установки
```bash
git clone https://github.com/s-antoshkin/go-api-example
cd go-api-example
go get -u
$ psql -p 5432 -h localhost -U postgres -c "CREATE DATABASE phonebook"
$ psql -h localhost -p 5432 -d phonebook -U postgres -f phonebook.sql
```
### Скопировать `.env.example` и назвать его `.env`:
```
cp .env.example .env
```
Заполнить переменные окружения:
```bash
DB_USERNAME = postgres #Имя пользователя БД
DB_PASSWORD = postgres #Пароль от БД
DB_HOST = localhost #Хост для подключения к БД
DB_PORT = 5432 #Порт для подключения к БД
DB_NAME = phonebook #Имя базы данных
```

## Запуск приложения:
```
go run .
```

## Эндпоинты API:

- `GET("/api/v1/records")` - запускает метод `getRecords` и возвращает данные всех записей;
- `GET("/api/v1/records/:id")` - запускает метод `getRecors` и возвращает данные записи (одной, по `id`);
- `POST("/api/v1/records")` - запускает метод `addRecors`, добавляет запись в БД;
- `PUT("/api/v1/records/:id")` - запускает метод `updateRecord`, обновляет данные записи;
- `DELETE("/api/v1/records/:id")` - запускает метод `deleteRecord`, удаляет запись в БД.

