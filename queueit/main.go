package main

import (
	"coding-lesson/internal/models"
	"coding-lesson/internal/queue"
	session2 "coding-lesson/internal/session"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	queue          *queue.Queue
	sessionStorage session2.SessionStorage
}

func main() {
	var eventRegistry = models.NewEventRegistry()
	cfg := &config{}
	cfg.queue = queue.NewQueue(queue.NewMemoryQueue())
	cfg.sessionStorage = session2.NewMemorySessionStorage()

	// Create a new instance of the HTTP router (ServeMux)
	mux := http.NewServeMux()

	// Define the root handler function
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		eventId := r.URL.Query().Get("id")
		event, exists := eventRegistry.GetEvent(eventId)
		if !exists {
			_, _ = w.Write([]byte(`<p>Event not found</p>`))
			return
		}
		handleQueue(w, r, event, cfg)
	}

	// Register the middleware and handlers to their respective routes
	mux.HandleFunc("/", rootHandler)

	// Start the HTTP server
	log.Println("Server listening on port 8081...")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handleQueue(w http.ResponseWriter, r *http.Request, event models.Event, cfg *config) {
	session := session2.NewSession(w, r, cfg.sessionStorage)
	value, sesErr := session.GetValue()
	if sesErr == false {
		log.Printf("unknown error fetching session %s", sesErr)
	}
	// TODO:
	// if session.GetId() not in the allowed list
	// check if it's time to allow one more or refresh allow list (including expiry mechanism )
	// if session.GetId() in the allowed list
	// redirect him to event page (with a signed jwt token)
	if value == nil {
		cfg.queue.Push(event.GetId(), session.GetId())
		session.SetValue("1")
	}

	html := fmt.Sprintf(`<p>You are in the queue for %s</p>`, event.GetName())

	if event.ShowQueueLength {
		html = html + fmt.Sprintf(`<p>There are %d people waiting</p>`, cfg.queue.ApproximateLength(event.GetId()))
	}

	_, err := w.Write([]byte(html))

	if err != nil {
		return
	}
}
