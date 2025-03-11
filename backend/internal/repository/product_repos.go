package repository

import (
	"backend/internal/model"
	"context"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	FindByID(ctx context.Context, id int) (*model.Product, error)
	FindAll(ctx context.Context) ([]model.Product, error)
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	query := `INSERT INTO products (name, description, price, stock) 
              VALUES (:name, :description, :price, :stock)
              RETURNING id, created_at, updated_at`

	_, err := r.db.NamedQueryContext(ctx, query, product)
	return err
}

func (r *productRepository) FindByID(ctx context.Context, id int) (*model.Product, error) {
	var product model.Product
	query := `SELECT * FROM products WHERE id = $1`
	err := r.db.GetContext(ctx, &product, query, id)
	return &product, err
}

func (r *productRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	query := `SELECT * FROM products`
	err := r.db.SelectContext(ctx, &products, query)
	return products, err
}
