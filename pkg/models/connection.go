package models

import (
	"database/sql"

	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (  
    username = "root"
    password = "iamnot19"
    hostname = "127.0.0.1:3306"
    dbname   = "books"
)

func dsn(dbName string) string {  
    return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func Connection() (*sql.DB, error) {  
	db, err := sql.Open("mysql", dsn(dbname))  
	if err != nil {  
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(20)
    db.SetConnMaxLifetime(time.Minute * 5)

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    err = db.PingContext(ctx)
    if err != nil {
        log.Printf("Errors %s pinging DB", err)
        return nil, err
    }
    log.Printf("Connected to DB %s successfully\n", dbname)
	return db, err
}

