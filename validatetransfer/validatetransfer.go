package validatetransfer

import "errors"

type transferRequest struct {
	amount               float64
	currency             string
	originAccountID      string
	destinationAccountID string
}

func validateTransferRequest(request transferRequest) error {
	if request.amount <= 0 {
		return errors.New("invalid amount")
	}

	if request.currency != "USD" {
		return errors.New("invalid currency: must be USD")
	}

	if request.originAccountID == "" {
		return errors.New("empty origin account ID")
	}

	if request.destinationAccountID == "" {
		return errors.New("empty destination account ID")
	}

	return nil // request is valid so we return nil
}
