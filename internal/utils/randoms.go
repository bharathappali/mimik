package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func getRandomInRange_UINT8(
	// Params:
	// 1 - Minimum number to generate
	min uint8,
	// 2 - Maximum number to generate
	max uint8) (uint8, error) { // Returns the random generated and an error

	if min == max {
		return min, nil
	}

	if min > max {
		return 0, fmt.Errorf("Min: %d cannot be greater than Max: %d\n", min, max)
	}

	// Create range
	rng := big.NewInt(int64(max - min + 1))

	// Generate a random number within the specified range
	n, err := rand.Int(rand.Reader, rng)
	if err != nil {
		return 0, err
	}

	// Add min to the generated random number to get a number in the desired range
	return uint8(n.Uint64()) + min, nil
}

func getRandomInRange_UINT64(
	// Params:
	// 1 - Minimum number to generate
	min uint64,
	// 2 - Maximum number to generate
	max uint64) (uint64, error) { // Returns the random generated and an error

	if min == max {
		return min, nil
	}

	if min > max {
		return min, fmt.Errorf("Min: %d cannot be greater than Max: %d\n", min, max)
	}

	// Create range
	rng := big.NewInt(int64(max - min + 1))

	// Generate a random number within the specified range
	n, err := rand.Int(rand.Reader, rng)
	if err != nil {
		return min, err
	}

	// Add min to the generated random number to get a number in the desired range
	return n.Uint64() + min, nil
}

func getRandomInRange_FLOAT64( // Params:
	// 1 - Minimum number to generate
	min float64,
	// 2 - Maximum number to generate
	max float64) (float64, error) { // Returns the random generated and an error
	if min == max {
		return min, nil
	}

	if min > max {
		return min, fmt.Errorf("Min: %f cannot be greater than Max: %f\n", min, max)
	}

	// Calculate the range
	rangeSize := max - min

	// Generate a random number within the specified range
	n, err := rand.Int(rand.Reader, big.NewInt(int64(rangeSize*1e18)))
	if err != nil {
		return min, err
	}

	// Convert the random number to a float64 and scale and shift it to the desired range
	r := min + (float64(n.Int64()) / 1e18)

	return r, nil
}
