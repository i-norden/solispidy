package parser_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLogKill(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Test Suite")
}
