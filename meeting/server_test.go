package meeting

import (
    "testing"
    "bytes"
)

func TestGreet(test *testing.T) {
    buffer := bytes.Buffer{}
    greet(&buffer)
    
    got := buffer.String()
    want := "Hello world"
    
    if got != want {
        test.Errorf("got %q want %q", got, want)
    }
}

func TestBye(test *testing.T) {
    buffer := bytes.Buffer{}
    bye(&buffer)
    
    got := buffer.String()
    want := "Good bye"
    
    if got != want {
        test.Errorf("got %q want %q", got, want)
    }
}
