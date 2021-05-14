package main

import "fmt"

func (dc *dgraphClient) addUsers(id string, index int) ([]string, error) {
	inputs := []map[string]interface{}{}
	for _, user := range demoUsers {
		input := map[string]interface{}{
			"email": fmt.Sprintf("%s.%s@%s.org", user.firstName, user.lastName, demoOrgs[index]),
			"org": map[string]string{
				"id": id,
			},
			"firstName": user.firstName,
			"lastName":  user.lastName,
			// auth0ID field is being populated since this is a "direct"
			// invocation of the generated "add" mutation
			"auth0ID": fmt.Sprintf("auth0|%s", randomID()),
		}

		inputs = append(inputs, input)
	}

	variables := map[string]interface{}{
		"input": inputs,
	}

	var response struct {
		Body struct {
			Data struct {
				User []struct {
					ID string `json:"id"`
				} `json:"user"`
			} `json:"data"`
		} `json:"data"`
	}

	if err := dc.Mutate(
		addUser,
		variables,
		&response,
	); err != nil {
		return nil, err
	}

	ids := []string{}
	for _, user := range response.Body.Data.User {
		ids = append(ids, user.ID)
	}

	return ids, nil
}
