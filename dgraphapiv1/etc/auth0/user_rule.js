// NOTE: this rule is used in conjunction with the @custom
// directives that populate the values in the app_metadata
// and are then added to the claims on login.
// 
// This rule only applies to user authentication flows.

function addAttributes(user, context, callback) {
  const claims = {
    isAuthenticated: 'true', // string because of dgraph requirement
    orgID: user.app_metadata.orgID,
    isApp: 'false', // string because of dgraph requirement
  };

  context.idToken['https://folivora.io/jwt/claims'] = claims;
  callback(null, user, context);
}