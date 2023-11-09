package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateNumber(t *testing.T) {
	num := generateNumber()
	fmt.Println("Generated number ", num)
	assert.GreaterOrEqual(t, num, 0)
	assert.LessOrEqual(t, num, 100)
}
