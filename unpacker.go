package bitpack

type Unpacker pack

// Set also resets unpacker
func (u *Unpacker) Set(slice []byte) {
	u.slice = slice
	u.offset = 0
}

func (u *Unpacker) Unpack(n int) (v uint64) {
	v, n = Unpack(u.slice, u.offset, n)
	u.offset = n
	return
}

// Unpack n bits to v where 0 < n <= 64 (unchecked)
func Unpack(slice []byte, offset, n int) (v uint64, o int) {

	// []byte index
	i := offset >> 3

	// remainder bits (hi)
	j := offset & 0x7

	o = offset + n

	// pop any bits in first octet
	if j > 0 {

		// bits remaining (low)
		k := 8 - j

		// all done
		if n <= k {
			v = uint64(slice[i] << j >> (8-n))
			return
		}

		// highest k bits of v
		n -= k
		v = uint64(slice[i] << j >> j)
		i++
	}

	for n > 8 {
		n -= 8
		v <<= 8
		v |= uint64(slice[i])
		i++
	}

	// final n bits where 1 <= n <= 8
	v <<= n
	v |= uint64(slice[i] >> (8-n))

	return
}

// Unbackpack unpack n bits, from back to front, to v where 0 < n <= 64 (unchecked)
func Unbackpack(slice []byte, offset, n int) (v uint64, o int) {

	// []byte index
	i := (offset-1) >> 3

	// bits remaining (hi)
	j := offset & 0x7

	// bits of n shifted into v
	m := 0

	o = offset - n

	// pop any bits in first octet
	if j > 0 {

		// all done
		if n <= j {
			v = uint64(slice[i] << (j-n) >> (8-n))
			return
		}

		// lowest j bits of v
		v = uint64(slice[i] >> (8-j))
		m += j
		i--
	}

	for n-m > 8 {
		v |= uint64(slice[i]) << m
		m += 8
		i--
	}

	// final n-m bits where 1 <= (n-m) <= 8
	l := 8-(n-m)
	v |= uint64(slice[i] << l >> l) << m

	return
}
