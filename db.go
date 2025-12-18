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
func TambahDataPelanggan(Nama_Customer string, Email string) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/payment")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "INSERT INTO datacustomer (Nama_Customer,Email) VALUES (?,?)"
	result, err := db.Exec(query, Nama_Customer, Email)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data berhasil ditambahkan dengan ID  :", lastID)

}

// Hapus data pelanggan versi pertama  ada notifikasi tidak ada nama

// func HapusDataPelanggan(Nama_Customer string) {
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/payment")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	query := "DELETE FROM datacustomer WHERE Nama_Customer=?"
// 	result, err := db.Exec(query, Nama_Customer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	hapusPelanggan, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if hapusPelanggan == 0 {
// 		fmt.Println("Tidak ada nama tersebuut: ", Nama_Customer)

// 	} else {
// 		fmt.Printf("Berhasil menghapus %d data dengan Nama_Customer : %s\n ", hapusPelanggan, Nama_Customer)
// 	}

// }

// Delete data  versi kedua

func HapusDataPelanggan(Nama_Customer string) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/payment")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM datacustomer WHERE Nama_Customer =?", Nama_Customer)
	if err != nil {
		log.Fatal(err)
	}
	hapuspelanggan, _ := result.RowsAffected()
	fmt.Printf("Data terhapus: %d (Nama: %s)\n", hapuspelanggan, Nama_Customer)
}
