import axios from 'axios';
import {
	queryBloodSpecimen,
	bloodSpecimenPreQuery,
} from './mock/queries.js';
import { addBloodSpecimen } from './mock/mutations.js';


export default class API {
	constructor(url, token, secret) {
		this.url = url;
		this.token = token;
		this.secret = secret;
	}

	sendRequest(query, variables, callback) {
		if (process.env.VUE_APP_FOLIVORA_ENV === 'mock') {
			if (query.includes('QueryBloodSpecimen')) {
				callback(queryBloodSpecimen);
			} else if (query.includes('AddBloodSpecimen')) {
				callback(addBloodSpecimen);
			} else if (query.includes('BloodSpecimenPreQuery')) {
				callback(bloodSpecimenPreQuery);
			}
		} else {
			const options = {
				method: 'post',
				url: this.url,
				headers: {
					'X-Auth0-Token': this.token,
					'folivora-custom-secret': this.secret,
				},
				data: {
					query: query,
					variables: variables,
				},
			};

			axios(options)
				.then(resp => {
					callback(resp);
				})
				.catch(err => {

					// outline:
					// [ ] add correct error callback handling

					console.log('err:', err);
				});
		}
	}
}

// NOTES:
// [ ] https://vuejs.org/v2/cookbook/using-axios-to-consume-apis.html
// [ ] https://auth0.com/docs/quickstart/spa/vuejs/01-login
// [ ] https://auth0.com/docs/quickstart/spa/vuejs/02-calling-an-api#call-the-api-using-an-access-token
// [ ] https://cli.vuejs.org/guide/mode-and-env.html
