package sql2csv

import (
	"bytes"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCSVWriter_Write(t *testing.T) {
	w := bytes.NewBuffer([]byte{})
	csvWriter := NewCSVWriter([]byte("|"), []byte("\r\n"), w)

	err := csvWriter.Write(genRow([]byte{1}, []byte{2}, []byte{3}))
	assert.Nil(t, err)
	assert.Equal(t, w.Bytes(), []byte{1, '|', 2, '|', 3, '\r', '\n'})

	err = csvWriter.Write(genRow([]byte{4}, []byte{5}, []byte{6}))
	assert.Nil(t, err)
	assert.Equal(t, w.Bytes(), []byte{
		1, '|', 2, '|', 3, '\r', '\n',
		4, '|', 5, '|', 6, '\r', '\n'})
}

func genRow(cols ...[]byte) []interface{} {
	var row []interface{}
	for _, col := range cols {
		data := sql.RawBytes(col)
		row = append(row, &data)
	}
	return row
}
