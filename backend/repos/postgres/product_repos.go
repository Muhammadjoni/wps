// package postgres

// import (
// 	"context"
// 	"backend/internal/model"

// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type ProductRepository struct {
// 	db *pgxpool.Pool
// }

// func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
// 	return &ProductRepository{db: db}
// }

// func (r *ProductRepository) GetProducts(ctx context.Context) ([]model.Product, error) {
// 	query := `SELECT id, name, description, price, stock FROM products`
// 	rows, err := r.db.Query(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var products []model.Product
// 	for rows.Next() {
// 		var p model.Product
// 		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, p)
// 	}
// 	return products, nil
// }


package postgres

import (
    "context"
    "backend/internal/model"
    "errors"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/georgysavva/scany/v2/pgxscan"
)

type ProductRepository struct {
    pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
    return &ProductRepository{pool: pool}
}

func (r *ProductRepository) Create(ctx context.Context, product *model.Product) error {
    query := `
        INSERT INTO products (name, description, price, stock)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at
    `
    
    return r.pool.QueryRow(ctx, query,
        product.Name,
        product.Description,
        product.Price,
        product.Stock,
    ).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
}

func (r *ProductRepository) FindByID(ctx context.Context, id int) (*model.Product, error) {
    var product model.Product
    query := `SELECT * FROM products WHERE id = $1`
    
    err := pgxscan.Get(ctx, r.pool, &product, query, id)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]model.Product, error) {
    var products []model.Product
    query := `SELECT * FROM products`
    
    rows, err := r.pool.Query(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    if err := pgxscan.ScanAll(&products, rows); err != nil {
        return nil, err
    }
    return products, nil
}