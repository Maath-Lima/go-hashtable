package utils

import (
	"math"
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

	for rounds > 0 {
		a := float64(rand.Intn(max-min) + min)

		x := int(math.Pow(a, float64(d))) % m

		if !(x == 1 || x == (m-1)) {
			return false
		}

		s = s - 1
		hitedNMinusOneTime := false

		for s > 0 {
			x = int(math.Pow(a, 2)) % m

			if x == (m - 1) {
				hitedNMinusOneTime = true
				continue
			}

			if x == 1 && !hitedNMinusOneTime {
				return false
			}

			s--
		}

		rounds--
	}

	return true
}
