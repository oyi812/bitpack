package bitpack

import (
	"encoding/binary"
	"testing"
)

func expect(t *testing.T, have, must uint64) {
	if have != must {
		t.Errorf("expect %d, have %d", must, have)
	}
}

// TestPacker packs and unpacks alternating bit
// runs of 0s and 1s for each width in 1..64
func TestPacker(t *testing.T) {

	// length is double the sum of 1..64 bits in bytes
	const length = (2*32*65)>>3
	var buffer [1<<16]byte

	var ones uint64; ones--

	var p Packer
	p.Set(buffer[:])

	for w := 1; w <= 64; w++ {
		p.Pack(w, 0)
		p.Pack(w, ones)
	}

	if p.Len() != length {
		t.Errorf("expect %d, have %d", length, p.Len())
	}

	var u Unpacker
	u.Set(buffer[:p.Len()])

	for w := 1; w <= 64; w++ {

		if must, have := uint64(0), u.Unpack(w); have != must {
			t.Errorf("expect %d, have %d", must, have)
		}

		if must, have := uint64(1<<w-1), u.Unpack(w); have != must {
			t.Errorf("expect %d, have %d", must, have)
		}
	}
}

func TestBackpack(t *testing.T) {

	var buffer [64]byte

	s := buffer[:]
	o := len(s) << 3

	o = Backpack(s, o, 10, 666)
	o = Backpack(s, o, 1, 1)
	o = Backpack(s, o, 2, 2)
	o = Backpack(s, o, 7, 42)

	// 7+10+2+1=20 bits fit into 3 bytes
	l := len(s) - o>>3
	if l != 3 {
		t.Errorf("expect %d, have %d", 3, l)
	}

	o = Backpack(s, o, 4, 15)

	// 20+4=24 bits fit into 3 bytes
	l = len(s) - o>>3
	if l != 3 {
		t.Errorf("expect %d, have %d", 3, l)
	}

	_ = Backpack(s, o, 64, 1<<63 + 1)

	// Unbackpack
	var v uint64
	o = len(s) << 3

	v, o = Unbackpack(s, o, 10)
	expect(t, v, 666)

	v, o = Unbackpack(s, o, 1)
	expect(t, v, 1)

	v, o = Unbackpack(s, o, 2)
	expect(t, v, 2)

	v, o = Unbackpack(s, o, 7)
	expect(t, v, 42)

	v, o = Unbackpack(s, o, 4)
	expect(t, v, 15)

	v, _ = Unbackpack(s, o, 64)
	expect(t, v, 1<<63 + 1)
}

var sizes = [...]int{64, 64, 7, 10, 2, 1, 4, 64}

func BenchmarkPacker(b *testing.B) {

	var buffer [64]byte
	var p Packer

	// b.ResetTimer()
	for i := 0; i < b.N; i++ {

		p.Set(buffer[:])
		for _, w := range sizes {
			p.Pack(w, 1)
		}

		var u Unpacker
		var v uint64

		u.Set(buffer[:p.Len()])
		for _, w := range sizes {
			v += u.Unpack(w)
		}

		_ = v == uint64(len(sizes))
	}
}

// for comparison
func BenchmarkVarint(b *testing.B) {
	var buffer [64]byte

	for i := 0; i < b.N; i++ {

		s := buffer[:]
		n := 0

		for _, w := range sizes {
			n += binary.PutUvarint(s[n:], uint64(1)<<(w-1))
		}

		var v uint64
		o := 0

		for range sizes {
			x, n := binary.Uvarint(s[o:])
			o += n
			v += x
		}
	}
}
