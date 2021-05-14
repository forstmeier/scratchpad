import { specimen } from './specimen.js';


export const bloodSpecimen = {
	...specimen,
	bloodType: {
		input: 'select',
		label: 'BLOOD TYPE',
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
		query: '',
		value: '',
	},
	volume: {
		input: 'number',
		label: 'VOLUME',
		options: [],
		query: '',
		value: '',
	},
};
