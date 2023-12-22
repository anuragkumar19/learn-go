// This is reflection of interfaces Go's io package
package io

// Declaration of an interface
type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// As we do struct embedding we can also do interface embedding
type ReadWriter interface {
	Reader
	Writer
}

// ReadWriter is similar to
type ReadWriter2 interface {
	Write(p []byte) (n int, err error)
	Read(p []byte) (n int, err error)
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
