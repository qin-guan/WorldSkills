package payload

import (
	"net/http"
	"time"
)

type CatalogFindResponse struct {
	Asset AssetResponse `json:"asset,omitempty"`
}

func (c *CatalogFindResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type AssetResponse struct {
	ID                   int    `json:"id,omitempty"`
	AssetGroupID         int    `json:"assetGroupID,omitempty"`
	EmployeeID           int    `json:"employeeID,omitempty"`
	DepartmentLocationID int    `json:"departmentLocationID,omitempty"`
	AssetSN              int    `json:"assetSN,omitempty"`
	AssetName            string `json:"assetName,omitempty"`
	Description          string `json:"description,omitempty"`
	WarrantyDateUnix     int64  `json:"warrantyDate,omitempty"`
}

func NewCatalogFindResponse(
	id int,
	assetGroupID int,
	employeeID int,
	departmentLocationID int,
	assetSN int,
	assetName string,
	description string,
	warrantyDate time.Time,
) *CatalogFindResponse {
	return &CatalogFindResponse{
		Asset: AssetResponse{
			ID:                   id,
			AssetGroupID:         assetGroupID,
			EmployeeID:           employeeID,
			DepartmentLocationID: departmentLocationID,
			AssetSN:              assetSN,
			AssetName:            assetName,
			Description:          description,
			WarrantyDateUnix:     warrantyDate.Unix(),
		},
	}
}
