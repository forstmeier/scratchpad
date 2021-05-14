export const specimen = {
	externalID: {
		input: 'text',
		label: 'EXTERNAL ID',
		options: [],
		query: '',
		value: '',
	},
	description: {
		input: 'text',
		label: 'DESCRIPTION',
		options: [],
		query: '',
		value: '',
	},
	collectionDate: {
		input: 'date',
		label: 'COLLECTION DATE',
		options: [],
		query: '',
		value: '',
	},
	destructionDate: {
		input: 'date',
		label: 'DESTRUCTION DATE',
		options: [],
		query: '',
		value: '',
	},
	type: {
		input: 'select',
		label: 'TYPE',
		options: [
			{
				name: 'BLOOD',
				value: 'BLOOD',
			},
		],
		query: '',
		value: '',
	},
	container: {
		input: 'select',
		label: 'CONTAINER',
		options: [
			{
				name: 'VIAL',
				value: 'VIAL',
			},
		],
		query: '',
		value: '',
	},
	status: {
		input: 'select',
		label: 'STATUS',
		options: [
			{
				name: 'DESTROYED',
				value: 'DESTROYED',
			},
			{
				name: 'EXHAUSTED',
				value: 'EXHAUSTED',
			},
			{
				name: 'IN_INVENTORY',
				value: 'IN_INVENTORY',
			},
			{
				name: 'IN_TRANSIT',
				value: 'IN_TRANSIT',
			},
			{
				name: 'LOST',
				value: 'LOST',
			},
			{
				name: 'RESERVED',
				value: 'RESERVED',
			},
			{
				name: 'TRANSFERRED',
				value: 'TRANSFERRED',
			},
		],
		query: '',
		value: '',
	},
	donor: {
		input: 'select',
		label: 'DONOR',
		options: [],
		query: `queryDonor(filter: {}) {
			id
			dob
			sex
			race
	    }`,
		value: '',
	},
	consent: {
		input: 'select',
		label: 'CONSENT',
		options: [],
		query: `queryConsent(filter: {}) {
			id
			textBody
        }`,
		value: '',
	},
	tests: {
		input: 'multiple',
		label: 'TESTS',
		options: [],
		query: `queryTest(filter: {}) {
            id
        }`,
		value: [],
	},
	results: {
		input: 'multiple',
		label: 'RESULTS',
		options: [],
		query: `queryResult(filter: {}) {
            id
        }`,
		value: [],
	},
};
