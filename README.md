# Notification Service

Сервис уведомлений для проекта *Finance Tracker*, отвечающий за рассылку уведомлений пользователям.

## Запуск
```bash
git clone https://github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification
cd notification
DOCKER_BUILDKIT=1 docker buildx build --platform linux/amd64 -f build/Dockerfile -t notification .
docker run -d \
  -p 5001:5001 \
  -p 50055:50051 \
  -e APP_MODE=production \
  -e HTTP_SERVER_PORT=5001 \
  -e GRPC_SERVER_PORT=50051 \
  -e KAFKA_USERNAME=<your kafka username> \
  -e KAFKA_PASSWORD=<your kafka password> \
  -e KAFKA_BOOTSTRAP_SERVERS=<your kafka bootstrap servers> \
  -e KAFKA_SSL_ROOT_CERT=<your kafka ssl root cert path> \
  -e TG_BOT_TOKEN=<your tg bot token> \
  notification
```

## Структура проекта
```
notification/
├── cmd/                    # точка входа в сервис
├── config/                 # конфиги
├── internal/
│   ├── pkg/
│       │── generated/          # сгенерированный из proto код
│       │── grpcsrv/            # grpc сервер 
│       │── handlers/           # обработчики kafka событий
│       │── httpsrv/            # http сервер 
│       │── jobs/               # background джобы
│       │── managers/           # менеджеры бизнес логики
│       │── senders/            # отправители kafka событий
│       │── receivers/          # получатели kafka событий
│       └── tg/                 # интеграция с telegram
├── library/                # обёртки и расширения над стандартными библиотеками
├── go.mod
└── README.md
```
