package main

func FindNextPrime(n int) int {

	// Bertrand's Postulate -> (n > 3) n < p < 2n*p - 2
	for i := n + 1; i < (2*n - 2); i++ {
		// Miller-Rabin
		if millerRabin(i) {

		}
	}
}

func millerRabin(m int) bool {
	d := m - 1
	var s int

	for d%2 == 0 {
		d = d >> 2
		s++
	}
}
