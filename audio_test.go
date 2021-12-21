package lightbrite

import (
	"log"
	"testing"
)

func Test_AudioStream(t *testing.T) {
	na := NewAudioDefault()
	log.Printf("Na: %v\n", na)

	na.Start()
}
