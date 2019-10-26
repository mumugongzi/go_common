package doc

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReadText(t *testing.T) {
	file := "/Users/lisong/Downloads/simple.docx"

	res, err := ReadText(file)
	assert.Nil(t, err)
	t.Log(strings.Join(res, "\n"))
}
