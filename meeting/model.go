package meeting

type Meeting struct {
    ID string `json:"id"`
    Host string `json:"host"`
    Guest string `json:"guest"`
    Date string `json:"date"`
    Duration int `json:"duration"`
}

