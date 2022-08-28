package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Record struct {
	ID     int64
	Name   string
	Amount int64
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func run() error {
	db, err := sql.Open("sqlite3", "bank.db")
	if err != nil {
		return err
	}

	if err := createTable(db); err != nil {
		return err
	}

	for {
		if err := showRecords(db); err != nil {
			return err
		}

		if err := inputRecord(db); err != nil {
			return err
		}
	}
}

func createTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS accounts (
			id     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name   TEXT NOT NULL,
			amount TEXT NOT NULL
	);`

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func showRecords(db *sql.DB) error {
	fmt.Println("全件表示")
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		return err
	}
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Amount); err != nil {
			return err
		}
		fmt.Printf("[%d] Name:%s Amount:%d\n", r.ID, r.Name, r.Amount)
	}
	fmt.Println("--------")

	return nil
}

func inputRecord(db *sql.DB) error {
	var r Record

	fmt.Print("Name >")
	fmt.Scan(&r.Name)

	fmt.Print("Amount>")
	fmt.Scan(&r.Amount)

	const sql = "INSERT INTO accounts(name, Amount) values (?,?)"
	_, err := db.Exec(sql, r.Name, r.Amount)
	if err != nil {
		return err
	}

	return nil
}
