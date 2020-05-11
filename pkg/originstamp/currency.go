package originstamp

import (
	"fmt"
	"strings"
)

type Currency int8

const (
	BITCOIN Currency = iota
	ETHEREUM
	AION
	SUEDKURIER = 100
)

func CurrencyValues() []Currency {
	return []Currency{
		BITCOIN,
		ETHEREUM,
		AION,
		SUEDKURIER,
	}
}

func CurrencyStrValues() []string {
	vals := CurrencyValues()
	strs := make([]string, len(vals))
	for i, currency := range vals {
		strs[i] = currency.String()
	}
	return strs
}

func (c Currency) String() string {
	switch c {
	case BITCOIN:
		return "BITCOIN"
	case ETHEREUM:
		return "ETHEREUM"
	case AION:
		return "AION"
	case SUEDKURIER:
		return "SUEDKURIER"
	default:
		return fmt.Sprintf("Currency(%d)", c)
	}
}

func CurrencyFromString(str string) (Currency, error) {
	switch strings.ToUpper(str) {
	case "BITCOIN":
		return BITCOIN, nil
	case "ETHEREUM":
		return ETHEREUM, nil
	case "AION":
		return AION, nil
	case "SUEDKURIER":
		return SUEDKURIER, nil
	default:
		return 0, fmt.Errorf("%s is not a valid currency", str)
	}
}
