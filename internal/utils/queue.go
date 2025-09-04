package utils

import (
	"strconv"
	"strings"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/alfariiizi/vandor/internal/enum"
)

func GetQueueName(queue enum.Queue) string {
	name, _ := getQueueConfig(queue)
	return name
}

func GetQueuePriority(queue enum.Queue) int {
	_, priority := getQueueConfig(queue)
	return priority
}

func GetQueue(queue enum.Queue) (string, int) {
	return getQueueConfig(queue)
}

func GetAllQueueConfig() map[string]int {
	queues := make(map[string]int)
	for _, queue := range enum.AllQueues {
		name, priority := getQueueConfig(queue)
		queues[name] = priority
	}
	return queues
}

func getQueueConfig(queue enum.Queue) (string, int) {
	cfg := config.GetConfig()
	prefix := cfg.Jobs.QueuePrefix

	queueOne := strings.Split(queue.Label(), ":")
	if len(queueOne) < 2 {
		return prefix + queueOne[0], 1
	}
	priority, err := strconv.Atoi(queueOne[1])
	if err != nil {
		priority = 1
	}
	return prefix + "_" + queueOne[0], priority
}
