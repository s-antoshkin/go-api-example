package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func insert(name, phone string) (pgconn.CommandTag, error) {
	return db.Exec(context.TODO(), "INSERT INTO phonebook VALUES (default, $1, $2)", name, phone)
}

func remove(id int) (pgconn.CommandTag, error) {
	return db.Exec(context.Background(), "DELETE FROM phonebook WHERE id=$1", id)
}

func update(id int, name, phone string) (pgconn.CommandTag, error) {
	return db.Exec(context.Background(), "UPDATE phonebook SET name=$1, phone=$2 WHERE id=$3", name, phone, id)
}

func readOne(id int) (Record, error) {
	var rec Record
	row := db.QueryRow(context.Background(), "SELECT * FROM phonebook WHERE id=$1 ORDER BY id", id)
	return rec, row.Scan(&rec.ID, &rec.Name, &rec.Phone)
}

func readAll(str string) ([]Record, error) {
	var rows pgx.Rows
	var err error
	if str != "" {
		rows, err = db.Query(context.Background(), "SELECT * FROM phonebook WHERE name LIKE $1 ORDER BY id", "%"+str+"%")

	} else {
		rows, err = db.Query(context.Background(), "SELECT * FROM phonebook ORDER BY id")
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var recs = make([]Record, 0)
	var rec Record
	for rows.Next() {
		if err = rows.Scan(&rec.ID, &rec.Name, &rec.Phone); err != nil {
			return nil, err
		}
		recs = append(recs, rec)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recs, nil
}
