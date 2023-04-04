package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/qin-guan/WorldSkills/Kazan2019/entity"
	"github.com/qin-guan/WorldSkills/Kazan2019/repository"
	"time"
)

type catalog struct {
	repository repository.Catalog
}

type Catalog interface {
	Find(ctx context.Context, sn int) (*entity.Asset, error)
	Search(
		ctx context.Context,
		name string,
		groupID int,
		departmentID int,
		warrantyDateStart time.Time,
		warrantyDateEnd time.Time,
	) ([]*entity.Asset, error)
}

func NewCatalog(repository repository.Catalog) Catalog {
	return &catalog{
		repository,
	}
}

func (c *catalog) Find(ctx context.Context, sn int) (*entity.Asset, error) {
	if sn < 1 {
		return nil, errors.New(fmt.Sprintf("asset sn is invalid: %d", sn))
	}

	asset, err := c.repository.Find(ctx, sn)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("asset with sn does not exist: %d", sn))
	}

	return asset, nil
}

func (c *catalog) Search(
	ctx context.Context,
	name string,
	groupID int,
	departmentID int,
	warrantyDateStart time.Time,
	warrantyDateEnd time.Time,
) ([]*entity.Asset, error) {
	if groupID < 1 {
		return nil, errors.New(fmt.Sprintf("group ID is invalid: %d", groupID))
	}

	if departmentID < 1 {
		return nil, errors.New(fmt.Sprintf("department ID is invalid: %d", groupID))
	}

	assets, err := c.repository.FindMany(ctx, name, groupID, departmentID, warrantyDateStart, warrantyDateEnd)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("database error"))
	}

	return assets, nil
}
