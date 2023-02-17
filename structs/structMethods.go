package structs

import (
	"fmt"
)

type ownStringImpl struct {
	value []byte
	len   int
}

func testMethods() {

	var normalString string
	var ownString *ownStringImpl

	normalString = "test"
	ownString = newOwnString([]byte("own implementation"))

	fmt.Printf("The normal string result is: %s and the own implementation is %v, returning value %s and len %d\n", normalString, ownString, ownString.String(), ownString.Len())
}

func newOwnString(input []byte) *ownStringImpl {
	return &ownStringImpl{value: input, len: len(input)}
}

// implementing it automatically returns string when parsing a value
func (s *ownStringImpl) String() string {
	return string(s.value)
}

func (s *ownStringImpl) Len() int {
	return s.len
}
