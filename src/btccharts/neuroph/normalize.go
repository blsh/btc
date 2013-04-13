package neuroph

import "math/big"

type Normalizable interface {
	Normalize() []big.Rat
}
