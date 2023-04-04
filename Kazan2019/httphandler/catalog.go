package httphandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/qin-guan/WorldSkills/Kazan2019/payload"
	"github.com/qin-guan/WorldSkills/Kazan2019/usecase"
	"strconv"

	"net/http"
)

type catalog struct {
	usecase usecase.Catalog
}

func NewCatalog(usecase usecase.Catalog) http.Handler {
	m := chi.NewMux()
	c := &catalog{usecase}
	m.Get("/asset/{assetID}", c.Find)
	m.Get("/search", c.Search)
	return m
}

func (c *catalog) Find(w http.ResponseWriter, r *http.Request) {
	assetIDString := chi.URLParam(r, "assetID")
	assetID, err := strconv.Atoi(assetIDString)
	if err != nil {
		render.DefaultResponder(w, r, ErrInvalidRequest(err))
		return
	}

	asset, err := c.usecase.Find(r.Context(), assetID)
	if err != nil {
		render.DefaultResponder(w, r, ErrInvalidRequest(err))
		return
	}

	_ = render.Render(w, r,
		payload.NewCatalogFindResponse(
			asset.ID,
			asset.AssetGroupID,
			asset.EmployeeID,
			asset.DepartmentLocationID,
			asset.AssetSN,
			asset.AssetName,
			asset.Description,
			asset.WarrantyDate,
		))
}

func (c *catalog) Search(w http.ResponseWriter, r *http.Request) {
}
