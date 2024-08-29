// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthmonitor/ent/adminpermission"
	"healthmonitor/ent/adminrolepermission"
	"healthmonitor/ent/adminroles"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AdminRolePermission is the model entity for the AdminRolePermission schema.
type AdminRolePermission struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdminRolePermissionQuery when eager-loading is set.
	Edges                   AdminRolePermissionEdges `json:"edges"`
	admin_permission_roles  *uuid.UUID
	admin_roles_permissions *uuid.UUID
	selectValues            sql.SelectValues
}

// AdminRolePermissionEdges holds the relations/edges for other nodes in the graph.
type AdminRolePermissionEdges struct {
	// Role holds the value of the role edge.
	Role *AdminRoles `json:"role,omitempty"`
	// Permission holds the value of the permission edge.
	Permission *AdminPermission `json:"permission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdminRolePermissionEdges) RoleOrErr() (*AdminRoles, error) {
	if e.Role != nil {
		return e.Role, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: adminroles.Label}
	}
	return nil, &NotLoadedError{edge: "role"}
}

// PermissionOrErr returns the Permission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdminRolePermissionEdges) PermissionOrErr() (*AdminPermission, error) {
	if e.Permission != nil {
		return e.Permission, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: adminpermission.Label}
	}
	return nil, &NotLoadedError{edge: "permission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AdminRolePermission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case adminrolepermission.FieldID:
			values[i] = new(sql.NullInt64)
		case adminrolepermission.ForeignKeys[0]: // admin_permission_roles
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case adminrolepermission.ForeignKeys[1]: // admin_roles_permissions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdminRolePermission fields.
func (arp *AdminRolePermission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case adminrolepermission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			arp.ID = int(value.Int64)
		case adminrolepermission.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field admin_permission_roles", values[i])
			} else if value.Valid {
				arp.admin_permission_roles = new(uuid.UUID)
				*arp.admin_permission_roles = *value.S.(*uuid.UUID)
			}
		case adminrolepermission.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field admin_roles_permissions", values[i])
			} else if value.Valid {
				arp.admin_roles_permissions = new(uuid.UUID)
				*arp.admin_roles_permissions = *value.S.(*uuid.UUID)
			}
		default:
			arp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AdminRolePermission.
// This includes values selected through modifiers, order, etc.
func (arp *AdminRolePermission) Value(name string) (ent.Value, error) {
	return arp.selectValues.Get(name)
}

// QueryRole queries the "role" edge of the AdminRolePermission entity.
func (arp *AdminRolePermission) QueryRole() *AdminRolesQuery {
	return NewAdminRolePermissionClient(arp.config).QueryRole(arp)
}

// QueryPermission queries the "permission" edge of the AdminRolePermission entity.
func (arp *AdminRolePermission) QueryPermission() *AdminPermissionQuery {
	return NewAdminRolePermissionClient(arp.config).QueryPermission(arp)
}

// Update returns a builder for updating this AdminRolePermission.
// Note that you need to call AdminRolePermission.Unwrap() before calling this method if this AdminRolePermission
// was returned from a transaction, and the transaction was committed or rolled back.
func (arp *AdminRolePermission) Update() *AdminRolePermissionUpdateOne {
	return NewAdminRolePermissionClient(arp.config).UpdateOne(arp)
}

// Unwrap unwraps the AdminRolePermission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (arp *AdminRolePermission) Unwrap() *AdminRolePermission {
	_tx, ok := arp.config.driver.(*txDriver)
	if !ok {
		panic("ent: AdminRolePermission is not a transactional entity")
	}
	arp.config.driver = _tx.drv
	return arp
}

// String implements the fmt.Stringer.
func (arp *AdminRolePermission) String() string {
	var builder strings.Builder
	builder.WriteString("AdminRolePermission(")
	builder.WriteString(fmt.Sprintf("id=%v", arp.ID))
	builder.WriteByte(')')
	return builder.String()
}

// AdminRolePermissions is a parsable slice of AdminRolePermission.
type AdminRolePermissions []*AdminRolePermission
