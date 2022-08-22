package stopwatch

import "time"

// StopSection stops the measure for a specific section
func (sw *Stopwatch) StopSection(key string) float64 {
	currentTime := time.Now().UnixNano()
	tmp := float64(currentTime-sw.sessions[key]) / ToMillisecondsDivisor
	delete(sw.sessions, key)
	return tmp
}
