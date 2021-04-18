package meeting

type Meeting struct {
    Id string
    Host string
    Guest string
    Date string
    Duration int
}


type DataStore interface {
    findMeeting(host string, date string) Meeting
}

type HardCodedDataStore struct {}

func (h *HardCodedDataStore) findMeeting(host string, date string) Meeting {
    return Meeting{"12e3", "sefa", "hayta", "20210417", 30}
}

type MeetingResource struct {
    Store *DataStore
}

func NewHardCodedDataStore() *DataStore {
    var store DataStore = &HardCodedDataStore{}
    return &store
}

func NewMeetingResource(store *DataStore) *MeetingResource {
    return &MeetingResource{Store : store}
}

func (m *MeetingResource) findMeeting(host string, date string) Meeting {
    store := *m.Store
    return store.findMeeting(host, date)
}

func initializeResource() *MeetingResource {
    store := NewHardCodedDataStore()
    return NewMeetingResource(store)
}

