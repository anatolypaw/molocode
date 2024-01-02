package entity

import (
	"fmt"
	"testing"
)

func TestValidateSerial(t *testing.T) {
	tests := []struct {
		serial   string
		expected error
	}{
		{"123abc", nil},             // корректная длина
		{"abc12xx", fmt.Errorf("")}, // длинный
		{"z", fmt.Errorf("")},       // короткий
		{"", fmt.Errorf("")},        // пустой
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(`"%s"`, test.serial), func(t *testing.T) {
			result := ValidateSerial(test.serial)
			if (test.expected == nil && result != nil) || (test.expected != nil && result == nil) {
				t.Errorf("Ожидалось: %v, получено: %v", test.expected, result)
			}
		})

	}
}

func TestValidateCrypto(t *testing.T) {
	tests := []struct {
		crypto   string
		expected error
	}{
		{"1234", nil},               // корректная длина
		{"abc12as", fmt.Errorf("")}, // длинный
		{"s", fmt.Errorf("")},       // короткий
		{"", fmt.Errorf("")},        // пустой
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(`"%s"`, test.crypto), func(t *testing.T) {
			result := ValidateCrypto(test.crypto)
			if (test.expected == nil && result != nil) || (test.expected != nil && result == nil) {
				t.Errorf("Ожидалось: %v, получено: %v", test.expected, result)
			}
		})

	}
}
