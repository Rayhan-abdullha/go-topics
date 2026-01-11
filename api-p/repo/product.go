package repo

import (
	"server/domain"
	"server/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	dbCon *sqlx.DB
}

func NewProductRepo(dbCon *sqlx.DB) ProductRepo {
	return &productRepo{
		dbCon: dbCon,
	}
}

func (r *productRepo) Create(pd *domain.Product) (*domain.Product, error) {
	query := `INSERT INTO products (title, description, price, image_url)
VALUES ($1, $2, $3, $4)
RETURNING id, title, description, price, image_url;`

	err := r.dbCon.QueryRow(query, pd.Title, pd.Description, pd.Price, pd.ImgUrl).
		Scan(&pd.ID, &pd.Title, &pd.Description, &pd.Price, &pd.ImgUrl)
	if err != nil {
		return nil, err
	}
	return pd, nil
}
func (r *productRepo) List() ([]*domain.Product, error) {
	query := `
		SELECT id, title, description, price, image_url
		FROM products
		ORDER BY id DESC
	`

	rows, err := r.dbCon.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*domain.Product, 0)

	for rows.Next() {
		var p domain.Product
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.Price,
			&p.ImgUrl,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
