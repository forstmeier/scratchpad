package main

const addBloodSpecimenUpdate = `mutation AddBloodSpecimenUpdate($input: [AddBloodSpecimenUpdateInput!]!) {
	data: addBloodSpecimenUpdate(input: $input) {
		bloodSpecimenUpdate {
			id
		}
	}
}`

const addBloodSpecimen = `mutation AddBloodSpecimen($input: [AddBloodSpecimenInput!]!) {
	data: addBloodSpecimen(input: $input) {
		bloodSpecimen {
			id
		}
	}
}`

const addDonorConsentAction = `mutation AddDonorConsentAction($input: [AddDonorConsentActionInput!]!) {
	data: addDonorConsentAction(input: $input) {
		donorConsentAction {
			id
		}
	}
}`

const addConsent = `mutation AddConsent($input: [AddConsentInput!]!) {
	data: addConsent(input: $input) {
		consent {
			id
		}
	}
}`

const addDonor = `mutation AddDonor($input: [AddDonorInput!]!) {
	data: addDonor(input: $input) {
		donor {
			id
		}
	}
}`

const addOrg = `mutation AddOrg($input: [AddOrgInput!]!) {
	data: addOrg(input: $input) {
		org {
			id
		}
	}
}`

const addUser = `mutation AddUser($input: [AddUserInput!]!) {
	data: addUser(input: $input) {
		user {
			id
			org {
				id
			}
			email
			firstName
			lastName
			auth0ID
		}
	}
}`
