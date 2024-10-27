package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	/* How we want to write */
	enc = binary.BigEndian
)

const (
	/* Eight bytes per record */
	lenWidth = 8
)

/*
Store holds all of our records
*/
type store struct {
	*os.File
	mu sync.Mutex
	/* Storing data before writing to file */
	buf *bufio.Writer
	/* Size of log file */
	size uint64
}

/*
Creates a store for a given file
*/
func newStore(f *os.File) (*store, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	size := uint64(fi.Size())
	return &store{
		File: f,
		size: size,
		buf:  bufio.NewWriter(f),
	}, nil
}

/*
Adding a new record to our store
*/
func (s *store) Append(p []byte) (n uint64, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Where new data will be written
	pos = s.size
	// Writing length to buffer
	if err := binary.Write(s.buf, enc, uint64(len(p))); err != nil {
		return 0, 0, err
	}
	w, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}
	w += lenWidth
	s.size += uint64(w)
	return uint64(w), pos, nil
}

/*
Reading a record given an offset
*/
func (s *store) Read(pos uint64) (p []byte, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Flushing everything from buffer first
	if err := s.buf.Flush(); err != nil {
		return nil, err
	}
	// Preparing byte array for reading record size
	size := make([]byte, lenWidth)
	if _, err := s.File.ReadAt(size, int64(pos)); err != nil {
		return nil, err
	}
	// Preparing byte array for reading record data
	b := make([]byte, enc.Uint64(size))
	// Skipping over first 8 bytes (record size)
	if _, err := s.File.ReadAt(b, int64(pos+lenWidth)); err != nil {
		return nil, err
	}
	return b, nil
}

/*
Reading data in p byte slice
*/
func (s *store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.buf.Flush(); err != nil {
		return 0, err
	}
	return s.File.ReadAt(p, off)
}

/*
Closing the file
*/
