package users

func createUserDgraphReq(req Auth0CreateUserRequest, auth0ID string) map[string]interface{} {
	return map[string]interface{}{
		"input": []map[string]interface{}{
			map[string]interface{}{
				"org": map[string]string{
					"id": req.OrgID,
				},
				"email":     req.Email,
				"firstName": req.FirstName,
				"lastName":  req.LastName,
				"auth0ID":   auth0ID,
			},
		},
	}
}

func updateUserDgraphReq(req Auth0UpdateUserRequest) map[string]interface{} {
	return map[string]interface{}{
		"input": map[string]interface{}{
			"filter": map[string]interface{}{
				"auth0ID": map[string]string{
					"eq": req.Auth0ID,
				},
			},
			// there is not anything to update currently
			"set": map[string]string{
				"auth0ID": req.Auth0ID,
			},
		},
	}
}
