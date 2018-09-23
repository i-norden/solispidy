package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/i-norden/solispidy/parser"
	"github.com/i-norden/solispidy/types"
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

	expectedTokens := make([]types.Symbol, 8)
	expectedTokens[0] = types.CnstStr{Data: "This", Line: 0}
	expectedTokens[1] = types.CnstStr{Data: "is", Line: 0}
	expectedTokens[2] = types.CnstStr{Data: "a", Line: 0}
	expectedTokens[3] = types.CnstStr{Data: "test", Line: 0}
	expectedTokens[4] = types.CnstStr{Data: "split", Line: 2}
	expectedTokens[5] = types.CnstStr{Data: "me", Line: 2}
	expectedTokens[6] = types.CnstStr{Data: "up", Line: 2}
	ui := [4]uint64{3}
	expectedTokens[7] = types.CnstInt{Data: ui, Line: 2}

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
		Expect(err).ToNot(HaveOccurred())
		Expect(tokens).To(Equal(expectedTokens))
	})
})
