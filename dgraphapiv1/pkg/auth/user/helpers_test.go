package user

import "testing"

func Test_checkCreateUserValues(t *testing.T) {
	tests := []struct {
		inputs   []string
		check    bool
		response string
	}{
		{
			inputs:   []string{"", "", "", "", ""},
			check:    false,
			response: "orgID, email, password, firstName, lastName",
		},
		{
			inputs:   []string{"orgID", "email", "password", "firstName", "lastName"},
			check:    true,
			response: "",
		},
	}

	for _, test := range tests {
		check, response := checkCreateUserValues(
			test.inputs[0],
			test.inputs[1],
			test.inputs[2],
			test.inputs[3],
			test.inputs[4],
		)

		if check != test.check {
			t.Errorf("check received: %t, expected: %t\n", check, test.check)
		}

		if response != test.response {
			t.Errorf("response received: %s, expected: %s\n", response, test.response)
		}
	}
}

func Test_checkUpdateUserValues(t *testing.T) {
	tests := []struct {
		inputs   []string
		check    bool
		response string
	}{
		{
			inputs:   []string{"", ""},
			check:    false,
			response: "password, auth0ID",
		},
		{
			inputs:   []string{"password", "auth0ID"},
			check:    true,
			response: "",
		},
	}

	for _, test := range tests {
		check, response := checkUpdateUserValues(
			test.inputs[0],
			test.inputs[1],
		)

		if check != test.check {
			t.Errorf("check received: %t, expected: %t\n", check, test.check)
		}

		if response != test.response {
			t.Errorf("response received: %s, expected: %s\n", response, test.response)
		}
	}
}
