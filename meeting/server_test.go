package meeting

import (
    "testing"
    "os"
    "bytes"
    "net/http"
    "net/http/httptest"
    "encoding/json"
)

func TestMain(m *testing.M) {
    var aWire wireDependencies = &AppWire{}
    Application.Initialize(&aWire)
    code := m.Run()
    os.Exit(code)
}

func TestGreet(test *testing.T) {
    buffer := bytes.Buffer{}
    greet(&buffer, "world")
    
    got := buffer.String()
    want := "Hello, world\n"
    
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

func TestGetHello(test *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		test.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(greeterHandler)
	handler.ServeHTTP(responseRecorder, req)
	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

    got := responseRecorder.Body.String()
    want := "Hello, world\n"
    if got != want {
        test.Errorf("handler returned unexpected body: got %q want %q", got, want)
    }
}

func TestGetBye(test *testing.T) {
	req, err := http.NewRequest("GET", "/bye", nil)
	if err != nil {
		test.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(byeHandler)
	handler.ServeHTTP(responseRecorder, req)
	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

    got := responseRecorder.Body.String()
    want := "Good bye"
    if got != want {
        test.Errorf("handler returned unexpected body: got %q want %q", got, want)
    }
}

func TestGetMeeting(test *testing.T) {
	req, err := http.NewRequest("GET", "/meeting/ali", nil)
	if err != nil {
		test.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
    Application.router.ServeHTTP(responseRecorder, req)
	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
    var got Meeting
    json.Unmarshal([]byte(responseRecorder.Body.String()), &got)
    want := MEETINGS[0]
    if got != want {
        test.Errorf("handler returned unexpected body: got %q want %q", got, want)
    }
}
