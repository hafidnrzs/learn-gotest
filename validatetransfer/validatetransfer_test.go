package validatetransfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorShouldError(t *testing.T) {
	type errorTestCases struct {
		description   string
		input         transferRequest
		expectedError string
	}

	for _, scenario := range []errorTestCases{
		{
			description: "invalid amount",
			input: transferRequest{
				amount:               0,
				currency:             "USD",
				originAccountID:      "checking",
				destinationAccountID: "savings",
			},
			expectedError: "invalid amount",
		},
		{
			description: "invalid currency",
			input: transferRequest{
				amount:               150.99,
				currency:             "INR",
				originAccountID:      "checking",
				destinationAccountID: "savings",
			},
			expectedError: "invalid currency: must be USD",
		},
		{
			description: "invalid origin account ID",
			input: transferRequest{
				amount:               150.99,
				currency:             "USD",
				originAccountID:      "",
				destinationAccountID: "savings",
			},
			expectedError: "empty origin account ID",
		},
		{
			description: "invalid destination account ID",
			input: transferRequest{
				amount:               150.99,
				currency:             "USD",
				originAccountID:      "checking",
				destinationAccountID: "",
			},
			expectedError: "empty destination account ID",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			err := validateTransferRequest(scenario.input)
			require.Error(t, err)
			assert.Equal(t, scenario.expectedError, err.Error())
		})
	}
}
