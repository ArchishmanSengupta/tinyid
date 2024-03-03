package tinyid

import (
	"crypto/rand"
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
