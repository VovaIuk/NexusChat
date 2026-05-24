# Полный нагрузочный тест k6 (~5 мин).
param(
    [string]$BaseUrl = "http://localhost:8003/api",
    [int]$ChatId = 1,
    [int]$HistoryLimit = 50,
    [int]$ScrollPages = 8
)

$ErrorActionPreference = "Stop"
$Root = Resolve-Path (Join-Path $PSScriptRoot "../..")
Set-Location $Root

if (-not (Get-Command k6 -ErrorAction SilentlyContinue)) {
    Write-Error "k6 не найден. Установите: https://grafana.com/docs/k6/latest/set-up/install-k6/"
}

Write-Host "Grafana: http://localhost:13000  |  API: $BaseUrl"
Write-Host "Starting load test..."

k6 run `
    -e "BASE_URL=$BaseUrl" `
    -e "CHAT_ID=$ChatId" `
    -e "HISTORY_LIMIT=$HistoryLimit" `
    -e "SCROLL_PAGES=$ScrollPages" `
    "$Root/k6/scenarios/load.js"
