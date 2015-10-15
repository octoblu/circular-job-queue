package circularqueue

// Queue represents the underlying job queue
type Queue interface {
  Pop() (Job,error)
}

// Conn is what all circular queues require
type Conn interface {
	Do(commandName string, args... interface{}) (interface{}, error)
}

// ConnQueue represents the underlying job queue
type ConnQueue struct {
	conn Conn
}

// New constructs a new instance of the circular queue
func New(conn Conn) (*ConnQueue) {
	return &ConnQueue{conn: conn}
}

// Pop returns a job from the queue and auto-readds it
func (queue *ConnQueue) Pop() (Job,error) {
  result, err := queue.conn.Do("RPOPLPUSH", "circular-job-queue", "circular-job-queue")
  if err != nil {
    return nil, err
  }

  key := string(result.([]uint8))
  return NewJob(key), nil
}
