package neuroph

import "math/big"

type NeurophRecord interface {
	Normalize() []big.Rat
}
