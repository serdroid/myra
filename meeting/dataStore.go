package meeting

var MEETINGS = []Meeting {
{"12e3", "sefa", "hayta", "20210417", 30},
{"12e4", "veli", "nick", "20210419", 60},
}

// defines an interface for data stores to decouple implementation
type dataStore interface {
    // finds meeting for given host and date
    findMeeting(host string, date string) Meeting
    // creates new meeting record and returns ID of the new record
    createMeeting(m *Meeting) error
}

type hardCodedDataStore struct {}

func (h *hardCodedDataStore) findMeeting(host string, date string) Meeting {
    return MEETINGS[0]
}

// creates new hardCodedDataStore instance as dataStore
func NewHardCodedDataStore() *dataStore {
    var store dataStore = &hardCodedDataStore{}
    return &store
}

func (h *hardCodedDataStore) createMeeting(m *Meeting) error {
    return nil
}
