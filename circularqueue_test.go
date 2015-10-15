package circularqueue_test

import (
	"github.com/garyburd/redigo/redis"
	"github.com/octoblu/circularqueue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("circularqueue", func() {
	Context("Pop", func() {
		It("should have a Pop", func() {
			sut := circularqueue.New()

			job, err := sut.Pop()
			Expect(err).To(BeNil())
			Expect(job).NotTo(BeNil())
		})

		Context("When there are two records in redis", func() {
			var redisConn redis.Conn
			var result interface{}
			var sut circularqueue.Queue

			BeforeEach(func(){
				var err error

				redisConn, err = redis.Dial("tcp", ":6379")
				Expect(err).To(BeNil())

				_, err = redisConn.Do("DEL", "circular-job-queue")
				Expect(err).To(BeNil())

				_, err = redisConn.Do("RPUSH", "circular-job-queue", "A", "B")
				Expect(err).To(BeNil())

				sut = circularqueue.New()
				result, err = sut.Pop()
				Expect(err).To(BeNil())
			})

			AfterEach(func(){
				redisConn.Close()
			})

			It("should move the last record to the first record", func(){
				resp,err := redisConn.Do("LINDEX", "circular-job-queue", 0)

				if err != nil {
					Fail(err.Error())
				}

				record := string(resp.([]uint8))
				Expect(record).To(Equal("B"))
			})

			It("should maintain the same list length", func(){
				resp,err := redisConn.Do("LLEN", "circular-job-queue")

				if err != nil {
					Fail(err.Error())
				}

				Expect(resp).To(Equal(int64(2)))
			})

			It("should return the last record", func(){
				job := result.(circularqueue.Job)

				Expect(job.GetKey()).To(Equal("B"))
			})

			Context("when pop is called a second time", func(){
				BeforeEach(func(){
					var err error
					result, err = sut.Pop()
					Expect(err).To(BeNil())
				})

				It("should return the other job", func(){
					job := result.(circularqueue.Job)
					Expect(job.GetKey()).To(Equal("A"))
				})
			})
		})
	})
})
