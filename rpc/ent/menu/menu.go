// Code generated by ent, DO NOT EDIT.

package menu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldMenuLevel holds the string denoting the menu_level field in the database.
	FieldMenuLevel = "menu_level"
	// FieldMenuType holds the string denoting the menu_type field in the database.
	FieldMenuType = "menu_type"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRedirect holds the string denoting the redirect field in the database.
	FieldRedirect = "redirect"
	// FieldComponent holds the string denoting the component field in the database.
	FieldComponent = "component"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldServiceName holds the string denoting the service_name field in the database.
	FieldServiceName = "service_name"
	// FieldPermission holds the string denoting the permission field in the database.
	FieldPermission = "permission"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldHideMenu holds the string denoting the hide_menu field in the database.
	FieldHideMenu = "hide_menu"
	// FieldHideBreadcrumb holds the string denoting the hide_breadcrumb field in the database.
	FieldHideBreadcrumb = "hide_breadcrumb"
	// FieldIgnoreKeepAlive holds the string denoting the ignore_keep_alive field in the database.
	FieldIgnoreKeepAlive = "ignore_keep_alive"
	// FieldHideTab holds the string denoting the hide_tab field in the database.
	FieldHideTab = "hide_tab"
	// FieldFrameSrc holds the string denoting the frame_src field in the database.
	FieldFrameSrc = "frame_src"
	// FieldCarryParam holds the string denoting the carry_param field in the database.
	FieldCarryParam = "carry_param"
	// FieldHideChildrenInMenu holds the string denoting the hide_children_in_menu field in the database.
	FieldHideChildrenInMenu = "hide_children_in_menu"
	// FieldAffix holds the string denoting the affix field in the database.
	FieldAffix = "affix"
	// FieldDynamicLevel holds the string denoting the dynamic_level field in the database.
	FieldDynamicLevel = "dynamic_level"
	// FieldRealPath holds the string denoting the real_path field in the database.
	FieldRealPath = "real_path"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the menu in the database.
	Table = "sys_menus"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "role_menus"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "sys_roles"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_menus"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "sys_menus"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSort,
	FieldParentID,
	FieldMenuLevel,
	FieldMenuType,
	FieldPath,
	FieldName,
	FieldRedirect,
	FieldComponent,
	FieldDisabled,
	FieldServiceName,
	FieldPermission,
	FieldTitle,
	FieldIcon,
	FieldHideMenu,
	FieldHideBreadcrumb,
	FieldIgnoreKeepAlive,
	FieldHideTab,
	FieldFrameSrc,
	FieldCarryParam,
	FieldHideChildrenInMenu,
	FieldAffix,
	FieldDynamicLevel,
	FieldRealPath,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"role_id", "menu_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort uint32
	// DefaultParentID holds the default value on creation for the "parent_id" field.
	DefaultParentID uint64
	// DefaultPath holds the default value on creation for the "path" field.
	DefaultPath string
	// DefaultRedirect holds the default value on creation for the "redirect" field.
	DefaultRedirect string
	// DefaultComponent holds the default value on creation for the "component" field.
	DefaultComponent string
	// DefaultDisabled holds the default value on creation for the "disabled" field.
	DefaultDisabled bool
	// DefaultServiceName holds the default value on creation for the "service_name" field.
	DefaultServiceName string
	// DefaultHideMenu holds the default value on creation for the "hide_menu" field.
	DefaultHideMenu bool
	// DefaultHideBreadcrumb holds the default value on creation for the "hide_breadcrumb" field.
	DefaultHideBreadcrumb bool
	// DefaultIgnoreKeepAlive holds the default value on creation for the "ignore_keep_alive" field.
	DefaultIgnoreKeepAlive bool
	// DefaultHideTab holds the default value on creation for the "hide_tab" field.
	DefaultHideTab bool
	// DefaultFrameSrc holds the default value on creation for the "frame_src" field.
	DefaultFrameSrc string
	// DefaultCarryParam holds the default value on creation for the "carry_param" field.
	DefaultCarryParam bool
	// DefaultHideChildrenInMenu holds the default value on creation for the "hide_children_in_menu" field.
	DefaultHideChildrenInMenu bool
	// DefaultAffix holds the default value on creation for the "affix" field.
	DefaultAffix bool
	// DefaultDynamicLevel holds the default value on creation for the "dynamic_level" field.
	DefaultDynamicLevel uint32
	// DefaultRealPath holds the default value on creation for the "real_path" field.
	DefaultRealPath string
)

// OrderOption defines the ordering options for the Menu queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByMenuLevel orders the results by the menu_level field.
func ByMenuLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuLevel, opts...).ToFunc()
}

// ByMenuType orders the results by the menu_type field.
func ByMenuType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuType, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByRedirect orders the results by the redirect field.
func ByRedirect(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRedirect, opts...).ToFunc()
}

// ByComponent orders the results by the component field.
func ByComponent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComponent, opts...).ToFunc()
}

// ByDisabled orders the results by the disabled field.
func ByDisabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisabled, opts...).ToFunc()
}

// ByServiceName orders the results by the service_name field.
func ByServiceName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServiceName, opts...).ToFunc()
}

// ByPermission orders the results by the permission field.
func ByPermission(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPermission, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByIcon orders the results by the icon field.
func ByIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIcon, opts...).ToFunc()
}

// ByHideMenu orders the results by the hide_menu field.
func ByHideMenu(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHideMenu, opts...).ToFunc()
}

// ByHideBreadcrumb orders the results by the hide_breadcrumb field.
func ByHideBreadcrumb(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHideBreadcrumb, opts...).ToFunc()
}

// ByIgnoreKeepAlive orders the results by the ignore_keep_alive field.
func ByIgnoreKeepAlive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIgnoreKeepAlive, opts...).ToFunc()
}

// ByHideTab orders the results by the hide_tab field.
func ByHideTab(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHideTab, opts...).ToFunc()
}

// ByFrameSrc orders the results by the frame_src field.
func ByFrameSrc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrameSrc, opts...).ToFunc()
}

// ByCarryParam orders the results by the carry_param field.
func ByCarryParam(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCarryParam, opts...).ToFunc()
}

// ByHideChildrenInMenu orders the results by the hide_children_in_menu field.
func ByHideChildrenInMenu(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHideChildrenInMenu, opts...).ToFunc()
}

// ByAffix orders the results by the affix field.
func ByAffix(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAffix, opts...).ToFunc()
}

// ByDynamicLevel orders the results by the dynamic_level field.
func ByDynamicLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDynamicLevel, opts...).ToFunc()
}

// ByRealPath orders the results by the real_path field.
func ByRealPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRealPath, opts...).ToFunc()
}

// ByRolesCount orders the results by roles count.
func ByRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesStep(), opts...)
	}
}

// ByRoles orders the results by roles terms.
func ByRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
