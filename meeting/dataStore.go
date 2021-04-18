package meeting

var MEETINGS = []Meeting {
{"12e3", "sefa", "hayta", "20210417", 30},
}

type dataStore interface {
    findMeeting(host string, date string) Meeting
}

type hardCodedDataStore struct {}

func (h *hardCodedDataStore) findMeeting(host string, date string) Meeting {
    return MEETINGS[0]
}

func NewHardCodedDataStore() *dataStore {
    var store dataStore = &hardCodedDataStore{}
    return &store
}

