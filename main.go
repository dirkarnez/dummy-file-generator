package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	GiB  uint64
	MiB  uint64
	KiB  uint64
	Byte uint64
)

func main() {
	flag.Uint64Var(&GiB, "g", 0, "GiB (1024 MiBs), default is 0")
	flag.Uint64Var(&MiB, "m", 0, "MiB (1024 KiBs), default is 0")
	flag.Uint64Var(&KiB, "k", 0, "KiB (1024 Bytes), default is 0")
	flag.Uint64Var(&Byte, "b", 0, "Byte, default is 0")
	flag.Parse()

	inputSize := (GiB * 1024 * 1024 * 1024) + (MiB * 1024 * 1024) + (KiB * 1024) + Byte
	if inputSize == 0 {
		fmt.Println("--help")
		return
	}

	slice := make([]byte, inputSize)
	for i := 0; i < len(slice); i++ {
		slice[i] = 255
	}

	file, err := os.Create("test.bin")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	count, err := file.Write(slice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Written %d bytes\n", count)
}
