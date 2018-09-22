package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/i-norden/solispidy/parser"
)

var _ = Describe("Parser", func() {

	mockProgram := `This is a test
						; reject this
						(split me up 3)`
	line1 := parser.Line{
		Text: "This is a test",
		Number: 1,
		Tokens: []string{"This","is","a","test"},
	}
	line2 := parser.Line{
		Text: "(split me up 3)",
		Number: 3,
		Tokens: []string{"(","split","me","up","3",")"},
	}

	It("Tests the Tokenize function", func() {
		lines := parser.Tokenize(mockProgram)
		Expect(lines[0]).To(Equal(line1))
		Expect(lines[1]).To(Equal(line2))
	})
})