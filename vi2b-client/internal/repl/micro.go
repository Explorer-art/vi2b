package repl

import (
	"log"

	"github.com/pion/mediadevices"
	_ "github.com/pion/mediadevices/pkg/driver/microphone"
)

func MicroCommand(args []string) {

	stream, err := mediadevices.GetUserMedia(mediadevices.MediaStreamConstraints{
		Audio: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	audioTrack := stream.GetAudioTracks()[0]
	defer audioTrack.Close()

}
