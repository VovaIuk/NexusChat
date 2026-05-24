CREATE INDEX IF NOT EXISTS idx_messages_chat_id_time ON messages (chat_id, time DESC);
