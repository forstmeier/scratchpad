package specimens

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/folivoralabs/api/pkg/auth/tokens"
	"github.com/folivoralabs/api/pkg/graphql"
)

// Handler is an HTTP listener for BloodSpecimenUpdate type @custom directive events.
func Handler(dgraphURL string, client tokens.Tokens) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appToken, err := client.GetAppToken()
		if err != nil {
			http.Error(w, errorGettingAppToken, http.StatusBadRequest)
			return
		}

		var req AddBloodSpecimenUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, errorIncorrectRequestBody, http.StatusBadRequest)
			return
		}

		dgraphClient := graphql.New(
			dgraphURL,
			appToken,
		)

		timestamp := time.Now().Format(time.RFC3339)

		variables := map[string]interface{}{
			"input": map[string]interface{}{
				"orgID": map[string]string{
					"id": req.OrgID,
				},
				"bloodSpecimen": map[string]string{
					"id": req.BloodSpecimenID,
				},
				"bloodType":   req.BloodType,
				"container":   req.Container,
				"volume":      req.Volume,
				"description": req.Description,
				"timestamp":   timestamp,
			},
		}

		response := &Response{}

		if err := dgraphClient.Mutate(
			updateBloodSpecimen,
			variables,
			&response,
		); err != nil {
			http.Error(w, errorDgraphBloodSpecimenMutation, http.StatusInternalServerError)
			return
		}

		if err := dgraphClient.Mutate(
			addBloodSpecimenUpdate,
			variables,
			&response,
		); err != nil {
			http.Error(w, errorDgraphBloodSpecimenUpdateMutation, http.StatusInternalServerError)
			return
		}

		responseBody, err := json.Marshal(response.Body.Data.BloodSpecimenUpdate[0])
		if err != nil {
			http.Error(w, errorDgraphJSONMarshal, http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(responseBody))
	})
}
