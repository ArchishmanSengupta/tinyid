package tinyid

import (
	"crypto/rand"
	"errors"
	"math"
	"math/bits"
	"strings"
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

	// calculate mask to ensure that random bytes are mapped properly to the given alphabet
	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(len(alphabet)-1|1))) - 1

	// calulates the step size based on the desired size and alphabet length to ensure an even distribution of characters
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(len(alphabet))))

	// initialize string builder to efficiently generate the ID
	id := new(strings.Builder)

	id.Grow(size)

	// iterating to generate the ID
	for {
		// generate a random buffer of bytes
		randomBuffer, err := randomBufferGenerator(step)
		if err != nil {
			return "", err
		}

		for i := 0; i < step; i++ {
			// mapping each byte to a character in the alphabet using bitwise operations and the calculated mask
			currentIndex := int(randomBuffer[i]) & mask
			if currentIndex < len(alphabet) {
				if err := id.WriteByte(alphabet[currentIndex]); err != nil {
					return "", err
				} else if id.Len() == size {
					return id.String(), nil
				}
			}
		}
	}
}

// NewTinyID generates a random string with default settings.
func NewTinyID() (string, error) {
	return generateTinyId(DefaultAlphabet, DefaultSize)
}
