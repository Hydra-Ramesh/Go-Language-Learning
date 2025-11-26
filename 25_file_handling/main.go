package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open file once and ensure it's closed
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File or Folder:", fileInfo.IsDir())
	fmt.Println("File Size in bytes:", fileInfo.Size())
	fmt.Println("File Mode (permissions):", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())

	// Read File (reset offset to start)
	if _, err := f.Seek(0, 0); err != nil {
		panic(err)
	}

	buf := make([]byte, 100)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println("Read bytes:", n)
	fmt.Println("Content:", string(buf[:n]))

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content using ReadFile:", string(data))

	// Read Directory
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileInformation, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInformation {
		fmt.Println("File Name:", fileInfo.Name())
	}

	// Create File
	newFile, err := os.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	// // ensure file is closed and check close error
	// defer func() {
	// 	if err := newFile.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// num, err := newFile.WriteString("This is a newly created file.\n")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Wrote bytes:", num)

	bytes := []byte("This is a newly created file using byte slice.\n")
	num, err := newFile.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote bytes:", num)

	// close newFile BEFORE attempting to remove it later
	if err := newFile.Close(); err != nil {
		panic(err)
	}

	// Read and Write to Another File (Streaming)

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("copyfile.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		er := writer.WriteByte(b)
		if er != nil {
			panic(er)
		}
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}

	// Delete File
	err = os.Remove("newfile.txt")
	if err != nil {
		panic(err)
	}
}
