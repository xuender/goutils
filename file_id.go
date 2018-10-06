package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"hash"
	"hash/fnv"
)

// FileID 文件唯一标识
type FileID struct {
	hash hash.Hash
	size int64
}

// NewFileID 新建文件ID
func NewFileID(file string) (*FileID, error) {
	id := new(FileID)
	id.hash = fnv.New128()
	id.size = 0
	if file == "" {
		return id, nil
	}
	err := ReadBuf(file, func(bs []byte) { id.Write(bs) })
	return id, err
}

func (f *FileID) Write(data []byte) (int, error) {
	f.size += int64(len(data))
	return f.hash.Write(data)
}

// ID 文件ID
func (f *FileID) ID() []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(f.size))
	return bytes.Join([][]byte{
		f.hash.Sum(nil),
		removeVacant(bs),
	}, nil)
}
func (f *FileID) String() string {
	return hex.EncodeToString(f.ID())
}
func removeVacant(bytes []byte) []byte {
	l := len(bytes)
	for idx := range bytes {
		if bytes[l-idx-1] != 0 {
			return bytes[:l-idx]
		}
	}
	return bytes
}
