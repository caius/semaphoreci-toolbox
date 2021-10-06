package files

import (
	"fmt"
)

func HumanReadableSize(b int64) string {
	const unit = 1024

	if b < unit {
		return fmt.Sprintf("%.1f", float64(b))
	}

	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f%c", float64(b)/float64(div), "KMGTPE"[exp])
}
