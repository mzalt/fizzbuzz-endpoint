// fizzbuzz_test.go
package handler

import (
	"fmt"
	"testing"

	"github.com/fizzbuzz-endpoint/internal/model"
	"github.com/stretchr/testify/assert"
)

type Fixture struct {
	testCases     []*model.FizzbuzzInpute
	expectedCases []string
}

func TestMyFizzBuzz(t *testing.T) {

	fixture := prepareFixtures()
	assert := assert.New(t)
	for i := 0; i < len(fixture.testCases); i++ {
		res, err := MyFizzBuzz(fixture.testCases[i])
		assert.NoError(err)
		assert.Equal(fixture.expectedCases[i], res.Output)
	}

	_, err := MyFizzBuzz(&model.FizzbuzzInpute{Number1: 0, Number2: 5, Limit: 10, Str1: "fizz", Str2: "buzz"})
	assert.Equal(err, fmt.Errorf("bad values for number1 or/and number2"))
}

func prepareFixtures() *Fixture {
	return &Fixture{
		testCases: []*model.FizzbuzzInpute{
			{Number1: 3, Number2: 5, Limit: 10, Str1: "fizz", Str2: "buzz"},
			{Number1: 3, Number2: 5, Limit: 10, Str1: "fizz", Str2: ""},
			{Number1: 3, Number2: 5, Limit: -1, Str1: "fizz", Str2: "buzz"},
			{Number1: 3, Number2: 5, Limit: 0, Str1: "fizz", Str2: "buzz"},
		},
		expectedCases: []string{
			"1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz",
			"1, 2, fizz, 4, 5, fizz, 7, 8, fizz, 10",
			"1",
			"1",
		},
	}
}
