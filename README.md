
# go-mpg123
[![GoDoc](https://godoc.org/github.com/oandrew/go-mpg123?status.svg)](https://godoc.org/github.com/oandrew/go-mpg123)

mp123 bindings for Go

For full example using PortAudio check [examples/play_file.go](examples/play_file.go)

```go
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
  // Process samples - save to disk, stream to speakers, etc.
  if err == io.EOF {
    break
  }

}
```
