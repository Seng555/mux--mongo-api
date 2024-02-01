// Parse the request body into credentials
// model/response.go
package model

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
