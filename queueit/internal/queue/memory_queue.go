package queue

import "log"

// MemoryQueue Define a struct to represent the memory queue
type MemoryQueue struct {
	queues map[string][]string
}

func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{queues: make(map[string][]string)}
}

// Push Implement the Push method for MemoryQueue
func (mq *MemoryQueue) Push(key string, item string) {
	mq.queues[key] = append(mq.queues[key], item)
	log.Printf("Data in q %s", mq.queues[key])
}

// Shift Implement the Shift method for MemoryQueue
func (mq *MemoryQueue) Shift(key string) string {
	if len(mq.queues[key]) == 0 {
		return ""
	}

	firstItem := mq.queues[key][0]
	mq.queues[key] = mq.queues[key][1:]
	return firstItem
}

func (mq *MemoryQueue) ApproximateLength(key string) int {
	return len(mq.queues[key])
}
