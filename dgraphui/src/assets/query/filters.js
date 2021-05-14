export {
	specimen,
	consent,
	donor,
};

const _specimen = [
	{
		type: 'specimen',
		field: 'externalID',
		label: 'EXTERNAL ID',
		input: 'text',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'EQUALS',
				value: 'eq',
			},
		],
	},
	{
		type: 'specimen',
		field: 'description',
		label: 'DESCRIPTION',
		input: 'text',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'CONTAINS',
				value: 'anyoftext',
			},
		],
	},
	{
		type: 'specimen',
		field: 'collectionDate',
		label: 'COLLECTION DATE',
		input: 'date',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'LESS THAN',
				value: 'lt',
			},
			{
				name: 'EQUAL TO',
				value: 'eq',
			},
			{
				name: 'GREATER THAN',
				value: 'gt',
			},
		],
	},
	{
		type: 'specimen',
		field: 'destructionDate',
		label: 'DESTRUCTION DATE',
		input: 'date',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'LESS THAN',
				value: 'lt',
			},
			{
				name: 'EQUAL TO',
				value: 'eq',
			},
			{
				name: 'GREATER THAN',
				value: 'gt',
			},
		],
	},
	{
		type: 'specimen',
		field: 'type',
		label: 'TYPE',
		input: 'select',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'INCLUDES',
				value: 'regexp',
			},
		],
		options: [
			{
				name: 'BLOOD',
				value: 'BLOOD',
			},
		],
	},
	{
		type: 'specimen',
		field: 'container',
		label: 'CONTAINER',
		input: 'select',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'INCLUDES',
				value: 'regexp',
			},
		],
		options: [
			{
				name: 'VIAL',
				value: 'VIAL',
			},
		],
	},
	{
		type: 'specimen',
		field: 'status',
		label: 'STATUS',
		input: 'select',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'INCLUDES',
				value: 'regexp',
			},
		],
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
	},
];

const _bloodSpecimen = [
	{
		type: 'bloodSpecimen',
		field: 'bloodType',
		label: 'BLOOD TYPE',
		input: 'select',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'INCLUDES',
				value: 'regexp',
			},
		],
		options: [
			{
				name: 'O_NEG',
				value: 'O_NEG',
			},
			{
				name: 'O_POS',
				value: 'O_POS',
			},
			{
				name: 'A_NEG',
				value: 'A_NEG',
			},
			{
				name: 'A_POS',
				value: 'A_POS',
			},
			{
				name: 'B_NEG',
				value: 'B_NEG',
			},
			{
				name: 'B_POS',
				value: 'B_POS',
			},
			{
				name: 'AB_NEG',
				value: 'AB_NEG',
			},
			{
				name: 'AB_POS',
				value: 'AB_POS',
			},
		],
	},
	{
		type: 'bloodSpecimen',
		field: 'volume',
		label: 'VOLUME',
		input: 'number',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'LESS THAN',
				value: 'lt',
			},
			{
				name: 'EQUAL TO',
				value: 'eq',
			},
			{
				name: 'GREATER THAN',
				value: 'gt',
			},
		],
	},
];

const specimen = _specimen.concat(_bloodSpecimen);

const consent = [
	{
		type: 'consent',
		field: 'textBody',
		label: 'TEXT BODY',
		input: 'text',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'CONTAINS',
				value: 'anyoftext',
			},
		],
	},
	{
		type: 'consent',
		field: 'destructionDate',
		label: 'DESTRUCTION DATE',
		input: 'date',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'LESS THAN',
				value: 'lt',
			},
			{
				name: 'EQUAL TO',
				value: 'eq',
			},
			{
				name: 'GREATER THAN',
				value: 'gt',
			},
		],
	},
];

const donor = [
	{
		type: 'donor',
		field: 'dob',
		label: 'DOB',
		input: 'date',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'LESS THAN',
				value: 'lt',
			},
			{
				name: 'EQUAL TO',
				value: 'eq',
			},
			{
				name: 'GREATER THAN',
				value: 'gt',
			},
		],
	},
	{
		type: 'donor',
		field: 'sex',
		label: 'SEX',
		input: 'select',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'INCLUDES',
				value: 'regexp',
			},
		],
		options: [
			{
				name: 'MALE',
				value: 'MALE',
			},
			{
				name: 'FEMALE',
				value: 'FEMALE',
			},
		],
	},
	{
		type: 'donor',
		field: 'race',
		label: 'RACE',
		input: 'select',
		operators: [
			{
				name: '',
				value: '',
			},
			{
				name: 'INCLUDES',
				value: 'regexp',
			},
		],
		options: [
			{
				name: 'AMERICAN_INDIAN_OR_ALASKA_NATIVE',
				value: 'AMERICAN_INDIAN_OR_ALASKA_NATIVE',
			},
			{
				name: 'ASIAN',
				value: 'ASIAN',
			},
			{
				name: 'BLACK_OR_AFRICAN_AMERICAN',
				value: 'BLACK_OR_AFRICAN_AMERICAN',
			},
			{
				name: 'HISPANIC_OR_LATINO',
				value: 'HISPANIC_OR_LATINO',
			},
			{
				name: 'WHITE',
				value: 'WHITE',
			},
		],
	},
];