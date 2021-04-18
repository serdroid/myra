package meeting


type dataStore interface {
    findMeeting(host string, date string) Meeting
}

type hardCodedDataStore struct {}

func (h *hardCodedDataStore) findMeeting(host string, date string) Meeting {
    return Meeting{"12e3", "sefa", "hayta", "20210417", 30}
}

func NewHardCodedDataStore() *dataStore {
    var store dataStore = &hardCodedDataStore{}
    return &store
}

