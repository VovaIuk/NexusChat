#!/usr/bin/env bash
# Заполнение БД тестовыми сообщениями (ручной запуск).
# Учётные данные: корневой .env, затем backend/.env
#
#   ./k6/scripts/seed.sh
#   MESSAGE_COUNT=500000 ./k6/scripts/seed.sh

set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
cd "$ROOT"

COUNT="${MESSAGE_COUNT:-100000}"

load_env_file() {
  local file="$1"
  [[ -f "$file" ]] || return 0
  set -a
  # shellcheck disable=SC1090
  source "$file"
  set +a
}

load_env_file "$ROOT/.env"
load_env_file "$ROOT/backend/.env"

if [[ -n "${POSTGRES_URL:-}" ]] && { [[ -z "${POSTGRES_USER:-}" ]] || [[ -z "${POSTGRES_DB:-}" ]]; }; then
  if [[ "$POSTGRES_URL" =~ postgres(ql)?://([^:]+):[^@]+@[^/]+/([^?]+) ]]; then
    POSTGRES_USER="${POSTGRES_USER:-${BASH_REMATCH[2]}}"
    POSTGRES_DB="${POSTGRES_DB:-${BASH_REMATCH[3]}}"
    export POSTGRES_USER POSTGRES_DB
  fi
fi

if [[ -z "${POSTGRES_USER:-}" || -z "${POSTGRES_DB:-}" ]]; then
  echo "В .env нужны POSTGRES_USER и POSTGRES_DB (или POSTGRES_URL в backend/.env)" >&2
  exit 1
fi

echo "Postgres: user=${POSTGRES_USER}, db=${POSTGRES_DB}"

INDEX_SQL="CREATE INDEX IF NOT EXISTS idx_messages_chat_id_time ON messages (chat_id, time DESC);"
BODY="$(sed "s/__COUNT__/${COUNT}/g" "$ROOT/k6/seed/bulk_messages.sql.template")"

echo "Seeding ${COUNT} messages into chat_id=1 ..."
{
  echo "$INDEX_SQL"
  echo "$BODY"
} | docker compose exec -T db psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -v ON_ERROR_STOP=1

echo "Done. Messages in chat 1:"
docker compose exec -T db psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -t -c \
  "SELECT COUNT(*) FROM messages WHERE chat_id = 1;"
