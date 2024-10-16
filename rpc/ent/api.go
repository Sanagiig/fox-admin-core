// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Sanagiig/fox-admin-core/rpc/ent/api"
)

// Api is the model entity for the Api schema.
type Api struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// API path | API 路径
	Path string `json:"path,omitempty"`
	// API description | API 描述
	Description string `json:"description,omitempty"`
	// API group | API 分组
	APIGroup string `json:"api_group,omitempty"`
	// Service name | 服务名称
	ServiceName string `json:"service_name,omitempty"`
	// HTTP method | HTTP 请求类型
	Method string `json:"method,omitempty"`
	// Whether is required | 是否必选
	IsRequired   bool `json:"is_required,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Api) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case api.FieldIsRequired:
			values[i] = new(sql.NullBool)
		case api.FieldID:
			values[i] = new(sql.NullInt64)
		case api.FieldPath, api.FieldDescription, api.FieldAPIGroup, api.FieldServiceName, api.FieldMethod:
			values[i] = new(sql.NullString)
		case api.FieldCreatedAt, api.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Api fields.
func (a *Api) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case api.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case api.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case api.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case api.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				a.Path = value.String
			}
		case api.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case api.FieldAPIGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_group", values[i])
			} else if value.Valid {
				a.APIGroup = value.String
			}
		case api.FieldServiceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service_name", values[i])
			} else if value.Valid {
				a.ServiceName = value.String
			}
		case api.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				a.Method = value.String
			}
		case api.FieldIsRequired:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_required", values[i])
			} else if value.Valid {
				a.IsRequired = value.Bool
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Api.
// This includes values selected through modifiers, order, etc.
func (a *Api) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Api.
// Note that you need to call Api.Unwrap() before calling this method if this Api
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Api) Update() *APIUpdateOne {
	return NewAPIClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Api entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Api) Unwrap() *Api {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Api is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Api) String() string {
	var builder strings.Builder
	builder.WriteString("Api(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(a.Path)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("api_group=")
	builder.WriteString(a.APIGroup)
	builder.WriteString(", ")
	builder.WriteString("service_name=")
	builder.WriteString(a.ServiceName)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(a.Method)
	builder.WriteString(", ")
	builder.WriteString("is_required=")
	builder.WriteString(fmt.Sprintf("%v", a.IsRequired))
	builder.WriteByte(')')
	return builder.String()
}

// Apis is a parsable slice of Api.
type Apis []*Api
