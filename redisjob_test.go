package circularqueue_test

import (
	"github.com/octoblu/circularqueue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Redisjob", func() {
	Context("New with a key", func() {
		var sut circularqueue.Job

		BeforeEach(func(){
			sut = circularqueue.NewJob("old-map")
		})

		It("should set the key", func(){
			Expect(sut.GetKey()).To(Equal("old-map"))
		})
	})

})
