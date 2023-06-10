# Ozon Intership
Реализация тестового задания по созданию сервиса генерации коротких ссылок.

## Описание функционала:
Проект имеет две реализации: grpc и rest. Выбор запускаемой реализации происходит при вызове цели Makefile (об этом
подробнее далее). Сервисы имеют сценарии работы как PostgreSQL и Redis. Выбор хранилища также определяется 
с помощью Makefile. Генерация ссылок происходит путем создания и обрезки хэша поданной пользователем ссылки.

## Запуск сервиса:
Сервис запускается путем использования одной из следующих целей Makefile:
``` 
make grpc_with_redis // запуск grpc-сервиса с использования redis в качестве хранилища
make grpc_with_postgres // запуск grpc-сервиса с использования postgres в качестве хранилища
make rest_with_redis // запуск rest-сервиса с использования redis в качестве хранилища
make rest_with_postgres // запуск rest-сервиса с использования postgres в качестве хранилища
make tests // запуск контейнера тестирования. Процесс тестов описан далее.
```

## Тестирование функционала:
Далее приложены изображения, демонстрирующие способ взаимодействие с запущенным сервисом. Файл .proto, необходимый для
тестирования работы grpc-сервиса, находится pkg/proto/file.proto. Адреса для тестрования:
```
    grpc - localhost:8080
    rest - localhost:3000
```
### Тестирование функицонала REST:
#### Post:
![Image alt](https://github.com/alexGrap/OzonIntership/blob/main/readmeImages/1.jpg)
#### Get:
![Image alt](https://github.com/alexGrap/OzonIntership/blob/main/readmeImages/2.jpg)

### Тестирование функционала GRPC:
#### Post:
![Image alt](https://github.com/alexGrap/OzonIntership/blob/main/readmeImages/3.jpg)
#### Get:
![Image alt](https://github.com/alexGrap/OzonIntership/blob/main/readmeImages/4.jpg)



## Запуск тестов:
Контейнер тестрования производит тестрование фукнций генерации, записи и возвращения ссылок. Файл test.log, который после
самостоятельной остановки работы контейнера находится в файлах контейнера test/test.log содержит 
информацию о ходе выполнения тестов. 
![Image alt](https://github.com/alexGrap/OzonIntership/blob/main/readmeImages/5.jpg)
