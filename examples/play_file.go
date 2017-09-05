package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"net/http"
	"os"
	_ "os"
	"strings"

	"github.com/gordonklaus/portaudio"
	mpg "github.com/oandrew/go-mpg123"
)

func open(path string) (io.ReadCloser, error) {
	if strings.HasPrefix(path, "http") {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Provide a file path or an URL")
	}

	s, err := open(os.Args[1])
	if err != nil {
		log.Fatalf("Could not open the stream: %v", err)
	}
	defer s.Close()

	portaudio.Initialize()
	defer portaudio.Terminate()

	r := mpg.NewReaderConfig(bufio.NewReader(s), mpg.ReaderConfig{
		OutputFormat: &mpg.OutputFormat{
			Channels: 2,
			Rate:     44100,
			Encoding: mpg.EncodingInt16,
		},
	})

	r.Read(nil)
	log.Printf("Format: %+v", r.OutputFormat())
	log.Printf("FrameInfo: %+v", r.FrameInfo())
	log.Printf("Offset: %+v", r.Offset())
	log.Printf("Meta ID3v2:  %#v", r.Meta().ID3v2)

	samples := make([]int16, 320)
	stream, err := portaudio.OpenDefaultStream(0, 2, 44100, len(samples), &samples)

	stream.Start()
	defer stream.Stop()

	for {
		err := binary.Read(r, binary.LittleEndian, samples)
		stream.Write()
		if err == io.EOF {
			break
		}
	}

}
