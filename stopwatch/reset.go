package stopwatch

import "time"

// Reset the Time to now
func (sw *Stopwatch) Reset() {
	sw.start = time.Now().UnixNano()
}
