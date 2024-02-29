package models

type Event struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Concurrency         int    `json:"concurrency"`
	SessionExpirySecond int    `json:"sessionExpirySecond"`
	ShowQueueLength     bool   `json:"showQueueLength"`
}

func (e Event) GetConcurrency() int {
	return e.Concurrency
}

func (e Event) GetSessionExpirySecond() int {
	return e.SessionExpirySecond
}

func (e Event) GetId() string {
	return e.ID
}

func (e Event) GetName() string {
	return e.Name
}

func (e Event) GetShowQueueLength() bool {
	return e.ShowQueueLength
}
