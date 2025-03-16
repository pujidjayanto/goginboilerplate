package product

import (
	"slices"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ProductNameLike(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name ILIKE ?", "%"+name+"%")
	}
}

func Sort(sortBy, sortDirection string) func(db *gorm.DB) *gorm.DB {
	column := defaultColumnToSort
	// can use map, but let make this simple for now since it is small array
	if slices.Contains(sortableColumns, sortBy) {
		column = sortBy
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: column}, Desc: sortDirection == "desc"})
	}
}
