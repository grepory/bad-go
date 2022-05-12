# Bad Go examples

I've tried to put together a basic, approachable tutorial to some common
mistakes I've seen people make in Go with some general advice. My examples
and organization aren't as good as some publicly available tutorials and
documentation, but I thought it would be fun to put together, and it was
a good refresher for me.

## How to read this code

Each of the packages in this repo contains some example code. It contains
one or more sets of files of the following pattern:

| FOO.go | The bad example |
| FOO_test.go | Tests that illustrate the bad and good |
| better_FOO_test.go | Implementation of the working example |

Read them in that order. They create a bit of a narrative.

## Better resources than this

The [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) is one of my favorite
go-to docs. It covers a lot more than what's here, and what it covers 
includes some examples that are much more advanced than what I do.
Should you approach that guide, feel free to reach out to chat about
anything in it.

That guide also contains some really good links itself:

[Effective Go](https://go.dev/doc/effective_go)
[Golang Common Mistakes](https://github.com/golang/go/wiki/CommonMistakes)
[Common Golang Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)