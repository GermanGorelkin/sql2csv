package sql2csv

import (
	"context"
	"database/sql"
)

type Writer interface {
	WriteStrings(cols []string) error
	Write(row []interface{}) error
	Flush() error
}

func NewSQLReader(db *sql.DB) SQLReader {
	return SQLReader{DB: db}
}

type SQLReader struct {
	DB      *sql.DB
	Columns bool
}

func (pool SQLReader) Read(ctx context.Context, query string, w Writer) error {
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

	if pool.Columns {
		if err = w.WriteStrings(cols); err != nil {
			return err
		}
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

	return w.Flush()
}
