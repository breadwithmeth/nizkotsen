// filepath: /Users/shrvse/go/src/github.com/breadwithmeth/nizkotsen/internal/database/mysql.go
package database

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/breadwithmeth/nizkotsen/internal/config"

    _ "github.com/go-sql-driver/mysql" // Импорт драйвера MySQL
)

func Connect() *sql.DB {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Строка подключения к базе данных
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBName,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }

    // See "Important settings" section.
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    return db
}