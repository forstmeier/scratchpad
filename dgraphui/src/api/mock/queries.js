export {
	queryBloodSpecimen,
	bloodSpecimenPreQuery,
}

const queryBloodSpecimen = {
	data: {
		data: {
			queryBloodSpecimen: [
				{
					id: 'specimen_0',
					status: 'IN_INVENTORY',
					externalID: 'external_id_A',
					description: 'So uncivilized',
					collectionDate: '2020-11-15',
					type: 'BLOOD',
					container: 'VIAL',
					donor: {
						dob: '1992-04-16',
						sex: 'MALE',
						race: 'WHITE',
					},
					consent: {
						textBody: 'May the Force be with you.',
					},
					bloodType: 'A_NEG',
					volume: 10,
				},
				{
					id: 'specimen_1',
					status: 'IN_TRANSIT',
					externalID: 'external_id_B',
					description: "Of course I know him. He's me.",
					collectionDate: '2020-11-15',
					type: 'BLOOD',
					container: 'VIAL',
					donor: {
						dob: '1992-04-16',
						sex: 'MALE',
						race: 'WHITE',
					},
					consent: {
						textBody: 'May the Force be with you.',
					},
					bloodType: 'A_NEG',
					volume: 10,
				},
				{
					id: 'specimen_2',
					status: 'IN_INVENTORY',
					externalID: 'external_id_C',
					description: 'Hello, there!',
					collectionDate: '2020-11-15',
					type: 'BLOOD',
					container: 'VIAL',
					donor: {
						dob: '1991-10-20',
						sex: 'FEMALE',
						race: 'WHITE',
					},
					consent: {
						textBody: 'May the Force be with you.',
					},
					bloodType: 'O_POS',
					volume: 10,
				},
			],
		},
	},
};

const bloodSpecimenPreQuery = {
	data: {
		data: {
			donor: [
				{
					id: 'donor_0',
					dob: '1991-10-20',
					sex: 'FEMALE',
					race: 'WHITE',
				},
				{
					id: 'donor_1',
					dob: '1992-04-16',
					sex: 'MALE',
					race: 'WHITE',
				},
			],
			consent: [
				{
					id: 'consent_0',
					textBody: 'May the Force be with you.',
				},
			],
			tests: [
				{
					'id': 'test_0'
				},
				{
					'id': 'test_1'
				},
			],
			results: [
				{
					'id': 'result_0'
				},
				{
					'id': 'result_1'
				},
			],
		},
	},
};