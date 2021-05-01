package meeting

import (
    "testing"
    "os"
    "bytes"
    "net/http"
    "net/http/httptest"
    "encoding/json"
)

// Test implementation of the DependencyWeaver interface
type TestWire struct {}

func (t TestWire) weave() *dependencies {
    store := NewHardCodedDataStore()
    return &dependencies{store}
}

func TestMain(m *testing.M) {
    var wr DependencyWeaver = &TestWire{}
    App.initDeps(&wr)
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
    App.router.ServeHTTP(responseRecorder, req)
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

func TestMarshalMeeting(test *testing.T) {
    nm := Meeting{Host:"efe", Guest:"kahya", Date:"20210429", Duration:30}
    by, _ := json.Marshal(&nm)
    got := string(by)
    want := `{"id":"","host":"efe","guest":"kahya","date":"20210429","duration":30}`

    if got != want {
        test.Errorf("marshall err: got %v want %v", got, want)
    }
}


func TestCreateMeeting(test *testing.T) {
    var jsonStr = []byte(`{"host":"efe","guest":"kahya","date":"20210429","duration":30}`)
    req, _ := http.NewRequest("POST", "/meeting", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
	responseRecorder := executeRequest(req)
	checkResponseStatus(test, responseRecorder.Code)
    
    //var m map[string]interface{}
    //json.Unmarshal(response.Body.Bytes(), &m)	
    got := responseRecorder.Body.String()
    want := "{\"id\":\"ef32\"}\n"
    if got != want {
        test.Errorf("handler returned unexpected body: got %v want %v", got, want)
    }
}

func BenchmarkRandomStr(bench *testing.B) {
    for i:= 0; i < bench.N; i++ {
        randomString(16)
    }
}

