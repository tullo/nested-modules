package lib

import "time"

// CurrentTime returns current time in format HH:mm:ss
func CurrentTime() string {
	return time.Now().Format("15:04:05")
}
