# WebSocket нагрузка: создание сообщений через /v1/ws (~3 мин).
param(
    [string]$BaseUrl,
    [int]$ChatId = 1,
    [int]$AuthWaitMs = 800,
    [int]$MessagesPerSession = 2
)

$ErrorActionPreference = "Stop"
$Root = Resolve-Path (Join-Path $PSScriptRoot "../..")
Set-Location $Root

. (Join-Path $Root "k6/config.ps1")
if (-not $BaseUrl) { $BaseUrl = $K6_BASE_URL }

if (-not (Get-Command k6 -ErrorAction SilentlyContinue)) {
    Write-Error "k6 не найден. Установите: https://grafana.com/docs/k6/latest/set-up/install-k6/"
}

Write-Host "Grafana: http://localhost:13000  |  API: $BaseUrl"
Write-Host "WebSocket load (chat messages)..."

k6 run `
    -e "BASE_URL=$BaseUrl" `
    -e "CHAT_ID=$ChatId" `
    -e "AUTH_WAIT_MS=$AuthWaitMs" `
    -e "MESSAGES_PER_SESSION=$MessagesPerSession" `
    "$Root/k6/scenarios/ws-messages.js"
