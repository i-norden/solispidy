package parser_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"

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

	expectedTokens := make([]types.Symbol, 12)
	expectedTokens[0] = types.LeftPar{LPId: 1, Line: 0}
	expectedTokens[1] = types.CnstStr{Data: "This", Line: 0}
	expectedTokens[2] = types.CnstStr{Data: "is", Line: 0}
	expectedTokens[3] = types.CnstStr{Data: "a", Line: 0}
	expectedTokens[4] = types.CnstStr{Data: "test", Line: 0}
	expectedTokens[5] = types.RightPar{RPId: 1, Line: 0}
	expectedTokens[6] = types.LeftPar{LPId: 1, Line: 2}
	expectedTokens[7] = types.CnstStr{Data: "split", Line: 2}
	expectedTokens[8] = types.CnstStr{Data: "me", Line: 2}
	expectedTokens[9] = types.CnstStr{Data: "up", Line: 2}
	ui := [4]uint64{3}
	expectedTokens[10] = types.CnstInt{Data: ui, Line: 2}
	expectedTokens[11] = types.RightPar{RPId: 1, Line: 2}

	mockProgram2 := "(this is (test))"

	line3 := parser.Line{
		Text:   "(this is (test))",
		Number: 0,
		Tokens: []string{"(", "this", "is", "(", "test", ")", ")"},
	}

	expectedTokens2 := make([]types.Symbol, 7)
	expectedTokens2[0] = types.LeftPar{LPId: 1, Line: 0}
	expectedTokens2[1] = types.CnstStr{Data: "this", Line: 0}
	expectedTokens2[2] = types.CnstStr{Data: "is", Line: 0}
	expectedTokens2[3] = types.LeftPar{LPId: 1, Line: 0}
	expectedTokens2[4] = types.CnstStr{Data: "test", Line: 0}
	expectedTokens2[5] = types.RightPar{RPId: 1, Line: 0}
	expectedTokens2[6] = types.RightPar{RPId: 1, Line: 0}

	var AST1, AST2, AST3, AST4 types.AST
	AST1.Here = &expectedTokens2[1]
	AST1.Next = &AST2
	AST2.Here = &expectedTokens2[2]
	AST2.Next = &AST3
	AST3.Here = &expectedTokens2[4]
	AST3.Next = &AST4

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

	It("Tests the MakeAST function", func() {
		lines, err := parser.Tokenize(mockProgram2)
		Expect(err).ToNot(HaveOccurred())
		Expect(*lines[0]).To(Equal(line3))

		tokens, err := parser.ReadFromLines(lines)
		Expect(err).ToNot(HaveOccurred())
		Expect(tokens).To(Equal(expectedTokens2))

		ast, err := parser.MakeAST(tokens, types.AST{}, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(ast).To(Equal(AST1))
		fmt.Fprintf(os.Stderr, "AST here: %v\r\n AST next here: %v\r\n", *ast.Here, *ast.Next.Here)
	})
})
