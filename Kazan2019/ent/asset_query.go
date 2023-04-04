// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/asset"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assetgroup"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assetphoto"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assettransferlog"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/departmentlocation"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/employee"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// AssetQuery is the builder for querying Asset entities.
type AssetQuery struct {
	config
	ctx                    *QueryContext
	order                  []OrderFunc
	inters                 []Interceptor
	predicates             []predicate.Asset
	withAssetPhoto         *AssetPhotoQuery
	withAssetTransferLog   *AssetTransferLogQuery
	withDepartmentLocation *DepartmentLocationQuery
	withEmployee           *EmployeeQuery
	withAssetGroup         *AssetGroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AssetQuery builder.
func (aq *AssetQuery) Where(ps ...predicate.Asset) *AssetQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AssetQuery) Limit(limit int) *AssetQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AssetQuery) Offset(offset int) *AssetQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AssetQuery) Unique(unique bool) *AssetQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AssetQuery) Order(o ...OrderFunc) *AssetQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryAssetPhoto chains the current query on the "AssetPhoto" edge.
func (aq *AssetQuery) QueryAssetPhoto() *AssetPhotoQuery {
	query := (&AssetPhotoClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(assetphoto.Table, assetphoto.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, asset.AssetPhotoTable, asset.AssetPhotoColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssetTransferLog chains the current query on the "AssetTransferLog" edge.
func (aq *AssetQuery) QueryAssetTransferLog() *AssetTransferLogQuery {
	query := (&AssetTransferLogClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(assettransferlog.Table, assettransferlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, asset.AssetTransferLogTable, asset.AssetTransferLogColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDepartmentLocation chains the current query on the "DepartmentLocation" edge.
func (aq *AssetQuery) QueryDepartmentLocation() *DepartmentLocationQuery {
	query := (&DepartmentLocationClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(departmentlocation.Table, departmentlocation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asset.DepartmentLocationTable, asset.DepartmentLocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmployee chains the current query on the "Employee" edge.
func (aq *AssetQuery) QueryEmployee() *EmployeeQuery {
	query := (&EmployeeClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asset.EmployeeTable, asset.EmployeeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssetGroup chains the current query on the "AssetGroup" edge.
func (aq *AssetQuery) QueryAssetGroup() *AssetGroupQuery {
	query := (&AssetGroupClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(assetgroup.Table, assetgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asset.AssetGroupTable, asset.AssetGroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Asset entity from the query.
// Returns a *NotFoundError when no Asset was found.
func (aq *AssetQuery) First(ctx context.Context) (*Asset, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asset.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AssetQuery) FirstX(ctx context.Context) *Asset {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Asset ID from the query.
// Returns a *NotFoundError when no Asset ID was found.
func (aq *AssetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asset.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AssetQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Asset entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Asset entity is found.
// Returns a *NotFoundError when no Asset entities are found.
func (aq *AssetQuery) Only(ctx context.Context) (*Asset, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asset.Label}
	default:
		return nil, &NotSingularError{asset.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AssetQuery) OnlyX(ctx context.Context) *Asset {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Asset ID in the query.
// Returns a *NotSingularError when more than one Asset ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AssetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = &NotSingularError{asset.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AssetQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Assets.
func (aq *AssetQuery) All(ctx context.Context) ([]*Asset, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Asset, *AssetQuery]()
	return withInterceptors[[]*Asset](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AssetQuery) AllX(ctx context.Context) []*Asset {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Asset IDs.
func (aq *AssetQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(asset.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AssetQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AssetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AssetQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AssetQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AssetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AssetQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AssetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AssetQuery) Clone() *AssetQuery {
	if aq == nil {
		return nil
	}
	return &AssetQuery{
		config:                 aq.config,
		ctx:                    aq.ctx.Clone(),
		order:                  append([]OrderFunc{}, aq.order...),
		inters:                 append([]Interceptor{}, aq.inters...),
		predicates:             append([]predicate.Asset{}, aq.predicates...),
		withAssetPhoto:         aq.withAssetPhoto.Clone(),
		withAssetTransferLog:   aq.withAssetTransferLog.Clone(),
		withDepartmentLocation: aq.withDepartmentLocation.Clone(),
		withEmployee:           aq.withEmployee.Clone(),
		withAssetGroup:         aq.withAssetGroup.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithAssetPhoto tells the query-builder to eager-load the nodes that are connected to
// the "AssetPhoto" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithAssetPhoto(opts ...func(*AssetPhotoQuery)) *AssetQuery {
	query := (&AssetPhotoClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAssetPhoto = query
	return aq
}

// WithAssetTransferLog tells the query-builder to eager-load the nodes that are connected to
// the "AssetTransferLog" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithAssetTransferLog(opts ...func(*AssetTransferLogQuery)) *AssetQuery {
	query := (&AssetTransferLogClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAssetTransferLog = query
	return aq
}

// WithDepartmentLocation tells the query-builder to eager-load the nodes that are connected to
// the "DepartmentLocation" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithDepartmentLocation(opts ...func(*DepartmentLocationQuery)) *AssetQuery {
	query := (&DepartmentLocationClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withDepartmentLocation = query
	return aq
}

// WithEmployee tells the query-builder to eager-load the nodes that are connected to
// the "Employee" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithEmployee(opts ...func(*EmployeeQuery)) *AssetQuery {
	query := (&EmployeeClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withEmployee = query
	return aq
}

// WithAssetGroup tells the query-builder to eager-load the nodes that are connected to
// the "AssetGroup" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithAssetGroup(opts ...func(*AssetGroupQuery)) *AssetQuery {
	query := (&AssetGroupClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAssetGroup = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AssetGroupID int `json:"AssetGroupID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Asset.Query().
//		GroupBy(asset.FieldAssetGroupID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AssetQuery) GroupBy(field string, fields ...string) *AssetGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AssetGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = asset.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AssetGroupID int `json:"AssetGroupID,omitempty"`
//	}
//
//	client.Asset.Query().
//		Select(asset.FieldAssetGroupID).
//		Scan(ctx, &v)
func (aq *AssetQuery) Select(fields ...string) *AssetSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AssetSelect{AssetQuery: aq}
	sbuild.label = asset.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AssetSelect configured with the given aggregations.
func (aq *AssetQuery) Aggregate(fns ...AggregateFunc) *AssetSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AssetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !asset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AssetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Asset, error) {
	var (
		nodes       = []*Asset{}
		_spec       = aq.querySpec()
		loadedTypes = [5]bool{
			aq.withAssetPhoto != nil,
			aq.withAssetTransferLog != nil,
			aq.withDepartmentLocation != nil,
			aq.withEmployee != nil,
			aq.withAssetGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Asset).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Asset{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withAssetPhoto; query != nil {
		if err := aq.loadAssetPhoto(ctx, query, nodes,
			func(n *Asset) { n.Edges.AssetPhoto = []*AssetPhoto{} },
			func(n *Asset, e *AssetPhoto) { n.Edges.AssetPhoto = append(n.Edges.AssetPhoto, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAssetTransferLog; query != nil {
		if err := aq.loadAssetTransferLog(ctx, query, nodes,
			func(n *Asset) { n.Edges.AssetTransferLog = []*AssetTransferLog{} },
			func(n *Asset, e *AssetTransferLog) { n.Edges.AssetTransferLog = append(n.Edges.AssetTransferLog, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withDepartmentLocation; query != nil {
		if err := aq.loadDepartmentLocation(ctx, query, nodes, nil,
			func(n *Asset, e *DepartmentLocation) { n.Edges.DepartmentLocation = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withEmployee; query != nil {
		if err := aq.loadEmployee(ctx, query, nodes, nil,
			func(n *Asset, e *Employee) { n.Edges.Employee = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAssetGroup; query != nil {
		if err := aq.loadAssetGroup(ctx, query, nodes, nil,
			func(n *Asset, e *AssetGroup) { n.Edges.AssetGroup = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AssetQuery) loadAssetPhoto(ctx context.Context, query *AssetPhotoQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *AssetPhoto)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Asset)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AssetPhoto(func(s *sql.Selector) {
		s.Where(sql.InValues(asset.AssetPhotoColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.asset_asset_photo
		if fk == nil {
			return fmt.Errorf(`foreign-key "asset_asset_photo" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "asset_asset_photo" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AssetQuery) loadAssetTransferLog(ctx context.Context, query *AssetTransferLogQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *AssetTransferLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Asset)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.AssetTransferLog(func(s *sql.Selector) {
		s.Where(sql.InValues(asset.AssetTransferLogColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AssetID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "AssetID" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AssetQuery) loadDepartmentLocation(ctx context.Context, query *DepartmentLocationQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *DepartmentLocation)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Asset)
	for i := range nodes {
		fk := nodes[i].DepartmentLocationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(departmentlocation.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "DepartmentLocationID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AssetQuery) loadEmployee(ctx context.Context, query *EmployeeQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *Employee)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Asset)
	for i := range nodes {
		fk := nodes[i].EmployeeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(employee.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "EmployeeID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AssetQuery) loadAssetGroup(ctx context.Context, query *AssetGroupQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *AssetGroup)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Asset)
	for i := range nodes {
		fk := nodes[i].AssetGroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(assetgroup.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "AssetGroupID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AssetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AssetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(asset.Table, asset.Columns, sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asset.FieldID)
		for i := range fields {
			if fields[i] != asset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AssetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(asset.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = asset.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AssetGroupBy is the group-by builder for Asset entities.
type AssetGroupBy struct {
	selector
	build *AssetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AssetGroupBy) Aggregate(fns ...AggregateFunc) *AssetGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AssetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AssetQuery, *AssetGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AssetGroupBy) sqlScan(ctx context.Context, root *AssetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AssetSelect is the builder for selecting fields of Asset entities.
type AssetSelect struct {
	*AssetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AssetSelect) Aggregate(fns ...AggregateFunc) *AssetSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AssetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AssetQuery, *AssetSelect](ctx, as.AssetQuery, as, as.inters, v)
}

func (as *AssetSelect) sqlScan(ctx context.Context, root *AssetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
