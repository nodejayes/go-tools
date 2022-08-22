package stopwatch

import "time"

// StartSection starts the Measure for a specific section
func (sw *Stopwatch) StartSection(key string) {
	sw.sessions[key] = time.Now().UnixNano()
}
