package meeting

type Meeting struct {
    Id string
    Host string
    Guest string
    Date string
    Duration int
}


type MeetingResource struct {
    Store *dataStore
}

func NewMeetingResource(store *dataStore) *MeetingResource {
    return &MeetingResource{Store : store}
}

func (m *MeetingResource) findMeeting(host string, date string) Meeting {
    store := *m.Store
    return store.findMeeting(host, date)
}

func initializeMeetingResource() *MeetingResource {
    store := NewHardCodedDataStore()
    return NewMeetingResource(store)
}

