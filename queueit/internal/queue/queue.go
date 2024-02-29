package queue

type Queue struct {
	provider QueueProvider
}

func NewQueue(provider QueueProvider) *Queue {
	return &Queue{provider: provider}
}

func (q Queue) SetProvider(provider QueueProvider) {
	q.provider = provider
}

func (q Queue) Push(key string, session string) {
	q.provider.Push(key, session)
}

func (q Queue) Shift(key string) string {
	return q.provider.Shift(key)
}

func (q Queue) ApproximateLength(key string) int {
	return q.provider.ApproximateLength(key)
}
