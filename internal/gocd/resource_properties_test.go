package gocd

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestResourceProperties(t *testing.T) {
	t.Run("Basic", testResourcePropertiesBasic)
	t.Run("NewFrame", testResourcePropertiesNewFrame)
	t.Run("MarshallCSV", testResourcePropertiesMarshallCSV)
	t.Run("UnmarshallCSVWithHeader", testResourcePropertiesUnmarshallCSVWithHeader)
	t.Run("UnmarshallCSVWithoutHeader", testResourcePropertiesUnmarshallCSVWithoutHeader)
	t.Run("IOWriterInterface", testResourcePropertiesIOWriter)
}

func testResourcePropertiesIOWriter(t *testing.T) {
	var w1 io.Writer
	var p1 Properties
	p1 = Properties{}
	w1 = &p1
	w1.Write([]byte("one"))
	assert.Equal(t, "one", p1.DataFrame[0][0])
}

func testResourcePropertiesBasic(t *testing.T) {
	p1 := Properties{}
	p1.SetRow(0, []string{"one", "two"})
	assert.Equal(t, []string{"one", "two"}, p1.DataFrame[0])
	assert.Len(t, p1.DataFrame, 1)

	p2 := Properties{
		Header: []string{"1", "2"},
	}
	p2.SetRow(4, []string{"one", "two"})
	assert.Equal(t, []string{"one", "two"}, p2.DataFrame[4])
	assert.Len(t, p2.DataFrame, 5)
	p2.SetRow(0, []string{"three", "four"})
	assert.Equal(t, "one", p2.Get(4, "1"))
}

func testResourcePropertiesNewFrame(t *testing.T) {
	p := NewPropertiesFrame([][]string{
		{"1", "2"},
		{"one", "two"},
		{"three", "four"},
	})

	assert.Equal(t, p.Header, []string{"1", "2"})
	assert.Equal(t, []string{"one", "two"}, p.DataFrame[0])
	assert.Equal(t, []string{"three", "four"}, p.DataFrame[1])
}

func testResourcePropertiesMarshallCSV(t *testing.T) {
	p := NewPropertiesFrame([][]string{
		{"1", "2"},
		{"one", "two"},
		{"thr,ee", "fo\"ur"},
	})

	raw, err := p.MarshallCSV()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, `1,2
one,two
"thr,ee","fo""ur"
`, raw)
}

func testResourcePropertiesUnmarshallCSVWithHeader(t *testing.T) {
	p := Properties{
		UnmarshallWithHeader: true,
	}
	err := p.UnmarshallCSV(`1,2
one,two
"thr,ee","fo""ur"
`)
	assert.True(t, p.UnmarshallWithHeader)

	assert.Equal(t, []string{"1", "2"}, p.Header)
	assert.Equal(t, []string{"one", "two"}, p.DataFrame[0])
	assert.Equal(t, []string{"thr,ee", "fo\"ur"}, p.DataFrame[1])
	assert.Nil(t, err)
}

func testResourcePropertiesUnmarshallCSVWithoutHeader(t *testing.T) {
	p := Properties{}
	err := p.UnmarshallCSV(`1,2
one,two
"thr,ee","fo""ur"
`)

	assert.False(t, p.UnmarshallWithHeader)
	assert.Empty(t, p.Header)
	assert.Equal(t, []string{"1", "2"}, p.DataFrame[0])
	assert.Equal(t, []string{"one", "two"}, p.DataFrame[1])
	assert.Equal(t, []string{"thr,ee", "fo\"ur"}, p.DataFrame[2])
	assert.Nil(t, err)
}
