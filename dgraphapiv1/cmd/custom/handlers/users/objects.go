package users

// Auth0CreateUserRequest is a request received from the @custom directive
// from Dgraph sent to the custom server to create an Auth0 user.
type Auth0CreateUserRequest struct {
	OrgID     string `json:"orgID" valid:"required"`
	Email     string `json:"email" valid:"required"`
	Password  string `json:"password" valid:"required"`
	FirstName string `json:"firstName" valid:"required"`
	LastName  string `json:"lastName" valid:"required"`
}

// Auth0UpdateUserRequest is a request received from the @custom directive
// from Dgraph sent to the custom server to update an Auth0 user.
type Auth0UpdateUserRequest struct {
	Auth0ID  string `json:"authZeroID" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// Auth0DeleteUserRequest is a request received from the @custom directive
// from Dgraph sent to the custom server to delete an Auth0 user.
type Auth0DeleteUserRequest struct {
	Auth0ID string `json:"authZeroID" valid:"required"`
}

// Response holds the HTTP response body from Dgraph.
type Response struct {
	Body Body `json:"data"`
}

// Body holds the Dgraph response data.
type Body struct {
	Data Data `json:"data"`
}

// Data holds the Dgraph response User(s) data.
type Data struct {
	User []User `json:"user"`
}

// User holds the Dgraph response User data.
type User struct {
	ID        string `json:"id"`
	Org       Org    `json:"org"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Auth0ID   string `json:"auth0ID"`
}

// Org holds the Dgraph response Org data.
type Org struct {
	ID string `json:"id"`
}
