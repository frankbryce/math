// Package structure contains data structures relevant for other packages
// in math to use, and also can be used by other developers that would like to
// extend github.com/frankbryce/math
package structure

type Formula interface {
	Evaluation() Element
	PrintTree() string
	String() string
}

type formula struct {
}
