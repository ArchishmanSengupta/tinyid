package tinyid

import (
	"crypto/rand"
	"errors"
	"math"
	// "math/bits"
	// "strings"
)

// Represent the default alphabet and size for the generated tiny IDs.
const (
	// Default is the default alphabet for TinyIDs.
	DefaultAlphabet = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	// DefaultSize is the default size for Tiny IDs.
	DefaultSize = 31
)

// Generate Bytes represents random bytes buffer
// 'step' parameter representing the number of bytes to generate and returns a byte slice and an error
type GenerateBytes func(step int) ([]byte, error)

// Generates a random buffer of bytes using cryptographic randomness
func randomBufferGenerator(step int) ([]byte, error) {
	buffer := make([]byte, step)
	_, err := rand.Read(buffer)
	return buffer, err
}

// GenerateTinyId generates a random sttring based on the alphabet and size.
func generateTinyId(alphabet string, size int) (string, error) {
	if size <= 0 {
		return "", errors.New("tinyId: size must be greater than 0")
	}

	// Calculate mask based on alphabet length
	mask := len(alphabet) - 1

	// Calculate step size based on mask and size
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(len(alphabet))))

	// Initialize byte slice to hold ID
	id := make([]byte, size)

	// Generate ID
	for i := 0; i < size; {
		// Generate random buffer of bytes
		randomBuffer, err := randomBufferGenerator(step)
		if err != nil {
			return "", err
		}

		// Process random buffer
		for _, b := range randomBuffer {
			// Map byte to character in alphabet using mask
			index := int(b) & mask

			// Check if index is within alphabet range
			if index < len(alphabet) {
				// Add character to ID
				id[i] = alphabet[index]
				i++

				// Check if ID is complete
				if i == size {
					break
				}
			}
		}
	}

	return string(id), nil
}

// NewTinyID generates a random string with default settings.
func NewTinyID() (string, error) {
	return generateTinyId(DefaultAlphabet, DefaultSize)
}
