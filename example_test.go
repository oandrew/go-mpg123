package mpg123_test

import (
	"encoding/binary"
	"io"
	"log"
	"os"

	"github.com/oandrew/go-mpg123"
)

func Example() {

	file, err := os.Open("audio.mp3")
	if err != nil {
		return
	}

	cfg := mpg123.ReaderConfig{
		OutputFormat: &mpg123.OutputFormat{
			Channels: 2,
			Rate:     44100,
			Encoding: mpg123.EncodingInt16,
		},
	}

	r := mpg123.NewReaderConfig(file, cfg)

	//Parse metadata
	r.Read(nil)
	log.Printf("Format: %+v", r.OutputFormat())
	log.Printf("FrameInfo: %+v", r.FrameInfo())
	log.Printf("Meta ID3v2:  %#v", r.Meta().ID3v2)

	// 320 samples.  signed 16bit
	samples := make([]int16, 320)
	for {
		log.Printf("Offset: %+v", r.Offset())
		err := binary.Read(r, binary.LittleEndian, samples)
		// Process samples - save to disk, stream to speakers
		if err == io.EOF {
			break
		}

	}
}
