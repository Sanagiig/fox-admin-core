// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ExampleRpcsColumns holds the columns for the "example_rpcs" table.
	ExampleRpcsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ExampleRpcsTable holds the schema information for the "example_rpcs" table.
	ExampleRpcsTable = &schema.Table{
		Name:       "example_rpcs",
		Columns:    ExampleRpcsColumns,
		PrimaryKey: []*schema.Column{ExampleRpcsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ExampleRpcsTable,
	}
)

func init() {
}
