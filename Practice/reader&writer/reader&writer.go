package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

// FooReader полностью соответствует контракту io.Reader
// Значит, FooReader теперь тоже является io.Reader-ом
func (FooReader *FooReader) Read(b []byte) (int, error) { // ПАТТЕРН Decorator, оборачиваем этот интерфейс, чтобы добавить свои действия, например логирование
	fmt.Println("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (FooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Println("out > ")
	return os.Stdout.Write(b)
}

func main() {
	var reader FooReader
	var writer FooWriter

	/*
		input := make([]byte, 4096)

		s, err := reader.Read(input)
		if err != nil {
			log.Fatalln("Unable to read data")
		}

		fmt.Printf("Read %d bytes from stdin\n", s)

		// допустим прочитали hi, это будет [h,i,\n] (3 байта)
		s, err = writer.Write(input[:s]) // если s = 3, выведутся 0, 1 и 2
		if err != nil {
			log.Fatalln("Unable to write data")
		}

		fmt.Printf("Wrote %d bytes to stdout\n", s)
	*/

	// ЛИБО:

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
