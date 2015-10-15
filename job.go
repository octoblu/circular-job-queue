package circularqueue

// Job represents a job in the queue
type Job interface {
  GetKey() string
}

// KeyJob implements Job
type KeyJob struct {
  key string
}

// NewJob returns a new job
func NewJob(key string) *KeyJob {
  return &KeyJob{key: key}
}

// GetKey returns the key to store the information about the job
func (job *KeyJob) GetKey() string {
  return job.key
}
