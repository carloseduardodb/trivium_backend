@echo off

IF "%1"=="up" (
    go run cmd/migrate/main.go -command up
    GOTO End
)

IF "%1"=="down" (
    go run cmd/migrate/main.go -command down
    GOTO End
)

IF "%1"=="status" (
    go run cmd/migrate/main.go -command status
    GOTO End
)

IF "%1"=="create" (
    IF "%2"=="" (
        echo Name parameter is required for create command
        GOTO End
    )
    go run cmd/migrate/main.go -command create -name %2 -type sql
    GOTO End
)

echo Available commands: up, down, status, create
:End