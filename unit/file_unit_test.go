package unit

import (
	"fmt"
	"testing"
)

func TestFormatFileSize(t *testing.T) {
	// 118.07 PiB
	fmt.Print(FormatFileSize(132933635857711104))
}