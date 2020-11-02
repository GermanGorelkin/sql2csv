package sql2csv

import (
	"context"
	"database/sql"
	"io"
)

func NewCSVWriter(delimiter []byte, nl []byte, w io.Writer) CSVWriter {
	return CSVWriter{
		Delimiter: delimiter,
		NewLine:   nl,
		w:         w,
	}
}

type CSVWriter struct {
	Delimiter []byte
	NewLine   []byte
	w         io.Writer
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

func NewSQLReader(db *sql.DB) SQLReader {
	return SQLReader{DB: db}
}
type SQLReader struct {
	DB *sql.DB
}

func (pool SQLReader) Read(ctx context.Context, query string, w CSVWriter) error {
	if err := pool.DB.PingContext(ctx); err != nil {
		return err
	}

	rows, err := pool.DB.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	vals := make([]interface{}, len(cols))
	for i := range vals {
		vals[i] = new(sql.RawBytes)
	}

	for rows.Next() {
		if err = rows.Scan(vals...); err != nil {
			return err
		}
		if err = w.Write(vals); err != nil {
			return err
		}
	}
	if err = rows.Close(); err != nil {
		return err
	}
	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
