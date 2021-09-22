package assert_test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-assert"
)

func Test_FailRerun(t *testing.T) {
	{
		as := assert.New(t, assert.WithFailRerun(3))

		ins := new(failRerun)
		as.Run("", func(as *assert.Assertions) {
			as.True(ins.TrueWhenSecond())
		})
	}
}

type failRerun struct {
	i int
}

func (r *failRerun) TrueWhenSecond() bool {
	r.i++
	fmt.Println(r.i)
	if r.i <= 2 {
		return false
	}
	return true
}
