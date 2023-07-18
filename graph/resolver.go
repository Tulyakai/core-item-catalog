package graph

import "com.ktb.ai.core-item-catalog/app/infrastructure/gormrepository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	GormCatalogRepository *gormrepository.GormCatalogRepository
}
