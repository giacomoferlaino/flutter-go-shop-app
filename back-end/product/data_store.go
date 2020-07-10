package product

import (
	"database/sql"
)

type dataStore struct {
	db *sql.DB
}

func (store *dataStore) getAll() ([]Product, error) {
	query := "select * from product"
	rows, err := store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.ImageURL, &product.IsFavorite)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (store *dataStore) getByID(id int64) ([]Product, error) {
	query := "select * from product where id = ?"
	rows, err := store.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.ImageURL, &product.IsFavorite)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (store *dataStore) add(product Product) (int64, error) {
	query := `INSERT INTO product (title, description, price, imageUrl, isFavorite)
	VALUES(?, ?, ?, ?, ?)`
	result, err := store.db.Exec(query, product.Title, product.Description, product.Price, product.ImageURL, product.IsFavorite)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (store *dataStore) deleteByID(id int64) (int64, error) {
	query := `DELETE FROM product WHERE id=?`
	result, err := store.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affectedRows, nil
}
