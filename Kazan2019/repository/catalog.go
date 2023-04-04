package repository

import (
	"context"
	"github.com/qin-guan/WorldSkills/Kazan2019/entity"
	"time"
)

type CatalogReader interface {
	Find(ctx context.Context, sn int) (*entity.Asset, error)
	FindMany(
		ctx context.Context,
		name string,
		groupID int,
		departmentID int,
		warrantyDateStart time.Time,
		warrantyDateEnd time.Time,
	) ([]*entity.Asset, error)
}

type Catalog interface {
	CatalogReader
}
