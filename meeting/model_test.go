package meeting

import (
    "testing"
)

func TestFindMeeting(test *testing.T) {
    meetingResource := initializeMeetingResource()
    got := meetingResource.findMeeting("sefa", "20210417")
    want := MEETINGS[0]
    
    if got != want {
        test.Errorf("got %v, want %v", got, want)
    }
}
