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

// MODOP
// func GetWhatever(db *sql.DB, query string) {
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatalf("Erreur lors de la query : %v", err)
// 	}
// 	defer rows.Close() // Assurez-vous de fermer les lignes une fois terminé

// 	// Récupérer les colonnes
// 	colonnes, err := rows.Columns()
// 	if err != nil {
// 		log.Fatalf("Erreur lors de la récupération des colonnes : %v", err)
// 	}

// 	// Afficher les colonnes
// 	fmt.Println("Colonnes:", colonnes)

// 	// Préparer un slice d'interfaces pour stocker les valeurs des colonnes
// 	valeurs := make([]interface{}, len(colonnes))
// 	pointeurs := make([]interface{}, len(colonnes))

// 	// Associer chaque élément du slice à son pointeur respectif
// 	for i := range valeurs {
// 		pointeurs[i] = &valeurs[i]
// 	}

// 	// Parcourir chaque ligne
// 	for rows.Next() {
// 		// Scanner les colonnes de la ligne dans les pointeurs
// 		err := rows.Scan(pointeurs...)
// 		if err != nil {
// 			log.Fatalf("Erreur lors du scan des lignes : %v", err)
// 		}

// 		// Créer un slice pour stocker les valeurs converties en string
// 		valeursLigne := make([]string, len(colonnes))
// 		for i, val := range valeurs {
// 			if val != nil {
// 				valeursLigne[i] = fmt.Sprintf("%v", val)
// 			} else {
// 				valeursLigne[i] = "NULL" // Si la valeur est NULL
// 			}
// 		}

// 		// Afficher la ligne
// 		fmt.Println("Ligne:", valeursLigne)
// 	}

// 	// Vérifier les erreurs éventuelles après la boucle
// 	if err := rows.Err(); err != nil {
// 		log.Fatalf("Erreur après lecture des lignes : %v", err)
// 	}
// }
