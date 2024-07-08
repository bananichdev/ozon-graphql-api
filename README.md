# OZON graphql API
## Документация
Посмотреть всю документацию и выполнить запросы к API можно (Предварительно, запустив сервис), перейдя по ссылке [http://localhost:8000/](http://localhost:8000/)
## Как запустить?
1. Склонировать репозиторий
### Doker-compose
<b>Должен быть установлен docker</b><br>
Если хотите видеть логи контейнеров:
```shell
docker compose up
```
Если не хотите видеть логи контейнеров:
```shell
docker compose up -d
```
Остановить сервис:
```shell
docker compose stop
```

Чтобы поменять хранилище (in-memory или PostgreSQL): откройте файл .docker.env и установите нужное значение переменной MODE (inmemory или db)
