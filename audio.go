package lightbrite

import (
	"github.com/gordonklaus/portaudio"
)

func NewAudioDefault() *portaudio.Stream {
	sampleRate := 44100
	portaudio.Initialize()

	buffer := make([]float32, sampleRate*2)

	stream, err := portaudio.OpenDefaultStream(1, 0, float64(sampleRate), len(buffer), func(in []float32) {
		for i := range buffer {
			buffer[i] = in[i]
		}
	})

	if err != nil {
		panic(err)
	}
	return stream
}
