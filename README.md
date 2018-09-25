## Inspiration
Errors in smart contracts regularly cause millions of dollars in damage. Theorem proving software is incredibly powerful, but very underutilized. So what if we used a theorem prover to help prevent these bugs?

## What it does
Parses Lisp into an AST, performs type-checking, verifies assertions within the code using Z3, and compiles the AST to solidity.

## How we built it
We use Golang to parse Lisp, perform type-checking, and compilation down to solidity, and use Z3 to verify assertions within the code.

## Challenges we ran into
Compilers can be rather complicated and difficult to write. We have made significant progress on all of the pieces but still need to tie things completely together and handle fringe cases.

## Accomplishments that we're proud of
We made excellent progress

## What we learned
We learned about Lisp
We learned about Z3 and formal verification
We learned how to parse Lisp into an AST using Go
We learned how Go can be used to compile an AST into solidity

## What's next for Solispidy
Completing the compiler, and fleshing it out with more features, finishing tying things together and handling fringe cases.
