package bitpack

import (
	"encoding/binary"
	"testing"
)

func assert(t *testing.T, have, must uint64) {
	if have != must {
		t.Errorf("expect %d, have %d", must, have)
	}
}

func TestPacker(t *testing.T) {

	var buffer [64]byte

	var p Packer
	p.Set(buffer[:])

	p.Pack(7, 42)
	p.Pack(10, 666)
	p.Pack(2, 2)
	p.Pack(1, 1)

	// 7+10+2+1=20 bits fit into 3 bytes
	if p.Len() != 3 {
		t.Errorf("expect %d, have %d", 3, p.Len())
	}

	p.Pack(4, 15)

	// 20+4=24 bits fit into 3 bytes
	if p.Len() != 3 {
		t.Errorf("expect %d, have %d", 3, p.Len())
	}

	p.Pack(64, 1<<63)

	var u Unpacker
	u.Set(buffer[:p.Len()])

	assert(t, u.Unpack(7), 42)
	assert(t, u.Unpack(10), 666)
	assert(t, u.Unpack(2), 2)
	assert(t, u.Unpack(1), 1)
	assert(t, u.Unpack(4), 15)
	assert(t, u.Unpack(64), 1<<63)
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

	o = Backpack(s, o, 64, 1<<63)

	// Unbackpack
	var v uint64
	o = len(s) << 3

	v, o = Unbackpack(s, o, 10)
	assert(t, v, 666)

	v, o = Unbackpack(s, o, 1)
	assert(t, v, 1)

	v, o = Unbackpack(s, o, 2)
	assert(t, v, 2)

	v, o = Unbackpack(s, o, 7)
	assert(t, v, 42)

	v, o = Unbackpack(s, o, 4)
	assert(t, v, 15)

	v, o = Unbackpack(s, o, 64)
	assert(t, v, 1<<63)
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
