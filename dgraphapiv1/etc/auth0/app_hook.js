// NOTE: this hook is used for M2M JWT generation for
// programmatic access to the Dgraph API.
// 
// This hook only applies to M2M authentication flows.

module.exports = function (client, scope, audience, context, cb) {
	var access_token = {};
	access_token.scope = scope; // keep this required line

	access_token['https://folivora.io/jwt/claims'] = {
		isApp: 'true',
		isAuthenticated: 'true',
	};
	access_token.scope.push('extra');

	cb(null, access_token);
};