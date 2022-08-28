package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Record struct {
	ID     int64
	Name   string
	Amount int64
}

func run() error {
	fromID := flag.Int("from", -1, "source ID")
	toID := flag.Int("to", -1, "destination ID")
	amount := flag.Int("amount", -1, "amount")
	flag.Parse()

	if *fromID < 0 || *toID < 0 || *amount < 0 {
		return errors.New("Specified negative number")
	}

	db, err := sql.Open("sqlite3", "bank.db")
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var fromAmount int64
	var toAmount int64

	row := tx.QueryRow("SELECT amount FROM accounts WHERE id = ?", *fromID)
	if err := row.Scan(&fromAmount); err != nil {
		tx.Rollback()
		return err
	}
	row = tx.QueryRow("SELECT amount FROM accounts WHERE id = ?", *toID)
	if err := row.Scan(&toAmount); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("fromID: %d, toID: %d, amount: %d\n", fromID, toID, fromAmount)

	const updateSQL = "UPDATE accounts SET amount = ? WHERE id = ?"
	fromDiff := fromAmount - int64(*amount)
	// if _, err := tx.Exec(updateSQL, fromDiff, fromID); err != nil || fromDiff < 0 {
	if _, err := tx.Exec(updateSQL, fromDiff, fromID); err != nil {
		tx.Rollback()
		return err
	}

	toDiff := toAmount + int64(*amount)
	if _, err := tx.Exec(updateSQL, toDiff, toID); err != nil {
		tx.Rollback()
		return err
	}
	if fromDiff < 0 {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}
