// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Friends []*User `json:"friends"`
}
