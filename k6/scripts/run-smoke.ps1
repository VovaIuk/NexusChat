# Быстрый smoke-тест k6 (30 сек). Запускать из корня проекта не обязательно.
param([string]$BaseUrl)

$ErrorActionPreference = "Stop"
$Root = Resolve-Path (Join-Path $PSScriptRoot "../..")
Set-Location $Root

. (Join-Path $Root "k6/config.ps1")
if (-not $BaseUrl) { $BaseUrl = $K6_BASE_URL }

if (-not (Get-Command k6 -ErrorAction SilentlyContinue)) {
    Write-Error "k6 не найден. Установите: https://grafana.com/docs/k6/latest/set-up/install-k6/"
}

k6 run -e "BASE_URL=$BaseUrl" "$Root/k6/scenarios/smoke.js"
