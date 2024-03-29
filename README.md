<h1>HTTP API</h1>
 HTTP API представляет собой веб-приложение для управления сегментами и пользователями. В проекте используется язык программирования Go. Приложение позволяет создавать, удалять сегменты, добавлять пользователей к сегментам, а также просматривать историю действий пользователей.

**Структура проекта:**


**cmd/main.go:**

- Основной файл программы, содержит точку входа для запуска веб-сервера.

**config/conf.yaml:** 

- Файл конфигурации, содержащий параметры подключения к базе данных.

**db/segments/segments.go:**

- Модуль для работы с базой данных, содержит методы для создания, удаления и управления сегментами.

**db/db_connection.go:** 

- Модуль для установки соединения с базой данных.

**handlers/handlers.go:** 

- Модуль обработчиков HTTP-запросов, содержит обработчики для создания, удаления сегментов, управления пользователями и другие операции.

**internal/config/config.go:** 

- Модуль для чтения конфигурационных данных из файла YAML.

 

 
**Как запустить:**
Для создания сегмента отправьте POST-запрос на /segment/add.
Для удаления сегмента отправьте DELETE-запрос на /segment/delete.
Для добавления пользователя к сегменту отправьте POST-запрос на /user/add.
Для удаления пользователя из сегмента отправьте POST-запрос на /user/delete.
Для просмотра сегментов пользователя отправьте GET-запрос на /user/return?id=<user_id>.
Для просмотра истории действий пользователя отправьте GET-запрос на /user/history.
Замечания
Перед запуском убедитесь, что у вас установлена и настроена PostgreSQL база данных.
Не забудьте внести необходимые параметры подключения к базе данных в файл config/conf.yaml.

Для Docker написаны файлы Dockerfile и docker-compose.yml. Сам проект поднимался и тестировался на Ubuntu, запросы отправлялись с помощью Postman на localhost:8080
