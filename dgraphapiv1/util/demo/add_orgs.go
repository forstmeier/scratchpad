package main

func (dc *dgraphClient) addOrgs(names []string) ([]string, error) {
	inputs := []map[string]string{}
	for _, name := range names {
		input := map[string]string{
			"name": name,
		}

		inputs = append(inputs, input)
	}

	variables := map[string]interface{}{
		"input": inputs,
	}

	var response struct {
		Body struct {
			Data struct {
				Org []struct {
					ID string `json:"id"`
				} `json:"org"`
			} `json:"data"`
		} `json:"data"`
	}

	if err := dc.Mutate(
		addOrg,
		variables,
		&response,
	); err != nil {
		return nil, err
	}

	ids := []string{}
	for _, org := range response.Body.Data.Org {
		ids = append(ids, org.ID)
	}

	return ids, nil
}
