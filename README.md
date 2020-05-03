# numtowords

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/parithiban/numtowords)
[![Go Report Card](https://goreportcard.com/badge/github.com/parithiban/numtowords)](https://goreportcard.com/report/github.com/parithiban/numtowords)
[![Build Status](https://travis-ci.org/parithiban/numtowords.svg?branch=master)](https://travis-ci.org/parithiban/numtowords)
[![codecov.io](https://codecov.io/github/parithiban/numtowords/coverage.svg?branch=master)](https://codecov.io/github/parithiban/numtowords?branch=master)

[![GitHub contributors](https://img.shields.io/github/contributors/parithiban/numtowords.svg?style=plastic&color=blue)](https://GitHub.com/parithiban/numtowords/graphs/contributors/)
![Last Commit](https://img.shields.io/github/last-commit/parithiban/numtowords.svg?style=plastic)

This is a utility package written in go which is used to convert the number to words.

## Installation

This requires Go version 1.11 or later.

> go get -u github.com/parithiban/numtowords

## Example

```code
words  := numtowords.Transform(1234) //output: "one thousand two hundred thirty four"
words  := numtowords.Transform(100000000000) //output: "one hundred billion"
words  := numtowords.Transform(000000000) //output: "zero"
words  := numtowords.Transform(1100100100300800112) //output: "one quintillion one hundred quadrillion one hundred trillion one hundred billion three hundred million eight hundred thousand one hundred twelve"
```
[Playground Link](https://play.golang.org/p/7XhcBcMedCh)
