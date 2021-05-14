package main

func (dc *dgraphClient) addDonor(id string, count int) ([]string, error) {
	inputs := []map[string]interface{}{}
	for i := 0; i < count; i++ {
		input := map[string]interface{}{
			"org": map[string]string{
				"id": id,
			},
			"dob":  randomDate(),
			"sex":  randomString(sexes),
			"race": randomString(races),
		}

		inputs = append(inputs, input)
	}

	variables := map[string]interface{}{
		"input": inputs,
	}

	var response struct {
		Body struct {
			Data struct {
				Donor []struct {
					ID string `json:"id"`
				} `json:"donor"`
			} `json:"data"`
		} `json:"data"`
	}

	if err := dc.Mutate(
		addDonor,
		variables,
		&response,
	); err != nil {
		return nil, err
	}

	ids := []string{}
	for _, donor := range response.Body.Data.Donor {
		ids = append(ids, donor.ID)
	}

	return ids, nil
}
