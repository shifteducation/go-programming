package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(buildWithBuilder())
	fmt.Println(buildWithJoin())
}

func buildWithBuilder() string {
	builder := strings.Builder{}

	builder.WriteRune('a')
	builder.WriteRune('b')
	builder.WriteRune('c')

	return builder.String()
}

func buildWithJoin() string {
	return strings.Join([]string{"a", "b", "c"}, "")
}
