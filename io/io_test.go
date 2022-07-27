package io

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCopy(t *testing.T) {
	str := "some io.Reader stream to be read\n"
	r := strings.NewReader(str)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		t.Fatal(err)
	}
	if _, err := fmt.Fprint(os.Stdout, str); err != nil {
		t.Fatal(err)
	}
}

func TestCopyBuffer(t *testing.T) {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 4)

	// buf is used here...
	c, err := io.CopyBuffer(os.Stdout, r1, buf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %d", c)

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		t.Fatal(err)
	}
}

func TestCopyN(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read")

	c, err := io.CopyN(os.Stdout, r, 4)
	if err != nil {
		t.Fatal(err)
	}
	c, err = io.CopyN(os.Stdout, strings.NewReader("\n"), 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("re: %v", c)
}

/*
Pipe creates a synchronous in-memory pipe.
It can be used to connect code expecting an io.Reader with code expecting an io.Writer.
Reads and Writes on the pipe are matched one to one except when multiple Reads are needed to consume a single Write.

That is, each Write to the PipeWriter blocks until it has satisfied one or more Reads from the PipeReader that fully consume the written data.
The data is copied directly from the Write to the corresponding Read (or Reads); there is no internal buffering.

It is safe to call Read and Write in parallel with each other or with Close.
Parallel calls to Read and parallel calls to Write are also safe: the individual calls will be gated sequentially.
*/
func TestPip(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		t.Fatal(err)
	}
}

func TestReadAll(t *testing.T) {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.\n")
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s", b)
}

func TestReadAtLeast(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 14)
	c, err := io.ReadAtLeast(r, buf, 4)
	if err != nil {
		t.Error(err)
	}
	t.Logf("c:%v", c)
	t.Logf("%s\n", buf)

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		t.Error("error:", err)
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		t.Error("error:", err)
	}
}

func TestReadFull(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		t.Error(err)
	}
	t.Logf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		t.Error("error:", err)
	}
}

func TestWriteString(t *testing.T) {
	if _, err := io.WriteString(os.Stdout, "Hello World\n"); err != nil {
		t.Fatal(err)
	}
}
