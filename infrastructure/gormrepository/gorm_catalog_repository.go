package gormrepository

import (
	"com.ktb.ai.core-item-catalog/domain"
	"gorm.io/gorm"
)

type GormCatalogRepository struct {
	db *gorm.DB
}

func NewGormCatalogRepository(db *gorm.DB) *GormCatalogRepository {
	return &GormCatalogRepository{db: db}
}

// get all catalog
func (r *GormCatalogRepository) GetCatalogs() ([]*domain.ItemCatalog, error) {
	var catalogs []*domain.ItemCatalog
	if err := r.db.Find(&catalogs).Error; err != nil {
		return nil, err
	}
	return catalogs, nil
}

// get catalog by id
func (r *GormCatalogRepository) GetCatalogByID(id int) (*domain.ItemCatalog, error) {
	var catalog domain.ItemCatalog
	if err := r.db.First(&catalog, id).Error; err != nil {
		return nil, err
	}
	return &catalog, nil
}

// create catalog
func (r *GormCatalogRepository) CreateCatalog(catalog *domain.ItemCatalog) (*domain.ItemCatalog, error) {
	if err := r.db.Create(catalog).Error; err != nil {
		return nil, err
	}
	return catalog, nil
}

// update catalog
func (r *GormCatalogRepository) UpdateCatalog(catalog *domain.ItemCatalog) (*domain.ItemCatalog, error) {
	if err := r.db.Save(catalog).Error; err != nil {
		return nil, err
	}
	return catalog, nil
}

// delete catalog
func (r *GormCatalogRepository) DeleteCatalog(catalog *domain.ItemCatalog) error {
	if err := r.db.Delete(catalog).Error; err != nil {
		return err
	}
	return nil
}


