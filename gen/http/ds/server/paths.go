// Code generated by goa v3.11.3, DO NOT EDIT.
//
// HTTP request path constructors for the ds service.
//
// Command:
// $ goa gen ds/design

package server

import (
	"fmt"
)

// ListDsPath returns the URL path to the ds service list HTTP endpoint.
func ListDsPath() string {
	return "/ds"
}

// CompleteDsPath returns the URL path to the ds service complete HTTP endpoint.
func CompleteDsPath(token string) string {
	return fmt.Sprintf("/ds/complete/%v", token)
}

// DemoDsPath returns the URL path to the ds service demo HTTP endpoint.
func DemoDsPath(a int, b int) string {
	return fmt.Sprintf("/ds/multiply/%v/%v", a, b)
}
