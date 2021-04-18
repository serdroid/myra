package meeting

import (
    "testing"
)

func TestFindMeeting(test *testing.T) {
    meetingResource := initializeMeetingResource()
    got := meetingResource.findMeeting("sefa", "20210417")
    want := Meeting{"12e3", "sefa", "hayta", "20210417", 30}
    
    if got != want {
        test.Errorf("got %v, want %v", got, want)
    }
}
