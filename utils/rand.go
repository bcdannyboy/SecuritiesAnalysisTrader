package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math/rand"
	"time"
)

func GetRandomSeed() int64 {
	// Generate a UUID
	uuidBytes := make([]byte, 16)
	_, err := rand.Read(uuidBytes)
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}

	// Get current time in nanoseconds
	currentTime := time.Now().UnixNano()

	// Combine current time and UUID
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, currentTime)
	if err != nil {
		log.Fatalf("Failed to write time to buffer: %v", err)
	}
	combinedBytes := append(buf.Bytes(), uuidBytes...)

	// Hash the combined bytes
	hash := sha256.Sum256(combinedBytes)

	// Convert the first 8 bytes of the hash to an int64
	var seed int64
	err = binary.Read(bytes.NewReader(hash[:8]), binary.BigEndian, &seed)
	if err != nil {
		log.Fatalf("Failed to read seed from hash: %v", err)
	}

	return seed
}
