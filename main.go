package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/sqlite"
)

func main() {
	db := ConnectDB("base.sqlite")
	defer db.Close()
	query := "Select * from Unite"
	GetWhatever(db, query)

}

func ConnectDB(dbCible string) *sql.DB {
	db, err := sql.Open("sqlite", dbCible)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base de données : %v", err)
	}
	// Vérifie que la connexion est opérationnelle
	if err := db.Ping(); err != nil {
		log.Fatalf("Impossible de se connecter à la base : %v", err)
	}
	return db
}

func GetWhatever(db *sql.DB, query string) {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Erreur lors de la query : %v", err)
	}
	defer rows.Close()
	colonnes, _ := rows.Columns()
	fmt.Println(rows.Columns())
	for _, col := range colonnes {
		fmt.Println(col)
	}
	// Ici faut ajouter deux lignes de slice, une avec les pointeurs de l'autre
	rows.Scan()
}
