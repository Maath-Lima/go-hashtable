package tests

import (
	"hashtable/utils"
	"testing"
)

func TestUtils_FindNextPrime(t *testing.T) {
	n := 90

	rest := utils.FindNextPrime(n)
	resp := 97

	if rest != resp {
		t.Fatalf("%v does not match %v as expected next prime number.", rest, resp)
	}
}
