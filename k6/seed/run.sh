#!/bin/sh
# Запуск внутри контейнера с примонтированным /seed (опционально).
# Для ручного запуска с хоста используйте k6/scripts/seed.ps1 или seed.sh.
set -eu

COUNT="${MESSAGE_COUNT:-100000}"
TEMPLATE="/seed/bulk_messages.sql.template"
OUTPUT="/tmp/bulk_messages.sql"

echo "Seeding ${COUNT} messages into chat_id=1 ..."

psql -h db -U "$POSTGRES_USER" -d "$POSTGRES_DB" -v ON_ERROR_STOP=1 -c \
  "CREATE INDEX IF NOT EXISTS idx_messages_chat_id_time ON messages (chat_id, time DESC);"

sed "s/__COUNT__/${COUNT}/g" "$TEMPLATE" > "$OUTPUT"

psql -h db -U "$POSTGRES_USER" -d "$POSTGRES_DB" -v ON_ERROR_STOP=1 -f "$OUTPUT"

echo "Done. Total messages in chat 1:"
psql -h db -U "$POSTGRES_USER" -d "$POSTGRES_DB" -t -c \
  "SELECT COUNT(*) FROM messages WHERE chat_id = 1;"
