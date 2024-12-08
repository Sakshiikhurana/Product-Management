package imageprocessing

import (
	"testing"
)

func TestProcessImage(t *testing.T) {
	// Call your image processing function
	err := processImage("https://example.com/image.jpg")

	// Check if the function returns an error
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
