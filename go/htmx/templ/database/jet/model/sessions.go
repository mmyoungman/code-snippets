//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Sessions struct {
	ID           string `sql:"primary_key"`
	UserID       string
	AccessToken  string
	RefreshToken string
	TokenType    string
}
