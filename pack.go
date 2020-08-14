// Package bitpack is for packing and unpacking a sequence of variable bit
// width values into a compact byte array. Width and order information is not
// encoded and must be known to the caller or included as metadata values.
package bitpack

type pack struct {
	slice []byte
	offset int // in bits
}
