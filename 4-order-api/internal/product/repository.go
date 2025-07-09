package product

import (
	"purple/4-order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	conn *db.Db
}

func NewProductRepository(conn *db.Db) *ProductRepository {
	return &ProductRepository{
		conn: conn,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.conn.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.conn.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.conn.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.conn.DB.Delete(&Product{}, id)
	return result.Error
}

func (repo *ProductRepository) GetByIds(ids []uint) ([]Product, error) {
	var products []Product
	result := repo.conn.DB.Find(&products, ids)
	if result.Error != nil {
		return nil, result.Error 
	}
	return products, nil
}