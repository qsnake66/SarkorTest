package repository

import (
	"database/sql"

	SarkorTest "github.com/qsnake66/ProductWerehouse"
)

type ProductPostgresRepository struct {
	db *sql.DB
}

func NewProductPostgresRepository(db *sql.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{db: db}
}

func (p *ProductPostgresRepository) CreateProduct(product SarkorTest.Product) (int, error) {

	var id int

	query := `
        INSERT INTO products (name, price, description, quantity, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `
	err := p.db.QueryRow(query, product.Name, product.Price, product.Description, product.Quantity, product.CreatedAt, product.UpdatedAt).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (p *ProductPostgresRepository) GetProductById(id int) (SarkorTest.Product, error) {

	var product SarkorTest.Product

	query := `
	SELECT *
	FROM products
	WHERE id = $1;
	`
	err := p.db.QueryRow(query, id).Scan(
		&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt,
	)

	if err != nil {
		return SarkorTest.Product{}, err
	}

	return product, nil
}
func (p *ProductPostgresRepository) GetAllProduct() ([]SarkorTest.Product, error) {
	var products []SarkorTest.Product

	query := `
		SELECT *
		FROM products;
	`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var product SarkorTest.Product

		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
func (p *ProductPostgresRepository) UpdateProduct(id int, product SarkorTest.Product) error {

	query := `
	UPDATE products
	SET name = $1, price = $2, description = $3, quantity = $4, updated_at = $5
	WHERE id = $6;
	`
	_, err := p.db.Exec(query, product.Name, product.Price, product.Description, product.Quantity, product.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil
}
func (p *ProductPostgresRepository) DeleteProduct(id int) error {

	query := `
	DELETE FROM products
	WHERE id = $1;
	`
	_, err := p.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
