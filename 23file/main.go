package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	fileInfo, err := f.Stat()

	if err != nil {
		// log the error
		panic(err)
	}

	fmt.Println("file name:", fileInfo.Name())
	fmt.Println("is directory:", fileInfo.IsDir())

	// read file
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	buf := make([]byte, 10)

	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		fmt.Println("data", d, string(buf[i]))
	}

	fmt.Println("data", d, buf)

	// read folders
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	// <=0 : all files/folders
	// n number of files/folders
	fileInfo2, err := dir.ReadDir(-1)

	for _, fi := range fileInfo2 {
		fmt.Println(fi.Name())
	}

	// write files
	f3, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer f3.Close()

	f3.WriteString("hi go ")
	f3.WriteString("nice language")

	// read and write to another file (streaming fashion)

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

	fmt.Println("written to new file successfully")

	// delete a file
	err3 := os.Remove("example2.txt")
	if err3 != nil {
		panic(err3)
	}

	fmt.Println("File deleted sucessfully")
}
