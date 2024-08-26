// Code generated by ent, DO NOT EDIT.

package adminuserrole

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the adminuserrole type in the database.
	Label = "admin_user_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the adminuserrole in the database.
	Table = "admin_user_roles"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "admin_user_roles"
	// UserInverseTable is the table name for the AdminUser entity.
	// It exists in this package in order to avoid circular dependency with the "adminuser" package.
	UserInverseTable = "admin_users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "admin_user_roles"
	// RoleTable is the table that holds the role relation/edge.
	RoleTable = "admin_user_roles"
	// RoleInverseTable is the table name for the AdminRoles entity.
	// It exists in this package in order to avoid circular dependency with the "adminroles" package.
	RoleInverseTable = "admin_roles"
	// RoleColumn is the table column denoting the role relation/edge.
	RoleColumn = "admin_roles_users"
)

// Columns holds all SQL columns for adminuserrole fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "admin_user_roles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"admin_roles_users",
	"admin_user_roles",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the AdminUserRole queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByRoleField orders the results by role field.
func ByRoleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RoleTable, RoleColumn),
	)
}