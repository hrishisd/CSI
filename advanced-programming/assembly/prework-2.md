**Readings.** Begin by reading the following for background context on Go assembly:

* [A Primer on Go Assembly](https://github.com/teh-cmc/go-internals/tree/master/chapter1_assembly_primer)
* [A Quick Guide to Go's Assembler](https://golang.org/doc/asm)

Next, read the following:

* [Proposal: Register-based Go calling convention](https://github.com/golang/proposal/blob/master/design/40724-register-calling.md)
* [Go internal ABI specification](https://go.googlesource.com/go/+/refs/heads/dev.regabi/src/cmd/compile/internal-abi.md)

As you read, keep in mind the following questions:

* In what ways is Go assembly a "pseudo-assembly"? How is it different from something like x86-64?
* Where does Go's assembler fit into the build pipeline?
* What details are specified by an "Application Binary Interface", and what specific decisions does the Go ABI make?
* How do the two (stack-based or register-based) calling conventions handle Go's multiple return values?
* Can you experimentally check (e.g. by looking at assembler output) whether a locally-compiled program uses a stack-based or register-based calling convention?

**Further Resources.**

[Go 1.17 adds support for the new register-based calling convention](https://golang.org/doc/go1.17#compiler). However, there are still many open tasks, e.g. adding more fuzz testing, porting this change to more platforms; see [issue #40724 in the golang/go repo](https://github.com/golang/go/issues/40724) for a comprehensive task list.

Are you interested in getting involved with this effort? If so, you might find it useful to read Go's [Contribution Guide](https://golang.org/doc/contribute) and to follow along with the discussion for #40724, so as to get a sense of where and how you might be able to contribute.
