package asciiartfs

import "errors"

// Check for valid of characters by runes from 32 to 126
func IsValid(s string) error {
	if len(s) > 200 {
		return errors.New("large input text")
	}
	for _, ch := range s {
		if ch < ' ' && ch != 10 || ch > '~' {
			return errors.New("invalid character in input text")
		}
	}
	return nil
}
