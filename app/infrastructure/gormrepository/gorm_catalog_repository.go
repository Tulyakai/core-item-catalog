package gormrepository

import (
	"com.ktb.ai.core-item-catalog/app/domain"
	"gorm.io/gorm"
)

type GormCatalogRepository struct {
	db *gorm.DB
}

func NewGormCatalogRepository(db *gorm.DB) *GormCatalogRepository {
	return &GormCatalogRepository{db: db}
}

func (repo GormCatalogRepository) Get(id int) (*domain.ItemCatalog, error) {
	var item domain.ItemCatalog
	err := repo.db.First(&item, &domain.ItemCatalog{ID: id}).Error
	return &item, err
}

func (repo GormCatalogRepository) Update(catalog *domain.ItemCatalog) error {
	return repo.db.Updates(&catalog).Error
}

func (repo GormCatalogRepository) Delete(id string) error {
	return repo.db.Delete(&id).Error
}
