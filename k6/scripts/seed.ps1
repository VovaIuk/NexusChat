# Заполнение БД тестовыми сообщениями (ручной запуск).
# Учётные данные: корневой .env (как docker compose), затем backend/.env
#
#   .\k6\scripts\seed.ps1
#   .\k6\scripts\seed.ps1 -Count 500000

param(
    [int]$Count = 100000
)

$ErrorActionPreference = "Stop"
$Root = Resolve-Path (Join-Path $PSScriptRoot "../..")
Set-Location $Root

. (Join-Path $PSScriptRoot "load-env.ps1")
Import-ProjectEnv -Root $Root

$templatePath = Join-Path $Root "k6\seed\bulk_messages.sql.template"
$sql = @"
CREATE INDEX IF NOT EXISTS idx_messages_chat_id_time ON messages (chat_id, time DESC);
$((Get-Content $templatePath -Raw) -replace '__COUNT__', $Count)
"@

Write-Host "Seeding $Count messages into chat_id=1 via docker compose exec db ..."

$sql | docker compose exec -T db psql -U $env:POSTGRES_USER -d $env:POSTGRES_DB -v ON_ERROR_STOP=1
if ($LASTEXITCODE -ne 0) {
    Write-Error "seed failed (psql exit code $LASTEXITCODE)"
}

Write-Host "Done. Messages in chat 1:"
docker compose exec -T db psql -U $env:POSTGRES_USER -d $env:POSTGRES_DB -t -c `
    "SELECT COUNT(*) FROM messages WHERE chat_id = 1;"
if ($LASTEXITCODE -ne 0) {
    Write-Error "count query failed (psql exit code $LASTEXITCODE)"
}
