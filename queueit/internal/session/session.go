package session

import (
	"coding-lesson/internal/cookies"
	"coding-lesson/internal/utils"
	"net/http"
	"time"
)

type Session struct {
	id      string
	storage SessionStorage
	r       *http.Request
	w       http.ResponseWriter
	expiry  time.Time
}

const SessionKey = "session_id"
const SessionTtl = 1800 * time.Second

func NewSession(w http.ResponseWriter, r *http.Request, ss SessionStorage) *Session {
	session := &Session{r: r, w: w, storage: ss, expiry: time.Now().Add(SessionTtl)}
	session.id = session.prepareSessionId()
	return session
}

func (s Session) GetId() string {
	return s.id
}

func (s Session) GetCreatedSince() (duration time.Duration, success bool) {
	return s.storage.GetCreatedSince(s)
}

func (s Session) GetValue() (data interface{}, success bool) {
	return s.storage.Get(s)
}

func (s Session) SetValue(data interface{}) {
	s.storage.Set(s, data)
}

func (s Session) prepareSessionId() string {
	cookie, err := s.r.Cookie(SessionKey)
	var sessionId string

	if err == nil {
		sessionId = cookie.Value
	} else {
		sessionId = utils.RandomString(16)
		cookies.SetCookie(s.w, SessionKey, sessionId)
	}

	return sessionId
}
