package userdb

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type dbContainer struct {
	db *sql.DB
}

/* Does the user have an active subscription for a product, regardless of stripe plan name? */
func CheckAccess(email string, productName string) bool {
	DBC := getDB()
	defer DBC.db.Close()
	var productNameHolder string
	sqlUserExists := `SELECT product_name FROM products p JOIN transactions t ON p.stripe_product_id = t.stripe_product_id WHERE t.email=$1 AND t.paid_until_ts<extract(epoch from now())`
	rows, err := DBC.db.Query(sqlUserExists, email)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rowErr := rows.Scan(&productNameHolder)
		if rowErr != nil {
			panic(err)
		}
		if productNameHolder == productName {
			return true
		}
	}
	rowsErr := rows.Err()
	if rowsErr != nil {
		panic(rowsErr)
	}
	return false
}

func AddUser(email string, stripeCustomerID string) error {
	DBC := getDB()
	defer DBC.db.Close()
	sqlUserInsertion := `INSERT INTO users (email, stripe_customer_id) VALUES ('` + email + `', '` + stripeCustomerID + `');`
	_, err := DBC.db.Exec(sqlUserInsertion)
	return err
}

func getDB() dbContainer {
	db, dbErr := sql.Open("postgres", "host=localhost port=5432 user=useroni dbname=usersdb password=omitequalthumbroseparched sslmode=disable")
	if dbErr != nil {
		panic(dbErr)
	}

	dbc := dbContainer{db}
	return dbc
}
