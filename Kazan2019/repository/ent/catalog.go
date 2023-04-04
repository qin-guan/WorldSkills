package ent

import (
	"context"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/asset"
	"github.com/qin-guan/WorldSkills/Kazan2019/repository"
	"time"
)

type catalog struct {
	client *ent.Client
}

func NewCatalog(client *ent.Client) repository.Catalog {
	return &catalog{
		client,
	}
}

func (c *catalog) Find(ctx context.Context, sn int) (*ent.Asset, error) {
	return c.client.Asset.Query().
		Where(asset.AssetSN(sn)).
		Only(ctx)
}

func (c *catalog) FindMany(
	ctx context.Context,
	name string,
	groupID int,
	departmentLocationID int,
	warrantyDateStart time.Time,
	warrantyDateEnd time.Time,
) ([]*ent.Asset, error) {
	query := c.client.Asset.Query()

	if name != "" {
		query = query.Where(asset.AssetName(name))
	}

	if groupID > 0 {
		query = query.Where(asset.AssetGroupID(groupID))
	}

	if departmentLocationID > 0 {
		query = query.Where(asset.DepartmentLocationID(departmentLocationID))
	}

	if !warrantyDateStart.IsZero() {
		query = query.Where(asset.WarrantyDateGTE(warrantyDateStart))
	}

	if !warrantyDateEnd.IsZero() {
		query = query.Where(asset.WarrantyDateLTE(warrantyDateEnd))
	}

	return query.All(ctx)
}
