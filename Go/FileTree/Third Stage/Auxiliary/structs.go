package Auxiliary

import (
	"crypto/md5"
	"io"
	"log"
	"os"
)

// FileStruct represent a file
type FileStruct struct {
	Path             string
	DuplicatedNumber int
	Hash             [md5.Size]byte
	Size             int64
}

// FilesDataStructures represent files data structures
type FilesDataStructures struct {
	FileStructList        []FileStruct
	SizeFilesMap          map[int64][]FileStruct
	OrderedSizeSlice      []int64
	SizeHashesFilesMapMap map[int64]map[[md5.Size]byte][]FileStruct
}

// CreateFileStruct create new FileStruct from path.
// Calculates Hash, but DuplicatedNumber remains 0
func CreateFileStruct(path string, size int64) FileStruct {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	var HashArray [md5.Size]byte
	md5Creator := md5.New()
	_, err = io.Copy(md5Creator, file)
	if err != nil {
		return FileStruct{}
	}
	HashSlice := md5Creator.Sum(nil)
	for i, e := range HashSlice {
		HashArray[i] = e
	}
	return FileStruct{
		Path:             path,
		DuplicatedNumber: 0,
		Hash:             HashArray,
		Size:             size,
	}
}

// return a new container for the files data structures
func NewFilesDataStructures() FilesDataStructures {
	filesDataStructures := FilesDataStructures{
		SizeFilesMap:          make(map[int64][]FileStruct),
		SizeHashesFilesMapMap: make(map[int64]map[[16]byte][]FileStruct),
	}
	return filesDataStructures
}
