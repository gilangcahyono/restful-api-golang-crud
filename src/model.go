package src

import (
	_ "github.com/lib/pq"
)

type Product struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	Price      string     `json:"price"`
	Stock      string     `json:"stock"`
	Category   string     `json:"category"`
	Created_at string     `json:"created_at"`
	Updated_at NullString `json:"updated_at"`
}

func Get() ([]Product, error) {

	db := CreateConnection()
	defer db.Close()
	var products []Product
	sqlStatement := `SELECT * FROM products`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return products, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.Category, &product.Created_at, &product.Updated_at)
		products = append(products, product)
	}

	return products, err
}

func GetById(id int64) (Product, error) {

	db := CreateConnection()
	defer db.Close()
	var product Product
	sqlStatement := `SELECT * FROM products WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.Category, &product.Created_at, &product.Updated_at)
	return product, err
}

func Create(product Product) int64 {

	db := CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO products (name, price, stock, category) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, product.Name, product.Price, product.Stock, product.Category).Scan(&id)
	if err != nil {
		id = -1
	}

	return id
}

func Update(id int64, product Product) int64 {

	db := CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE products SET name=$2, price=$3, stock=$4, category=$5, updated_at=now() WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, product.Name, product.Price, product.Stock, product.Category)
	if err != nil {
		return -1
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return rowsAffected
	}

	return rowsAffected
}

func Delete(id int64) int64 {

	db := CreateConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM products WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return -1
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return rowsAffected
	}

	return rowsAffected
}
