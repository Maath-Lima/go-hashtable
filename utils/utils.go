package utils

import (
	"math/big"
	"math/rand"
)

func FindNextPrime(n int) int {
	// Bertrand's Postulate -> (n > 3) n < p < 2*n - 2
	i := n

	for i < (2*n - 2) {
		// Miller-Rabin
		if millerRabin(i) {
			break
		}

		i++
	}

	return i
}

func millerRabin(m int) bool {
	rounds := 20
	d := m - 1
	var s int

	for d%2 == 0 {
		d = d >> 1
		s++
	}

	min := 2
	max := (m - 2) + 1

	for i := 0; i < rounds; i++ {
		a := float64(rand.Intn(max-min) + min)

		x := int(
			new(big.Int).Exp(
				big.NewInt(int64(a)),
				big.NewInt(int64(d)),
				big.NewInt(int64(m))).Uint64())

		if x == 1 || x == (m-1) {
			continue
		}

		composite := true

		for j := 0; j < s; j++ {
			x = (x * x) % m

			if x == (m - 1) {
				composite = false
				break
			}

			if x == 1 {
				return false
			}
		}

		if composite == true {
			return false
		}
	}

	return true
}
