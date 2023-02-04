package lib

import (
	"strings"
	"testing"
)

func TestCurrentTime(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "simple test",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(strings.Split(CurrentTime(), ":")); got != tt.want {
				t.Errorf("CurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
