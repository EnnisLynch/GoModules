package DefaultValues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DefaultTest struct {
	Foo        string `default:"bar"`
	Rock       string `default:"star"`
	Empty      string
	CheckInt   int   `default:"-42"`
	CheckInt16 int16 `default:"-42"`
	CheckInt32 int32 `default:"-42"`
	CheckInt64 int64 `default:"-42"`

	CheckUInt   uint   `default:"84"`
	CheckUInt16 uint16 `default:"84"`
	CheckUInt32 uint32 `default:"84"`
	CheckUInt64 uint64 `default:"84"`

	CheckUInt_Bad   uint   `default:"-1"`
	CheckUInt16_Bad uint16 `default:"-1"`
	CheckUInt32_Bad uint32 `default:"-1"`
	CheckUInt64_Bad uint64 `default:"-1"`
}

func Test_DefaultTest_strings(t *testing.T) {
	assert := assert.New(t)

	var new_object *DefaultTest = New[DefaultTest]()

	assert.Equal(new_object.Foo, "bar", "Foo=bar")
	assert.Equal(new_object.Rock, "star", "Rock=star")
	assert.Equal(new_object.Empty, "", "Empty no default")

}

func Test_DefaultTest_signed_integers(t *testing.T) {
	assert := assert.New(t)

	var new_object *DefaultTest = New[DefaultTest]()

	assert.Equal(int(-42), new_object.CheckInt, "42 is the answer")
	assert.Equal(int16(-42), new_object.CheckInt16, "42 is the answer")
	assert.Equal(int32(-42), new_object.CheckInt32, "42 is the answer")
	assert.Equal(int64(-42), new_object.CheckInt64, "42 is the answer")
}
func Test_DefaultTest_unsigned_integers(t *testing.T) {
	assert := assert.New(t)

	var new_object *DefaultTest = New[DefaultTest]()

	assert.Equal(uint(84), new_object.CheckUInt, "CheckUInt")
	assert.Equal(uint16(84), new_object.CheckUInt16, "CheckUInt16")
	assert.Equal(uint32(84), new_object.CheckUInt32, "CheckUInt32")
	assert.Equal(uint64(84), new_object.CheckUInt64, "CheckUInt64")

	assert.Equal(uint(0), new_object.CheckUInt_Bad, "CheckUInt_Bad")
	assert.Equal(uint16(0), new_object.CheckUInt16_Bad, "CheckUInt16_Bad")
	assert.Equal(uint32(0), new_object.CheckUInt32_Bad, "CheckUInt32_Bad")
	assert.Equal(uint64(0), new_object.CheckUInt64_Bad, "CheckUInt64_Bad")
}
