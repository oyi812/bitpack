package bitpack

type Packer pack

// Set also resets packer
func (p *Packer) Set(slice []byte) {
	p.slice = slice
	p.offset = 0
}

func (p *Packer) Len() int {
	return Len(p.offset)
}

func (p *Packer) Pack(n int, v uint64) {
	p.offset = Pack(p.slice, p.offset, n, v)
}

// Len returns pack encoded byte length
// caller has access to original slice
// For Backpack encoded byte length: len(slice)-offset>>3
func Len(offset int) int {
	if offset&0x7 == 0 {
		return offset >> 3
	}
	return offset>>3 + 1
}

// Pack n bits of v where 0 < n <= 64 (unchecked)
func Pack(slice []byte, offset, n int, v uint64) (o int) {

	// []byte index
	i := offset >> 3

	// remainder bits (hi)
	j := offset & 0x7

	o = offset + n

	// use any remaining space in last octet
	if j > 0 {

		// bits remaining (low)
		k := 8 - j

		// all done
		if n <= k {
			slice[i] |= byte(v) << (8 - n) >> j
			return
		}

		// highest k bits of v
		n -= k
		slice[i] |= byte(v>>n) << j >> j
		i++
	}

	for n > 8 {
		n -= 8
		slice[i] = byte(v >> n)
		i++
	}

	// final n bits where 1 <= n <= 8
	slice[i] = byte(v) << (8 - n)

	return
}

// Backpack packs n bits of v, from back to front, where 0 < n <= 64 (unchecked)
// initial offset is LHS of last bit to be packed i.e. len(slice) << 3
func Backpack(slice []byte, offset, n int, v uint64) (o int) {

	// []byte index
	i := (offset - 1) >> 3

	// bits remaining (hi)
	j := offset & 0x7

	o = offset - n

	// use any remaining space in last octet
	if j > 0 {

		// all done
		if n <= j {
			slice[i] |= byte(v) << (8 - n) >> (j - n)
			return
		}

		// lowest j bits of v
		slice[i] |= byte(v) << (8 - j)
		v >>= j
		n -= j
		i--
	}

	for n > 8 {
		slice[i] = byte(v)
		v >>= 8
		n -= 8
		i--
	}

	// final n bits where 1 <= n <= 8
	l := 8 - n
	slice[i] = byte(v) << l >> l

	return
}
