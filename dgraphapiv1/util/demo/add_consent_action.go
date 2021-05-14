package main

func (dc *dgraphClient) addConsentAction(orgID, donorID, consentID string) (string, error) {
	input := []map[string]interface{}{
		map[string]interface{}{
			"org": map[string]string{
				"id": orgID,
			},
			"donor": map[string]string{
				"id": donorID,
			},
			"consent": map[string]string{
				"id": consentID,
			},
			"timestamp": randomDate(),
		},
	}

	variables := map[string]interface{}{
		"input": input,
	}

	var response struct {
		Body struct {
			Data struct {
				DonorConsentAction []struct {
					ID string `json:"id"`
				} `json:"donorConsentAction"`
			} `json:"data"`
		} `json:"data"`
	}

	if err := dc.Mutate(
		addDonorConsentAction,
		variables,
		&response,
	); err != nil {
		return "", err
	}

	return response.Body.Data.DonorConsentAction[0].ID, nil
}
