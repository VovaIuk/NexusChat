function Import-ProjectEnv {
    param([string]$Root)

    $files = @(
        (Join-Path $Root ".env"),
        (Join-Path $Root "backend\.env")
    )

    $loaded = @()
    foreach ($envFile in $files) {
        if (-not (Test-Path $envFile)) {
            continue
        }
        Get-Content $envFile | ForEach-Object {
            if ($_ -match '^\s*([^#=]+)=(.*)$') {
                $name = $matches[1].Trim()
                $value = $matches[2].Trim().Trim('"').Trim("'")
                Set-Item -Path "env:$name" -Value $value
            }
        }
        $loaded += $envFile
    }

    if ($env:POSTGRES_URL -and (-not $env:POSTGRES_USER -or -not $env:POSTGRES_DB)) {
        if ($env:POSTGRES_URL -match '^postgres(?:ql)?://([^:]+):([^@]+)@[^/]+/([^?]+)') {
            if (-not $env:POSTGRES_USER) {
                $env:POSTGRES_USER = [uri]::UnescapeDataString($matches[1])
            }
            if (-not $env:POSTGRES_DB) {
                $env:POSTGRES_DB = [uri]::UnescapeDataString($matches[3])
            }
        }
    }

    if ($loaded.Count -eq 0) {
        Write-Error "Не найден .env в корне проекта или backend\.env"
    }

    if (-not $env:POSTGRES_USER -or -not $env:POSTGRES_DB) {
        Write-Error "В .env должны быть POSTGRES_USER и POSTGRES_DB (или POSTGRES_URL в backend\.env)"
    }

    Write-Host "Postgres: user=$($env:POSTGRES_USER), db=$($env:POSTGRES_DB) (from $($loaded -join ', '))"
}
