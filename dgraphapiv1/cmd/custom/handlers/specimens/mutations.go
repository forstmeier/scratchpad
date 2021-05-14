package specimens

const updateBloodSpecimen = `mutation UpdateBloodSpecimens($input: [UpdateBloodSpecimenInput!]!) {
	data: updateBloodSpecimen(input: $input) {
		bloodSpecimen {
			id
		}
	}
}`

const addBloodSpecimenUpdate = `mutation AddBloodSpecimenUpdates($input: [AddBloodSpecimenUpdateInput!]!) {
	data: addBloodSpecimenUpdate(input: $input) {
		bloodSpecimenUpdate {
			id
		}
	}
}`
