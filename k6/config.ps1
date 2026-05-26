# Общие значения по умолчанию для k6 (PowerShell-скрипты в k6/scripts/).
# Единая точка BASE_URL: k6/base-url.txt

$BaseUrlFile = Join-Path $PSScriptRoot "base-url.txt"
if (Test-Path $BaseUrlFile) {
    $K6_BASE_URL = (Get-Content $BaseUrlFile -Raw).Trim()
}
else {
    $K6_BASE_URL = 'http://localhost:8004/api'
}
