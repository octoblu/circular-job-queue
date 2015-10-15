package circularqueue

import (
	"github.com/garyburd/redigo/redis"
)

// Queue represents the underlying job queue
type Queue struct {

}

// New constructs a new instance of the circular queue
func New() (Queue) {
  var queue Queue
  return queue
}

// Pop returns a job from the queue and auto-readds it
func (queue *Queue) Pop() (Job,error) {
  var err error
  var redisConn redis.Conn
  var result interface{}

  redisConn, err = redis.Dial("tcp", ":6379")
  if err != nil {
    return nil, err
  }

  result, err = redisConn.Do("RPOPLPUSH", "circular-job-queue", "circular-job-queue")
  if err != nil {
    return nil, err
  }

  key := string(result.([]uint8))
  return NewJob(key), nil
}
