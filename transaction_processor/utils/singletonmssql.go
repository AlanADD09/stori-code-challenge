package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"context"
	"database/sql"
	"encoding/json"
	"errors"

	_ "github.com/microsoft/go-mssqldb"
)

var mssql *sql.DB
var lock = &sync.Mutex{}

func StringConnection() (stringConnection string) {

	stringConnection = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable;",
		os.Getenv("MSSQL_HOST"), os.Getenv("MSSQL_USER"), os.Getenv("MSSQL_PASSWORD"), os.Getenv("MSSQL_PORT"), os.Getenv("MSSQL_NAME"))
	fmt.Println("Connection string: ", stringConnection)
	return

}

func poolConnection() (conn *sql.DB, err error) {

	// Create connection pool
	conn, err = sql.Open("sqlserver", StringConnection())
	if err != nil {
		log.Fatalf("\nError creating connection pool: %v\n", err.Error())
	}
	return

}

func getInstanceMssql() (*sql.DB, error) {
	if mssql == nil {
		lock.Lock()
		defer lock.Unlock()
		if mssql == nil {
			conn, err := poolConnection()
			if err != nil {
				return nil, err
			}
			mssql = conn
		}
	}
	return mssql, nil
}

type SqlArgs struct {
	Name  string
	Value interface{}
}

type AffectsRows struct {
	AffectsRows int     `json:"affects_rows"`
	Error       *string `json:"error,omitempty"`
	ID          *int    `json:"id,omitempty"`
}

func parseSqlArgsMssql(variables []SqlArgs) (args []sql.NamedArg) {
	for _, value := range variables {
		args = append(args, sql.Named(value.Name, value.Value))
	}
	return
}

func DoQuery(query string, variables []SqlArgs) (jsonData []byte, err error) {
	var stmt *sql.Stmt
	var rows *sql.Rows
	var columnTypes []*sql.ColumnType
	connection, _ := getInstanceMssql()
	ctx := context.Background()

	args := make([]interface{}, 0)
	for _, value := range variables {
		args = append(args, sql.Named(value.Name, value.Value))
	}

	if stmt, err = connection.Prepare(query); err == nil {
		defer stmt.Close()

		// rows, err = stmt.QueryContext(ctx, parseSqlArgsMssql(variables))
		if rows, err = stmt.QueryContext(ctx, args...); err != nil {
			return
		}

		columnTypes, err = rows.ColumnTypes()
		if err != nil {
			return
		}

		count := len(columnTypes)
		finalRows := []interface{}{}

		for rows.Next() {

			scanArgs := make([]interface{}, count)

			for i, v := range columnTypes {
				switch v.DatabaseTypeName() {
				case "DATE", "DATETIMEOFFSET", "DATETIME2", "SMALLDATETIME", "DATETIME", "TIME", "CHAR", "NCHAR", "VARCHAR", "TEXT", "NTEXT":
					scanArgs[i] = new(sql.NullString)
					break
				case "NVARCHAR":
					scanArgs[i] = new([]byte)
				case "BIT":
					scanArgs[i] = new(sql.NullBool)
					break
				case "BIGINT", "NUMERIC", "SMALLINT", "DECIMAL", "SMALLMONEY", "INT":
					scanArgs[i] = new(sql.NullInt64)
					break
				case "FLOAT", "REAL":
					scanArgs[i] = new(sql.NullFloat64)
					break
				default:
					fmt.Printf("default")
					scanArgs[i] = new(sql.NullString)
				}
			}
			rows.Scan()

			err = rows.Scan(scanArgs...)

			if err != nil {
				return
			}

			// var masterData map[string]interface{}
			masterData := map[string]interface{}{}

			for i, v := range columnTypes {

				if fmt.Sprintf("%T", scanArgs[i]) == "*[]uint8" || fmt.Sprintf("%T", scanArgs[i]) == "*[]byte" {
					var aux interface{}
					json.Unmarshal((*(scanArgs[i]).(*[]byte)), &aux)
					masterData[v.Name()] = aux
					continue
				}
				if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
					if z.Valid {
						masterData[v.Name()] = z.Bool
					}
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullString); ok {
					if z.Valid {
						masterData[v.Name()] = z.String
					}
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
					if z.Valid {
						masterData[v.Name()] = z.Int64
					}
					continue
				} else {
				}

				if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
					if z.Valid {
						masterData[v.Name()] = z.Float64
					}
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
					if z.Valid {
						masterData[v.Name()] = z.Int32
					}
					continue
				}
				masterData[v.Name()] = scanArgs[i]

			}

			finalRows = append(finalRows, masterData)
		}

		jsonData, err = json.Marshal(finalRows)
	}
	return
}

func DoMutation(query string, variables []SqlArgs) (result int64, err error) {
	connection, _ := getInstanceMssql()
	ctx := context.Background()

	args := make([]interface{}, 0)
	for _, value := range variables {
		args = append(args, sql.Named(value.Name, value.Value))
	}

	stmt, err := connection.Prepare(query)
	if err != nil {
		strError := fmt.Sprintf("%s error Prepare sql error", query)
		e := errors.New(strError)
		return result, e
	}
	defer stmt.Close()

	rows_afftected, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}
	result, err = rows_afftected.RowsAffected()
	return
}
