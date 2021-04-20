package meeting

type Meeting struct {
    Id string `json:"id"`
    Host string `json:"host"`
    Guest string `json:"guest"`
    Date string `json:"date"`
    Duration int `json:"duration"`
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

