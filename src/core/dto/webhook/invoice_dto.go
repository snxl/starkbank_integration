package webhook

type Request struct {
	Event Event `json:"event"`
}

type Event struct {
	Subscription string `json:"subscription"`
	Log          Log    `json:"log"`
}

type Invoice struct {
	Id     string `json:",omitempty"`
	Amount int    `json:",omitempty"`
	Status string `json:",omitempty"`
}

type Log struct {
	Invoice Invoice `json:"invoice"`
}
