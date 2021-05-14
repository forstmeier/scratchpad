package main

func (dc *dgraphClient) addConsents(id string, count int) ([]string, error) {
	inputs := []map[string]interface{}{}
	for i := 0; i < count; i++ {
		input := map[string]interface{}{
			"org": map[string]string{
				"id": id,
			},
			"textBody":        randomString(quotes),
			"destructionDate": randomDate(),
		}

		inputs = append(inputs, input)
	}

	variables := map[string]interface{}{
		"input": inputs,
	}

	var response struct {
		Body struct {
			Data struct {
				Consent []struct {
					ID string `json:"id"`
				} `json:"consent"`
			} `json:"data"`
		} `json:"data"`
	}

	if err := dc.Mutate(
		addConsent,
		variables,
		&response,
	); err != nil {
		return nil, err
	}

	ids := []string{}
	for _, consent := range response.Body.Data.Consent {
		ids = append(ids, consent.ID)
	}

	return ids, nil
}
