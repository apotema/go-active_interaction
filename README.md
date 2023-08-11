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

func (s DoubleA) Run() int {
	return s.A * 2
}

func main() {
	value, _ := Execute[int](&DoubleA{A: 2})
	fmt.Println(value)
}
```

For each the output would be 4.

### Validations



## Credits

ActiveInteraction is brought to you by [Alexandre Mondaini Calvão][].
Special thanks to [Aaron Lasseigne][] for creating the [ActiveInteraction][] gem.


If you want to contribute to ActiveInteraction, please read
[our contribution guidelines][]. A [complete list of contributors][] is
available on GitHub.

ActiveInteraction is licensed under [the MIT License][].

[go active interaction]: https://github.com/apotema/go-active-interaction
[activeinteraction]: https://github.com/AaronLasseigne/active_interaction
[API Documentation]: https://google.com
[alexandre mondaini calvão]: https://github.com/apotema
[aaron lasseigne]: https://github.com/AaronLasseigne
[our contribution guidelines]: CONTRIBUTING.md
[complete list of contributors]: https://google.com
[the MIT License]: LICENSE.md

<!-- ?? -->
[validator]: https://github.com/go-playground/validator
