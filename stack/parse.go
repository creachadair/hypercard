// Package stack handles the encoded format of a HyperCard stack.
package stack

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Block struct {
	Size uint32 // block size including size, type, and ID
	Type string // type code, 4 bytes
	ID   uint32 // block ID
	Data []byte
}

// ParseOne parses one block from r reports an error. At EOF, ParseOne returns
// nil, io.EOF. If an error is encountered reading the block body after the
// header has been consumed, ParseOne returns the incomplete block with the
// error.
func ParseOne(r io.Reader) (*Block, error) {
	var header [12]byte

	if _, err := io.ReadFull(r, header[:]); err == io.EOF {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("reading header: %w", err)
	}

	blk := &Block{
		Size: binary.BigEndian.Uint32(header[:4]),
		Type: string(header[4:8]),
		ID:   binary.BigEndian.Uint32(header[8:12]),
	}
	blk.Data = make([]byte, int(blk.Size)-len(header))
	if nr, err := io.ReadFull(r, blk.Data); err != nil {
		blk.Data = blk.Data[:nr]
		return blk, fmt.Errorf("reading block data: %w", err)
	}
	return blk, nil
}
