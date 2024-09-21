package Handler

type Artist struct {
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	YearOfBegin int      `json:"creationDate"`
	FirstAlbum  string   `json:"firstAlbum"`
	Members     []string `json:"members"`
}

type Locations struct {
	locations []string `json:"location"`
	dates     string   `json:"dates"`
}
type Dates struct {
	dates []string `json:"dates"`
}
type Relation struct {
	DatesLocations []string `json:"datesLocations"`
}
type Data struct {
	artists   string `json:"artists"`
	locations string `json:"locations"`
	dates     string `json:"dates"`
	relation  string `json:"relation"`
}
