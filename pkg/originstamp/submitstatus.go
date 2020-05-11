package originstamp

import (
	"fmt"
	"strings"
)

type SubmitStatus int8

const (
	// The hash was not broadcasted yet, but received from our backend
	RECEIVED SubmitStatus = iota

	// The hash was included into a transaction and broadcasted to the network, but not included into a block
	BROADCASTED

	// The transaction was included into the latest block
	INCLUDED

	// the timestamp for your hash was successfully created.
	TAMPER_PROOF
)

func SubmitStatusValues() []SubmitStatus {
	return []SubmitStatus{
		RECEIVED,
		BROADCASTED,
		INCLUDED,
		TAMPER_PROOF,
	}
}

func SubmitStatusStrValues() []string {
	vals := SubmitStatusValues()
	strs := make([]string, len(vals))
	for i, currency := range vals {
		strs[i] = currency.String()
	}
	return strs
}

func (c SubmitStatus) String() string {
	switch c {
	case RECEIVED:
		return "RECEIVED"
	case BROADCASTED:
		return "BROADCASTED"
	case INCLUDED:
		return "INCLUDED"
	case TAMPER_PROOF:
		return "TAMPER_PROOF"
	default:
		return fmt.Sprintf("SubmitStatus(%d)", c)
	}
}

func SubmitStatusFromString(str string) (SubmitStatus, error) {
	switch strings.ToUpper(str) {
	case "RECEIVED":
		return RECEIVED, nil
	case "BROADCASTED":
		return BROADCASTED, nil
	case "INCLUDED":
		return INCLUDED, nil
	case "TAMPER_PROOF":
		return TAMPER_PROOF, nil
	default:
		return 0, fmt.Errorf("%s is not a valid submit status", str)
	}
}
