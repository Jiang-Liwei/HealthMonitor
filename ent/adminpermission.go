// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthmonitor/ent/adminpermission"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AdminPermission is the model entity for the AdminPermission schema.
type AdminPermission struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt int `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdminPermissionQuery when eager-loading is set.
	Edges        AdminPermissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AdminPermissionEdges holds the relations/edges for other nodes in the graph.
type AdminPermissionEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*AdminRolePermission `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e AdminPermissionEdges) RolesOrErr() ([]*AdminRolePermission, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AdminPermission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case adminpermission.FieldCreatedAt, adminpermission.FieldUpdatedAt, adminpermission.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case adminpermission.FieldName, adminpermission.FieldDescription, adminpermission.FieldPath, adminpermission.FieldMethod:
			values[i] = new(sql.NullString)
		case adminpermission.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdminPermission fields.
func (ap *AdminPermission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case adminpermission.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ap.ID = *value
			}
		case adminpermission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ap.Name = value.String
			}
		case adminpermission.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ap.Description = value.String
			}
		case adminpermission.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				ap.Path = value.String
			}
		case adminpermission.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				ap.Method = value.String
			}
		case adminpermission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ap.CreatedAt = int(value.Int64)
			}
		case adminpermission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ap.UpdatedAt = int(value.Int64)
			}
		case adminpermission.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ap.DeletedAt = int(value.Int64)
			}
		default:
			ap.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AdminPermission.
// This includes values selected through modifiers, order, etc.
func (ap *AdminPermission) Value(name string) (ent.Value, error) {
	return ap.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the AdminPermission entity.
func (ap *AdminPermission) QueryRoles() *AdminRolePermissionQuery {
	return NewAdminPermissionClient(ap.config).QueryRoles(ap)
}

// Update returns a builder for updating this AdminPermission.
// Note that you need to call AdminPermission.Unwrap() before calling this method if this AdminPermission
// was returned from a transaction, and the transaction was committed or rolled back.
func (ap *AdminPermission) Update() *AdminPermissionUpdateOne {
	return NewAdminPermissionClient(ap.config).UpdateOne(ap)
}

// Unwrap unwraps the AdminPermission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ap *AdminPermission) Unwrap() *AdminPermission {
	_tx, ok := ap.config.driver.(*txDriver)
	if !ok {
		panic("ent: AdminPermission is not a transactional entity")
	}
	ap.config.driver = _tx.drv
	return ap
}

// String implements the fmt.Stringer.
func (ap *AdminPermission) String() string {
	var builder strings.Builder
	builder.WriteString("AdminPermission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ap.ID))
	builder.WriteString("name=")
	builder.WriteString(ap.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ap.Description)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(ap.Path)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(ap.Method)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ap.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ap.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ap.DeletedAt))
	builder.WriteByte(')')
	return builder.String()
}

// AdminPermissions is a parsable slice of AdminPermission.
type AdminPermissions []*AdminPermission
