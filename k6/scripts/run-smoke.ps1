# Быстрый smoke-тест k6 (30 сек). Запускать из корня проекта не обязательно.
$ErrorActionPreference = "Stop"
$Root = Resolve-Path (Join-Path $PSScriptRoot "../..")
Set-Location $Root

if (-not (Get-Command k6 -ErrorAction SilentlyContinue)) {
    Write-Error "k6 не найден. Установите: https://grafana.com/docs/k6/latest/set-up/install-k6/"
}

k6 run -e BASE_URL=http://localhost:8003/api "$Root/k6/scenarios/smoke.js"
