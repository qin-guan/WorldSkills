// Code generated by ent, DO NOT EDIT.

package department

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ID"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeDepartmentLocation holds the string denoting the departmentlocation edge name in mutations.
	EdgeDepartmentLocation = "DepartmentLocation"
	// Table holds the table name of the department in the database.
	Table = "departments"
	// DepartmentLocationTable is the table that holds the DepartmentLocation relation/edge.
	DepartmentLocationTable = "department_locations"
	// DepartmentLocationInverseTable is the table name for the DepartmentLocation entity.
	// It exists in this package in order to avoid circular dependency with the "departmentlocation" package.
	DepartmentLocationInverseTable = "department_locations"
	// DepartmentLocationColumn is the table column denoting the DepartmentLocation relation/edge.
	DepartmentLocationColumn = "department_id"
)

// Columns holds all SQL columns for department fields.
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
