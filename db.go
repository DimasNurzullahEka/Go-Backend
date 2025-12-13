package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DataCustomer struct {
	ID            int
	Nama_Customer string
	Email         string
}

func GetDataCustomer() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/payment")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Koneksi database berhasil...")

	rows, err := db.Query("SELECT id, Nama_Customer, Email FROM DataCustomer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Data Customer:")
	for rows.Next() {
		var g DataCustomer
		if err := rows.Scan(&g.ID, &g.Nama_Customer, &g.Email); err != nil {
			log.Println("Error scan:", err)
		}
		fmt.Printf("ID: %d | Nama_Customer: %s | Email: %s\n", g.ID, g.Nama_Customer, g.Email)
	}
}
