package session

import (
	"time"
)

type MemorySessionStorage struct {
	sessions map[string]storage
}

type storage struct {
	createdAt time.Time
	expiry    time.Time
	value     interface{}
}

func NewMemorySessionStorage() *MemorySessionStorage {
	return &MemorySessionStorage{sessions: make(map[string]storage)}
}

func (m MemorySessionStorage) Set(session Session, data interface{}) {
	m.sessions[session.GetId()] = storage{value: data, createdAt: time.Now(), expiry: session.expiry}
}

func (m MemorySessionStorage) Get(session Session) (data interface{}, success bool) {
	if stored, ok := m.sessions[session.GetId()]; ok {
		if time.Now().Before(stored.expiry) {
			return stored.value, true
		}
		delete(m.sessions, session.GetId())
	}
	return nil, false
}
func (m MemorySessionStorage) GetCreatedSince(session Session) (duration time.Duration, success bool) {
	if session, ok := m.sessions[session.GetId()]; ok {
		return time.Now().Sub(session.createdAt), true
	}
	return 0, false
}
