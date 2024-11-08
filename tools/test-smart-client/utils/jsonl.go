package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: DRY: copied from  https://github.com/simonfrey/jsonl/blob/main/writer.go
type JsonlWriter struct {
	w io.Writer
}

func NewJsonlWriter(w io.Writer) JsonlWriter {
	return JsonlWriter{
		w: w,
	}
}

func (w JsonlWriter) Close() error {
	if c, ok := w.w.(io.WriteCloser); ok {
		return c.Close()
	}
	return fmt.Errorf("given writer is no WriteCloser")
}

func (w JsonlWriter) Write(data interface{}) error {
	j, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("could not json marshal data: %w", err)
	}

	_, err = w.w.Write(j)
	if err != nil {
		return fmt.Errorf("could not write json data to underlying io.JsonlWriter: %w", err)
	}

	_, err = w.w.Write([]byte("\n"))
	if err != nil {
		return fmt.Errorf("could not write newline to underlying io.JsonlWriter: %w", err)
	}

	if f, ok := w.w.(http.Flusher); ok {
		// If http writer, flush as well
		f.Flush()
	}
	return nil
}
