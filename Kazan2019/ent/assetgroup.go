// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assetgroup"
)

// AssetGroup is the model entity for the AssetGroup schema.
type AssetGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AssetGroupQuery when eager-loading is set.
	Edges AssetGroupEdges `json:"edges"`
}

// AssetGroupEdges holds the relations/edges for other nodes in the graph.
type AssetGroupEdges struct {
	// Asset holds the value of the Asset edge.
	Asset []*Asset `json:"Asset,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AssetOrErr returns the Asset value or an error if the edge
// was not loaded in eager-loading.
func (e AssetGroupEdges) AssetOrErr() ([]*Asset, error) {
	if e.loadedTypes[0] {
		return e.Asset, nil
	}
	return nil, &NotLoadedError{edge: "Asset"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AssetGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case assetgroup.FieldID:
			values[i] = new(sql.NullInt64)
		case assetgroup.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AssetGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AssetGroup fields.
func (ag *AssetGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case assetgroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ag.ID = int(value.Int64)
		case assetgroup.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				ag.Name = value.String
			}
		}
	}
	return nil
}

// QueryAsset queries the "Asset" edge of the AssetGroup entity.
func (ag *AssetGroup) QueryAsset() *AssetQuery {
	return NewAssetGroupClient(ag.config).QueryAsset(ag)
}

// Update returns a builder for updating this AssetGroup.
// Note that you need to call AssetGroup.Unwrap() before calling this method if this AssetGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (ag *AssetGroup) Update() *AssetGroupUpdateOne {
	return NewAssetGroupClient(ag.config).UpdateOne(ag)
}

// Unwrap unwraps the AssetGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ag *AssetGroup) Unwrap() *AssetGroup {
	_tx, ok := ag.config.driver.(*txDriver)
	if !ok {
		panic("ent: AssetGroup is not a transactional entity")
	}
	ag.config.driver = _tx.drv
	return ag
}

// String implements the fmt.Stringer.
func (ag *AssetGroup) String() string {
	var builder strings.Builder
	builder.WriteString("AssetGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ag.ID))
	builder.WriteString("Name=")
	builder.WriteString(ag.Name)
	builder.WriteByte(')')
	return builder.String()
}

// AssetGroups is a parsable slice of AssetGroup.
type AssetGroups []*AssetGroup
