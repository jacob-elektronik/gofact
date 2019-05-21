package reader

import (
	"bufio"
	"io"
	"log"
	"os"
)

type EdiReader struct {
	file      *os.File
	buf       []byte
	BufReader *bufio.Reader
}

// NewEdiReader buffered reader
func NewEdiReader(fileStr string) *EdiReader {
	r := EdiReader{}
	r.file, _ = os.Open(fileStr)
	r.BufReader = bufio.NewReader(r.file)
	r.buf = make([]byte, 256)
	return &r
}

func (r *EdiReader) ReadFile(ch chan<- []byte) {
	for {
		n, err := r.BufReader.Read(r.buf)

		if n > 0 {
			ch <- append([]byte{}, r.buf[:n]...)

		}

		if err == io.EOF {
			close(ch)
			err = r.file.Close()
			break
		}

		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			err = r.file.Close()
			break
		}
	}
}
