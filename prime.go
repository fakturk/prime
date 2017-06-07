// Package prime cheks if the the given number is prime
package prime

import "math"

// IsPrime checks the given number is prime or not, returns boolean
func IsPrime(number int) bool {
	prime := false
	if number == 3 || number == 2 {
		prime = true
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {

		if number%i == 0 {
			prime = false
			break
		}
		prime = true
	}
	return prime
}
