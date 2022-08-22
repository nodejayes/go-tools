package stopwatch

import (
	"fmt"
	"strconv"
	"time"
)

// ElapsedMs get the Time in ms that has elapsed since start
func (sw *Stopwatch) ElapsedMs() float64 {
	currentTime := time.Now().UnixNano()
	return float64(currentTime-sw.start) / ToMillisecondsDivisor
}

// ElapsedMsAsString same value as ElapsedMs but converted into a string
func (sw *Stopwatch) ElapsedMsAsString() string {
	return strconv.FormatFloat(sw.ElapsedMs(), 'f', -1, 64)
}

// PrintElapsedMs logs the elapsed ms to the stdout
func (sw *Stopwatch) PrintElapsedMs(text string, reset bool) {
	println(fmt.Sprintf("[%0.2f ms] %v", sw.ElapsedMs(), text))
	if reset {
		sw.Reset()
	}
}
