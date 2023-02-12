package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello there Marcin! How are you doing?"))
	wc.Close()

	// You could also do type conversion
	// object . type to convert to
	// WARNING!!!
	// The type needs to be a pointer, because at least one of the methods of the type
	// requires a pointer receiver, hence the whole interface needs to be defined with a pointer
	bwc := wc.(*BufferedWriterCloser)
	// only prints address, so that bwc is used in the app
	fmt.Println(bwc)

	// This will fail and panic
	// Because io Reader requires a Read method, which we are not implementing
	//r := wc.(io.Reader)
	//fmt.Println(r)

	// To prevent the panic, we can use the "ok logic" like so
	// but it still fails
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed 1")
	}

	// This goes through, obviously
	r2, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(r2)
	} else {
		fmt.Println("Conversion failed 2")
	}

	// Empty interface, useful when type compatibility is an issue
	var myObj interface{} = NewBufferedWriterCloser()
	// To make the empty interace is useful either use reflection
	// or type conversion like here
	if wc2, ok := myObj.(WriterCloser); ok {
		wc2.Write([]byte("Hello again Marcin! Those examples are getting tricky"))
		wc2.Close()
	}
	r3, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r3)
	} else {
		fmt.Println("Conversion failed 3")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Interface composed of other interfaces
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// The function stores data in an internal buffer (defined by the struct)
// It also writes out in chunks of 8 bytes, but nothing if less than 8 bytes is left
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println("write", string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

// The closer empties the rest of the buffer
// and prints the data
// so it's basically taking care of anything left over by the Write method which may skip a chunk
// if it is shorter than 8 bytes
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println("close", string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// This is a constructor method, to initialize the buffer
// Mind, that the constructor defines it as a pointer
// That is because some methods of the type, require pointer receivers
// If declaring interface with a pointer, you've got access
// to methods with both pointer and value receivers
// If declaring interface without a pointer, only to methods with value receivers
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
