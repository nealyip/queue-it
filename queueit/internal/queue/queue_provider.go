package queue

type QueueProvider interface {
	Push(key string, session string)
	Shift(key string) string
	ApproximateLength(key string) int
}
