package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/cheggaaa/pb/v3"

	"github.com/folivoralabs/api/pkg/auth/tokens"
	"github.com/folivoralabs/api/pkg/config"
	"github.com/folivoralabs/api/pkg/graphql"
)

type dgraphClient struct {
	graphql.GraphQL
}

func loadDemo(cfg *config.Config) error {
	results := make(map[string]map[string]map[string][]string)
	results["orgs"] = make(map[string]map[string][]string)

	cfg, err := config.New("../../etc/config/config.json")
	if err != nil {
		log.Fatalf("error reading config file: %s", err.Error())
	}

	tokensClient := tokens.New(cfg)

	appToken, err := tokensClient.GetAppToken()
	if err != nil {
		log.Fatalf("get management token error: %s", err.Error())
	}

	typesCount := 7 // this should reflect the number of "add*" methods
	bar := pb.StartNew(len(demoOrgs) * typesCount)

	dgraphAppClient := &dgraphClient{
		graphql.New(
			"http://localhost:8080/graphql",
			appToken,
		),
	}

	orgIDs, err := dgraphAppClient.addOrgs(demoOrgs)
	if err != nil {
		log.Fatalf("add orgs error: %s", err.Error())
	}
	for i := 0; i < len(orgIDs); i++ {
		bar.Increment()
	}

	for orgIndex, orgID := range orgIDs {
		results["orgs"][orgID] = make(map[string][]string)

		if err := tokensClient.UpdateUserToken("TEST_FORSTMEIER", orgID, appToken); err != nil {
			log.Fatalf("error updating user token: %s", err.Error())
		}

		userToken, err := tokensClient.GetUserToken("TEST_FORSTMEIER")
		if err != nil {
			log.Fatalf("error getting user token: %s", err.Error())
		}

		dgraphUserClient := &dgraphClient{
			graphql.New(
				"http://localhost:8080/graphql",
				userToken,
			),
		}

		userIDs, err := dgraphAppClient.addUsers(orgID, orgIndex)
		if err != nil {
			log.Fatalf("add users error: %s", err.Error())
		}
		bar.Increment()
		results["orgs"][orgID]["users"] = userIDs

		donorCount := rand.Intn(50)
		donorIDs, err := dgraphUserClient.addDonor(orgID, donorCount)
		if err != nil {
			log.Fatalf("add donors error: %s", err.Error())
		}
		bar.Increment()
		results["orgs"][orgID]["donors"] = donorIDs

		consentCount := rand.Intn(5) + 1
		consentIDs, err := dgraphUserClient.addConsents(orgID, consentCount)
		if err != nil {
			log.Fatalf("add consents error: %s", err.Error())
		}
		bar.Increment()
		results["orgs"][orgID]["consents"] = consentIDs

		consentActionIDs := []string{}
		allBloodSpecimenIDs := []string{}
		allBloodSpecimenUpdateIDs := []string{}
		for _, donorID := range donorIDs {
			consentID := consentIDs[rand.Intn(consentCount)]
			consentActionID, err := dgraphUserClient.addConsentAction(orgID, donorID, consentID)
			if err != nil {
				log.Fatalf("add consent action error: %s", err.Error())
			}

			consentActionIDs = append(consentActionIDs, consentActionID)

			specimenCount := rand.Intn(10) + 2 // needs to be 2 since we want at least two updates
			bloodSpecimenIDs, bloodType, err := dgraphAppClient.addBloodSpecimens(orgID, donorID, consentID, specimenCount)
			if err != nil {
				log.Fatalf("add blood specimens error: %s", err.Error())
			}

			for _, bloodSpecimenID := range bloodSpecimenIDs {
				bloodSpecimenUpdateIDs, err := dgraphAppClient.addBloodSpecimenUpdates(orgID, bloodSpecimenID, "", specimenCount-1)
				if err != nil {
					log.Fatalf("add blood specimen updates error: %s", err.Error())
				}
				allBloodSpecimenUpdateIDs = append(allBloodSpecimenUpdateIDs, bloodSpecimenUpdateIDs...)

				// this is to ensure that the same blood type is set on the final update as is set on the original blood specimen
				bloodSpecimenUpdateID, err := dgraphAppClient.addBloodSpecimenUpdates(orgID, bloodSpecimenID, bloodType, 1)
				if err != nil {
					log.Fatalf("add blood specimen update error: %s", err.Error())
				}

				allBloodSpecimenIDs = append(allBloodSpecimenUpdateIDs, bloodSpecimenUpdateID[0])
			}

			allBloodSpecimenIDs = append(allBloodSpecimenIDs, bloodSpecimenIDs...)
		}
		results["orgs"][orgID]["consentActions"] = consentActionIDs
		results["orgs"][orgID]["bloodSpecimens"] = allBloodSpecimenIDs
		results["orgs"][orgID]["bloodSpecimensUpdates"] = allBloodSpecimenIDs
		bar.Increment()
		bar.Increment()
		bar.Increment()
	}

	bar.Finish()

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Fatalf("error marshalling results data: %s", err.Error())
	}

	if err := ioutil.WriteFile("results.json", jsonData, 0644); err != nil {
		log.Fatalf("error writing results to file: %s", err.Error())
	}

	return nil
}
