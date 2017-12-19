package util

import (
	"fmt"
	"time"
)

func PrettyPrintBytes(n int) string {
	b := float64(n)
	if n > 1000*1000*1000 {
		return fmt.Sprintf("%.4fGB", b/(1000*1000*1000))
	}
	if n > 1000*1000 {
		return fmt.Sprintf("%.4fMB", b/(1000*1000))
	}
	if n > 1000 {
		return fmt.Sprintf("%.4fKB", b/1000)
	}
	return fmt.Sprintf("%dB", n)
}

func PerSecond(n int, timeSpan time.Duration) int {
	return n / int(timeSpan.Seconds())
}
