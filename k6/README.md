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

### 3. Полный нагрузочный тест (~5 мин)

```powershell
.\k6\scripts\run-load.ps1
```

Во время теста смотрите в Grafana: HTTP RPS, latency, пагинация истории, строки в `messages`, goroutines.

## Запуск без скриптов (из корня проекта)

```powershell
k6 run k6/scenarios/smoke.js
k6 run k6/scenarios/load.js
```

С параметрами:

```powershell
k6 run -e BASE_URL=http://localhost:8003/api -e SCROLL_PAGES=10 k6/scenarios/load.js
```

## Переменные окружения k6

| Переменная      | По умолчанию                 | Описание                    |
|-----------------|------------------------------|-----------------------------|
| `BASE_URL`      | `http://localhost:8003/api`  | URL API backend             |
| `CHAT_ID`       | `1`                          | Чат для скролла истории     |
| `HISTORY_LIMIT` | `50`                         | `limit` в запросе истории   |
| `SCROLL_PAGES`  | `8`                          | Страниц пагинации за проход |

Пользователи из `db_init`: `alice_dev` / `alice1234`, `bob_codes` / `bob12345`, и т.д.

## Если API через nginx

```powershell
.\k6\scripts\run-load.ps1 -BaseUrl "http://localhost:8004/api"
```

## Что смотреть в Grafana

- **Сообщений в БД** — после `seed.ps1`
- **HTTP RPS / latency p95** — во время `run-load.ps1`
- **Запросы истории / пагинация** — скролл по большому чату
- **PostgreSQL** — подключения, cache hit
- **Goroutines / память** — нагрузка на Go backend
