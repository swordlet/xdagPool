package util

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetTargetHex(t *testing.T) {
	targetHex := GetTargetHex(500)
	expectedHex := "6e128300"
	if targetHex != expectedHex {
		t.Error("Invalid targetHex")
	}

	targetHex = GetTargetHex(15000)
	expectedHex = "7b5e0400"
	if targetHex != expectedHex {
		t.Error("Invalid targetHex")
	}

	targetHex = GetTargetHex(300000)
	fmt.Println(targetHex)

	targetHex = GetTargetHex(50000)
	fmt.Println(targetHex)
}

func TestGetHashDifficulty(t *testing.T) {
	hash := "8e3c1865f22801dc3df0a688da80701e2390e7838e65c142604cc00eafe34000"
	hashBytes, _ := hex.DecodeString(hash)
	shareDiff, ok := GetHashDifficulty(hashBytes)

	if !ok && shareDiff.String() != "1009" {
		t.Error("Invalid diff")
	}
}

func TestGetHashDifficulty2(t *testing.T) {
	hash := "8e3c1865f22801dc3df0a688da80701e2390e7838e65c0000000000000000000"
	hashBytes, _ := hex.DecodeString(hash)
	shareDiff, _ := GetHashDifficulty(hashBytes)

	fmt.Println("difficulty: ", shareDiff)
}

func TestGetHashDifficultyWithBrokenHash(t *testing.T) {
	hash := ""
	hashBytes, _ := hex.DecodeString(hash)
	shareDiff, ok := GetHashDifficulty(hashBytes)

	if ok || shareDiff != nil {
		t.Error("Must be no result and not ok")
	}
}
