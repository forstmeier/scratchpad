package main

import "time"

func (dc *dgraphClient) addBloodSpecimenUpdates(orgID, bloodSpecimenID, bloodType string, count int) ([]string, error) {
	inputs := []map[string]interface{}{}
	for i := 0; i < count; i++ {

		if bloodType == "" {
			bloodType = randomString(bloodTypes)
		}

		input := map[string]interface{}{
			"org": map[string]string{
				"id": orgID,
			},
			"bloodSpecimen": map[string]string{
				"id": bloodSpecimenID,
			},
			"bloodType": bloodType,
			"timestamp": time.Now().Format(time.RFC3339),
		}

		inputs = append(inputs, input)
	}

	variables := map[string]interface{}{
		"input": inputs,
	}

	var response struct {
		Body struct {
			Data struct {
				BloodSpecimenUpdate []struct {
					ID string `json:"id"`
				} `json:"bloodSpecimenUpdate"`
			} `json:"data"`
		} `json:"data"`
	}

	if err := dc.Mutate(
		addBloodSpecimenUpdate,
		variables,
		&response,
	); err != nil {
		return nil, err
	}

	ids := []string{}
	for _, bloodSpecimenUpdate := range response.Body.Data.BloodSpecimenUpdate {
		ids = append(ids, bloodSpecimenUpdate.ID)
	}

	return ids, nil
}
