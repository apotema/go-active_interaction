 # [Go Active Interaction][]
================================

Go ActiveInteractation is based on [activeinteraction][] active_interaction and aims to bring the power of the Ruby gem to the Go space.


- [Installation](#installation)
- [Basic usage](#basic-usage)
  - [Validations](#validations)
- [Advanced usage](#advanced-usage)
  - [Callbacks](#callbacks)
  - [Composition](#composition)
  - Defaults - Not implemented yet
  - [Errors](#errors)
- [Credits](#credits)


[API Documentation][]


## Installation

Run the following Go command to install the gin package:

``` sh
go get github.com/apotema/go-activeinteraction
```

## Basic usage

First you need to import Go Active Interaction package for using Go Active Interaction, one simplest example likes the follow example.go:

```go
package main

import (
	"fmt"

	. "github.com/apotema/go-active_interaction/active_interaction"
)

type DoubleA struct {
	A int
	InteractionUtils
}

func (d DoubleA) Run() int {
	return d.A * 2
}

func main() {
	value, _ := Execute[int](&DoubleA{A: 2})
	fmt.Println("Interaction result: ", value)
}
```

For each the output would be:

``` sh
Interaction result:  4
```

### Validations

Just like in Ruby`s Active Interaction gem, you can add validations to your Interaction Arguments.
[Go Active Interaction][] uses some of the [Package validator][] validations
A simple example of adding validation to your Interaction is as the following:
```go
package main

import (
	"fmt"

	. "github.com/apotema/go-active_interaction/active_interaction"
)

type DoubleAIfBiggerThan10 struct {
	A int `validate:"gte=10"`
	InteractionUtils
}

func (d DoubleAIfBiggerThan10) Run() int {
	return d.A * 2
}

func main() {
	value, error := Execute[int](&DoubleAIfBiggerThan10{A: 2})

	if error.HasError() {
		fmt.Println("Interaction error: ", error)
	} else {
		fmt.Println("Interaction result: ", value)
	}
}
```

For each the output would be:

``` sh
description: Field validation for 'A' failed on the 'gte' tag
Interaction error:  "A": ["Field validation for 'A' failed on the 'gte' tag"]
```

Using the same code snippet, if we fulfill the validations passing A = 11:

```go
func main() {
	value, error := Execute[int](&DoubleAIfBiggerThan10{A: 11})

	if error.HasError() {
		fmt.Println("Interaction error: ", error)
	} else {
		fmt.Println("Interaction result: ", value)
	}
}
```

The return won't display an error:

``` sh
Interaction result:  22
```

## Credits

ActiveInteraction is brought to you by [Alexandre Mondaini Calvão][].
Special thanks to [Aaron Lasseigne][] for creating the [ActiveInteraction][] gem
and for [Dean Karn][] for creating the [Package validator][] lib.


If you want to contribute to ActiveInteraction, please read
[our contribution guidelines][]. A [complete list of contributors][] is
available on GitHub.

ActiveInteraction is licensed under [the MIT License][].

[go active interaction]: https://github.com/apotema/go-active-interaction
[activeinteraction]: https://github.com/AaronLasseigne/active_interaction
[API Documentation]: https://google.com
[alexandre mondaini calvão]: https://github.com/apotema
[aaron lasseigne]: https://github.com/AaronLasseigne
[dean karn]: https://github.com/deankarn
[package validator]: https://github.com/go-playground/validator
[our contribution guidelines]: CONTRIBUTING.md
[complete list of contributors]: https://google.com
[the MIT License]: LICENSE.md
