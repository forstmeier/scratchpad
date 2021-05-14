package main

func (dc *dgraphClient) addBloodSpecimens(orgID, donorID, consentID string, count int) ([]string, string, error) {
	inputs := []map[string]interface{}{}
	bloodType := randomString(bloodTypes)
	for i := 0; i < count; i++ {
		input := map[string]interface{}{
			"org": map[string]string{
				"id": orgID,
			},
			"externalID":  randomID(),
			"description": randomString(quotes),
			"timestamp":   randomDate(),
			"type":        randomString(specimenTypes),
			"container":   randomString(containerTypes),
			"status":      randomString(specimenStatuses),
			"donor": map[string]string{
				"id": donorID,
			},
			"consent": map[string]string{
				"id": consentID,
			},
			"bloodType": bloodType,
			"volume":    1.00,
		}

		inputs = append(inputs, input)
	}

	variables := map[string]interface{}{
		"input": inputs,
	}

	var response struct {
		Body struct {
			Data struct {
				BloodSpecimen []struct {
					ID string `json:"id"`
				} `json:"bloodSpecimen"`
			} `json:"data"`
		} `json:"data"`
	}

	if err := dc.Mutate(
		addBloodSpecimen,
		variables,
		&response,
	); err != nil {
		return nil, "", err
	}

	ids := []string{}
	for _, bloodSpecimen := range response.Body.Data.BloodSpecimen {
		ids = append(ids, bloodSpecimen.ID)
	}

	return ids, bloodType, nil
}
