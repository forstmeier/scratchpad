package user

// CreateUserRequest is a User type @custom directive requset from the
// custom server to Auth0 to create an Auth0 user.
type CreateUserRequest struct {
	Email       string      `json:"email" valid:"required"`
	Password    string      `json:"password" valid:"required"`
	AppMetadata AppMetadata `json:"app_metadata" valid:"required"`
	FirstName   string      `json:"given_name" valid:"required"`
	LastName    string      `json:"family_name" valid:"required"`
	Connection  string      `json:"connection" valid:"required"`
}

// CreateUserResponse is the object returned by the client following
// an action being taken in the Auth0 account.
//
// Auth0ID is unmarshalled from the larger response payload received
// from Auth0 after a management API users request is made.
type CreateUserResponse struct {
	Auth0ID string `json:"user_id"`
}

// UpdateUserRequest is a User type @custom directive requset from the
// custom server to Auth0 to update the Auth0 user.
type UpdateUserRequest struct {
	Password string `json:"password,omitempty"`
}

// AppMetadata is the app_metadata object on the user in Auth0.
type AppMetadata struct {
	OrgID string `json:"orgID,omitempty"`
}
