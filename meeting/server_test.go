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

func createRequest(test *testing.T, method, path string) *http.Request {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		test.Fatal(err)
	}
	return req
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
    Application.router.ServeHTTP(responseRecorder, req)
    return responseRecorder
}

func checkResponseStatus(test *testing.T, status int) {
	if status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetHello(test *testing.T) {
	req := createRequest(test, "GET", "/hello")
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(greeterHandler)
	handler.ServeHTTP(responseRecorder, req)
	checkResponseStatus(test, responseRecorder.Code)

    got := responseRecorder.Body.String()
    want := "Hello, world\n"
    if got != want {
        test.Errorf("handler returned unexpected body: got %q want %q", got, want)
    }
}

func TestGetBye(test *testing.T) {
	req := createRequest(test, "GET", "/bye")
	responseRecorder := executeRequest(req)
	checkResponseStatus(test, responseRecorder.Code)

    got := responseRecorder.Body.String()
    want := "Good bye"
    if got != want {
        test.Errorf("handler returned unexpected body: got %q want %q", got, want)
    }
}

func TestGetMeeting(test *testing.T) {
	req := createRequest(test, "GET", "/meeting/ali")
	responseRecorder := executeRequest(req)
	checkResponseStatus(test, responseRecorder.Code)
	
    var got Meeting
    json.Unmarshal([]byte(responseRecorder.Body.String()), &got)
    want := MEETINGS[0]
    if got != want {
        test.Errorf("handler returned unexpected body: got %v want %v", got, want)
    }
}
