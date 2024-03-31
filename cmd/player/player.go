// Package player load audio streams and play to speaker
package player

import (
	"log"
	"net/http"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/mjoe/radiotray/cmd/config"
)

var _currentStreamer *beep.StreamSeekCloser

// Play radio to speaker => blocking function
func Play(radio *config.Radio) error {
	// Stop any previous stream play
	Stop()

	resp, err := http.Get(radio.Source)
	if err != nil {
		log.Fatal(err)
	}
	reader := resp.Body

	streamer, format, err := mp3.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	_currentStreamer = &streamer

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	return nil
}

// Stop player
func Stop() {
	if _currentStreamer != nil {
		(*_currentStreamer).Close()
		_currentStreamer = nil
	}

	speaker.Clear()
}
