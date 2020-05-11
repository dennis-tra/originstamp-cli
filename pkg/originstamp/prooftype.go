package originstamp

import (
	"fmt"
	"strings"
)

type ProofType int8

const (
	// Proof with a seed file (txt) or proof with a merkle tree (xml)
	SEED ProofType = iota
	PDF
)

func ProofTypeValues() []ProofType {
	return []ProofType{
		SEED,
		PDF,
	}
}

func ProofTypeStrValues() []string {
	vals := ProofTypeValues()
	strs := make([]string, len(vals))
	for i, proofType := range vals {
		strs[i] = proofType.String()
	}
	return strs
}

func (c ProofType) String() string {
	switch c {
	case SEED:
		return "SEED"
	case PDF:
		return "PDF"
	default:
		return fmt.Sprintf("ProofType(%d)", c)
	}
}

func ProofTypeFromString(str string) (ProofType, error) {
	switch strings.ToUpper(str) {
	case "SEED":
		return SEED, nil
	case "PDF":
		return PDF, nil
	default:
		return 0, fmt.Errorf("%s is not a valid proof type", str)
	}
}
