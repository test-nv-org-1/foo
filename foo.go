package foo

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

const version = "v0.1.0"

func Foo() string {
	return fmt.Sprintf("foo-%v %s", strconv.FormatUint(rand.Uint64(), 36), version)
}
