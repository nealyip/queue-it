package session

import (
	"time"
)

type SessionStorage interface {
	Set(session Session, data interface{})
	Get(session Session) (data interface{}, success bool)
	GetCreatedSince(session Session) (duration time.Duration, success bool)
}
