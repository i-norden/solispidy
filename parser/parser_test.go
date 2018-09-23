package parser_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/i-norden/solispidy/parser"
)

var _ = Describe("Parser", func() {

	//var expectedTokens []interface{}

	mockProgram := "(This is a test)\n; reject this\n(split    me up 3)"
	line1 := parser.Line{
		Text:   "(This is a test)",
		Number: 0,
		Tokens: []string{"(", "This", "is", "a", "test", ")"},
	}
	line2 := parser.Line{
		Text:   "(split    me up 3)",
		Number: 2,
		Tokens: []string{"(", "split", "me", "up", "3", ")"},
	}

	It("Tests the Tokenize function", func() {
		lines, err := parser.Tokenize(mockProgram)
		Expect(err).ToNot(HaveOccurred())
		Expect(*lines[0]).To(Equal(line1))
		Expect(*lines[1]).To(Equal(line2))
	})

	It("Tests the ReadFromLiens and ReadFromTokens function", func() {
		lines, err := parser.Tokenize(mockProgram)
		Expect(err).ToNot(HaveOccurred())
		Expect(*lines[0]).To(Equal(line1))
		Expect(*lines[1]).To(Equal(line2))

		tokens, err := parser.ReadFromLines(lines)
		fmt.Fprintf(os.Stderr, "tokens: %v\r\n", tokens)
		Expect(err).ToNot(HaveOccurred())
	})
})
