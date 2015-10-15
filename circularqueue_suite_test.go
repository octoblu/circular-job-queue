package circularqueue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCircularqueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Circularqueue Suite")
}
