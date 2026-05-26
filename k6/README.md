# Нагрузочное тестирование k6 (ручной запуск)

Тесты запускаются **локально** с установленным [k6](https://grafana.com/docs/k6/latest/set-up/install-k6/).  
Стек приложения и мониторинга — через `docker compose up -d`.

## Установка k6 (Windows)

```powershell
choco install k6
# или: winget install k6 --source winget
```

Проверка: `k6 version`

## Порядок для демонстрации на курсовой

### 0. Поднять сервисы

```powershell
docker compose up -d
```

Откройте Grafana: http://localhost:13000 → дашборд **NexusChat Overview**.

### 1. Заполнить БД (большие данные)

Из корня проекта:

```powershell
.\k6\scripts\seed.ps1
```

500 000 сообщений:

```powershell
.\k6\scripts\seed.ps1 -Count 500000
```

> Использует `docker compose exec db` и учётные данные из **корневого `.env`** (как `docker compose`), при необходимости дополняет из `backend/.env`.  
> Повторный запуск удаляет старые `k6 seed #...` и вставляет заново.

Linux/macOS:

```bash
chmod +x k6/scripts/seed.sh
./k6/scripts/seed.sh
```

### 2. Smoke-тест (30 сек)

```powershell
.\k6\scripts\run-smoke.ps1
```

### 3. Полный нагрузочный тест HTTP (~5 мин)

```powershell
.\k6\scripts\run-load.ps1
```

Во время теста смотрите в Grafana: HTTP RPS, latency, пагинация истории, строки в `messages`, goroutines.

### 4. WebSocket: создание сообщений (~3 мин)

```powershell
.\k6\scripts\run-ws.ps1
```

Логин по HTTP → `ws://…/v1/ws` → `auth` → `chat_message` → broadcast с `message_id` (запись в БД).

В Grafana: **WebSocket соединения**, **Messages created**, **Ws broadcast**.

## Запуск без скриптов (из корня проекта)

```powershell
k6 run k6/scenarios/smoke.js
k6 run k6/scenarios/load.js
k6 run k6/scenarios/ws-messages.js
```

С параметрами:

```powershell
k6 run -e BASE_URL=http://localhost:8004/api -e SCROLL_PAGES=10 k6/scenarios/load.js
```

## Конфигурация

URL API по умолчанию задаётся в **`k6/base-url.txt`** (единая точка).  
Его читают и PowerShell-скрипты (`k6/config.ps1` → `run-smoke.ps1`, `run-load.ps1`, `run-ws.ps1`), и прямой `k6 run` через `k6/lib/config.js` (переменная `BASE_URL`).

## Переменные окружения k6

| Переменная              | По умолчанию                 | Описание                              |
|-------------------------|------------------------------|---------------------------------------|
| `BASE_URL`              | см. `k6/base-url.txt`        | URL API backend                       |
| `CHAT_ID`               | `1`                          | Чат (история и WS-сообщения)          |
| `HISTORY_LIMIT`         | `50`                         | `limit` в запросе истории             |
| `SCROLL_PAGES`          | `8`                          | Страниц пагинации за проход           |
| `AUTH_WAIT_MS`          | `800`                        | Пауза после auth перед отправкой (WS) |
| `MESSAGES_PER_SESSION`  | `2`                          | Сообщений за одно WS-соединение       |
| `MESSAGE_INTERVAL_MS`   | `400`                        | Интервал между сообщениями (WS)       |
| `CLOSE_AFTER_MS`        | `4500`                       | Закрытие сокета после отправки (WS)   |

Пользователи из `db_init`: `alice_dev` / `alice1234`, `bob_codes` / `bob12345`, и т.д.  
WS-тест (`ws-messages.js`) использует только **alice** и **bob** — участников `CHAT_ID=1`.

## Если API через nginx

Измените URL в `k6/base-url.txt` или передайте параметр:

```powershell
.\k6\scripts\run-load.ps1 -BaseUrl "http://localhost:8004/api"
.\k6\scripts\run-ws.ps1 -BaseUrl "http://localhost:8004/api"
```

## Что смотреть в Grafana

- **Сообщений в БД** — после `seed.ps1`
- **HTTP RPS / latency p95** — во время `run-load.ps1`
- **Запросы истории / пагинация** — скролл по большому чату
- **PostgreSQL** — подключения, cache hit
- **Goroutines / память** — нагрузка на Go backend
- **WebSocket / Messages created** — во время `run-ws.ps1`
