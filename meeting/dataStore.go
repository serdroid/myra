package meeting

var MEETINGS = []Meeting {
{"12e3", "sefa", "hayta", "20210417", 30},
{"12e4", "veli", "nick", "20210419", 60},
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

type postgresDataStore struct {}

func (p *postgresDataStore) findMeeting(host string, date string) Meeting {
    return MEETINGS[1]
}

func NewPostgresDataStore() *dataStore {
    var store dataStore = &postgresDataStore{}
    return &store
}

