// Code generated by ent, DO NOT EDIT.

package assetgroup

const (
	// Label holds the string label denoting the assetgroup type in the database.
	Label = "asset_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ID"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeAsset holds the string denoting the asset edge name in mutations.
	EdgeAsset = "Asset"
	// Table holds the table name of the assetgroup in the database.
	Table = "asset_groups"
	// AssetTable is the table that holds the Asset relation/edge.
	AssetTable = "assets"
	// AssetInverseTable is the table name for the Asset entity.
	// It exists in this package in order to avoid circular dependency with the "asset" package.
	AssetInverseTable = "assets"
	// AssetColumn is the table column denoting the Asset relation/edge.
	AssetColumn = "asset_group_id"
)

// Columns holds all SQL columns for assetgroup fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
