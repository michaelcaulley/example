// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeTodos holds the string denoting the todos edge name in mutations.
	EdgeTodos = "todos"
	// EdgeModeratorUsers holds the string denoting the moderator_users edge name in mutations.
	EdgeModeratorUsers = "moderator_users"
	// EdgeModerators holds the string denoting the moderators edge name in mutations.
	EdgeModerators = "moderators"
	// EdgePeoplePartnerUsers holds the string denoting the people_partner_users edge name in mutations.
	EdgePeoplePartnerUsers = "people_partner_users"
	// EdgePeoplePartner holds the string denoting the people_partner edge name in mutations.
	EdgePeoplePartner = "people_partner"
	// EdgeModerator holds the string denoting the moderator edge name in mutations.
	EdgeModerator = "moderator"
	// EdgePeoplePartners holds the string denoting the people_partners edge name in mutations.
	EdgePeoplePartners = "people_partners"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TodosTable is the table that holds the todos relation/edge.
	TodosTable = "todos"
	// TodosInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodosInverseTable = "todos"
	// TodosColumn is the table column denoting the todos relation/edge.
	TodosColumn = "owner_id"
	// ModeratorUsersTable is the table that holds the moderator_users relation/edge. The primary key declared below.
	ModeratorUsersTable = "moderators"
	// ModeratorsTable is the table that holds the moderators relation/edge. The primary key declared below.
	ModeratorsTable = "moderators"
	// PeoplePartnerUsersTable is the table that holds the people_partner_users relation/edge. The primary key declared below.
	PeoplePartnerUsersTable = "people_partners"
	// PeoplePartnerTable is the table that holds the people_partner relation/edge. The primary key declared below.
	PeoplePartnerTable = "people_partners"
	// ModeratorTable is the table that holds the moderator relation/edge.
	ModeratorTable = "moderators"
	// ModeratorInverseTable is the table name for the Moderator entity.
	// It exists in this package in order to avoid circular dependency with the "moderator" package.
	ModeratorInverseTable = "moderators"
	// ModeratorColumn is the table column denoting the moderator relation/edge.
	ModeratorColumn = "user_id"
	// PeoplePartnersTable is the table that holds the people_partners relation/edge.
	PeoplePartnersTable = "people_partners"
	// PeoplePartnersInverseTable is the table name for the PeoplePartner entity.
	// It exists in this package in order to avoid circular dependency with the "peoplepartner" package.
	PeoplePartnersInverseTable = "people_partners"
	// PeoplePartnersColumn is the table column denoting the people_partners relation/edge.
	PeoplePartnersColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ModeratorUsersPrimaryKey and ModeratorUsersColumn2 are the table columns denoting the
	// primary key for the moderator_users relation (M2M).
	ModeratorUsersPrimaryKey = []string{"user_id", "moderator_user_id"}
	// ModeratorsPrimaryKey and ModeratorsColumn2 are the table columns denoting the
	// primary key for the moderators relation (M2M).
	ModeratorsPrimaryKey = []string{"user_id", "moderator_user_id"}
	// PeoplePartnerUsersPrimaryKey and PeoplePartnerUsersColumn2 are the table columns denoting the
	// primary key for the people_partner_users relation (M2M).
	PeoplePartnerUsersPrimaryKey = []string{"user_id", "people_partner_user_id"}
	// PeoplePartnerPrimaryKey and PeoplePartnerColumn2 are the table columns denoting the
	// primary key for the people_partner relation (M2M).
	PeoplePartnerPrimaryKey = []string{"user_id", "people_partner_user_id"}
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

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTodosCount orders the results by todos count.
func ByTodosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTodosStep(), opts...)
	}
}

// ByTodos orders the results by todos terms.
func ByTodos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByModeratorUsersCount orders the results by moderator_users count.
func ByModeratorUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newModeratorUsersStep(), opts...)
	}
}

// ByModeratorUsers orders the results by moderator_users terms.
func ByModeratorUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModeratorUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByModeratorsCount orders the results by moderators count.
func ByModeratorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newModeratorsStep(), opts...)
	}
}

// ByModerators orders the results by moderators terms.
func ByModerators(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModeratorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPeoplePartnerUsersCount orders the results by people_partner_users count.
func ByPeoplePartnerUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPeoplePartnerUsersStep(), opts...)
	}
}

// ByPeoplePartnerUsers orders the results by people_partner_users terms.
func ByPeoplePartnerUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPeoplePartnerUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPeoplePartnerCount orders the results by people_partner count.
func ByPeoplePartnerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPeoplePartnerStep(), opts...)
	}
}

// ByPeoplePartner orders the results by people_partner terms.
func ByPeoplePartner(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPeoplePartnerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByModeratorCount orders the results by moderator count.
func ByModeratorCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newModeratorStep(), opts...)
	}
}

// ByModerator orders the results by moderator terms.
func ByModerator(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModeratorStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPeoplePartnersCount orders the results by people_partners count.
func ByPeoplePartnersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPeoplePartnersStep(), opts...)
	}
}

// ByPeoplePartners orders the results by people_partners terms.
func ByPeoplePartners(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPeoplePartnersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTodosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TodosTable, TodosColumn),
	)
}
func newModeratorUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ModeratorUsersTable, ModeratorUsersPrimaryKey...),
	)
}
func newModeratorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ModeratorsTable, ModeratorsPrimaryKey...),
	)
}
func newPeoplePartnerUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PeoplePartnerUsersTable, PeoplePartnerUsersPrimaryKey...),
	)
}
func newPeoplePartnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PeoplePartnerTable, PeoplePartnerPrimaryKey...),
	)
}
func newModeratorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ModeratorInverseTable, ModeratorColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, ModeratorTable, ModeratorColumn),
	)
}
func newPeoplePartnersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PeoplePartnersInverseTable, PeoplePartnersColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, PeoplePartnersTable, PeoplePartnersColumn),
	)
}
