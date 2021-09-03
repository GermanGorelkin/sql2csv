package sql2csv

import (
	"bufio"
	"database/sql"
	"io"
)

func NewCSVWriter(delimiter []byte, nl []byte, w io.Writer) CSVWriter {
	return CSVWriter{
		Delimiter: delimiter,
		NewLine:   nl,
		w:         bufio.NewWriter(w),
	}
}

type CSVWriter struct {
	Delimiter []byte
	NewLine   []byte
	w         *bufio.Writer
}

func (wr CSVWriter) WriteStrings(cols []string) error {
	for n, field := range cols {
		if n > 0 {
			if _, err := wr.w.Write(wr.Delimiter); err != nil {
				return err
			}
		}
		if _, err := wr.w.Write([]byte(field)); err != nil {
			return err
		}
	}
	if _, err := wr.w.Write(wr.NewLine); err != nil {
		return err
	}

	return nil
}

func (wr CSVWriter) Write(row []interface{}) error {
	for n, field := range row {
		if n > 0 {
			if _, err := wr.w.Write(wr.Delimiter); err != nil {
				return err
			}
		}
		if _, err := wr.w.Write(*(field.(*sql.RawBytes))); err != nil {
			return err
		}
	}
	if _, err := wr.w.Write(wr.NewLine); err != nil {
		return err
	}

	return nil
}

func (wr CSVWriter) Flush() error {
	return wr.w.Flush()
}

func (wr CSVWriter) AddBOM() error {
	if _, err := wr.w.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
		return err
	}
	return nil
}
