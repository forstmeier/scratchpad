package users

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

const updateUser = `mutation UpdateUser($input: UpdateUserInput!) {
	data: updateUser(input: $input) {
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

const deleteUser = `mutation DeleteUser($filter: UserFilter!) {
	data: deleteUser(filter: $filter) {
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
