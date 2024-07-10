//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Sessions = newSessionsTable("", "Sessions", "")

type sessionsTable struct {
	sqlite.Table

	// Columns
	UserID       sqlite.ColumnString
	AccessToken  sqlite.ColumnString
	RefreshToken sqlite.ColumnString
	TokenType    sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type SessionsTable struct {
	sessionsTable

	EXCLUDED sessionsTable
}

// AS creates new SessionsTable with assigned alias
func (a SessionsTable) AS(alias string) *SessionsTable {
	return newSessionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SessionsTable with assigned schema name
func (a SessionsTable) FromSchema(schemaName string) *SessionsTable {
	return newSessionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SessionsTable with assigned table prefix
func (a SessionsTable) WithPrefix(prefix string) *SessionsTable {
	return newSessionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SessionsTable with assigned table suffix
func (a SessionsTable) WithSuffix(suffix string) *SessionsTable {
	return newSessionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSessionsTable(schemaName, tableName, alias string) *SessionsTable {
	return &SessionsTable{
		sessionsTable: newSessionsTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newSessionsTableImpl("", "excluded", ""),
	}
}

func newSessionsTableImpl(schemaName, tableName, alias string) sessionsTable {
	var (
		UserIDColumn       = sqlite.StringColumn("UserID")
		AccessTokenColumn  = sqlite.StringColumn("AccessToken")
		RefreshTokenColumn = sqlite.StringColumn("RefreshToken")
		TokenTypeColumn    = sqlite.StringColumn("TokenType")
		allColumns         = sqlite.ColumnList{UserIDColumn, AccessTokenColumn, RefreshTokenColumn, TokenTypeColumn}
		mutableColumns     = sqlite.ColumnList{UserIDColumn, AccessTokenColumn, RefreshTokenColumn, TokenTypeColumn}
	)

	return sessionsTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:       UserIDColumn,
		AccessToken:  AccessTokenColumn,
		RefreshToken: RefreshTokenColumn,
		TokenType:    TokenTypeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
