package tasktype

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"time"

	_ "github.com/go-sql-driver/mysql" // registry sqlite3 deive

	"github.com/labulaka521/crocodile/common/log"
	"go.uber.org/zap"
)

var _ TaskRuner = DataAPI{}

// DataAPI http req task
type DataSQL struct {
	DataType string   `json:"datatype" comment:"DataType"`
	Host     string   `json:"host" comment:"HOST"`
	Port     string   `json:"port" comment:"Port"`
	Opentx   bool     `json:"opentx" comment:"Opentx"`
	DataBase string   `json:"database" comment:"DataBase"`
	User     string   `json:"user" comment:"User"`
	PassWord string   `json:"password" comment:"PassWord"`
	SqlData  []string `json:"sqldata" comment:"SqlData"`
}

// Header
// Body
// Test

// Type return api
func (da DataSQL) Type() string {
	return "sql"
}

// Run implment TaskRun interface
func (da DataSQL) Run(ctx context.Context) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		var exitCode = DefaultExitCode
		defer pw.Close()
		defer func() {
			writeLog(pw, fmt.Sprintf("Run Finished,Return Code:%5d", exitCode))
		}()
		// 初始化连接
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", da.User, da.PassWord, da.Host, da.Port, da.DataBase))
		if err != nil {
			log.Error("Open Data Base failed", zap.Error(err))
			writeLog(pw, fmt.Sprintf("Open Data Base failed: %s", zap.Error(err)))
			return
		}
		defer db.Close()

		if da.Opentx {
			// start tx
			tx, err := db.Begin()
			if err != nil {
				log.Error("Open Tx failed", zap.Error(err))
				writeLog(pw, fmt.Sprintf("Open Tx failed: %s", zap.Error(err)))
				return
			}
			for i, sqlString := range da.SqlData {
				sqlRes, err := tx.Exec(sqlString)
				if err != nil {
					log.Error("Exec SQL With Tx failed ", zap.Error(err))
					return
				}
				rows, _ := sqlRes.RowsAffected()
				writeLog(pw, fmt.Sprintf("Run Sql %d With Tx Success,Number of rows affected :%d", i, rows))
			}
			err = tx.Commit()
			if err != nil {
				log.Error("Tx Commit failed", zap.Error(err))
				writeLog(pw, fmt.Sprintf("Tx Commit failed: %s", zap.Error(err)))
				return
			}
		} else {
			for i, sqlString := range da.SqlData {
				sqlRes, err := db.Exec(sqlString)
				if err != nil {
					log.Error("Exec SQL No Tx failed", zap.Error(err))
					writeLog(pw, fmt.Sprintf("Exec SQL No Tx failed: %s", zap.Error(err)))
					return
				}
				rows, _ := sqlRes.RowsAffected()
				writeLog(pw, fmt.Sprintf("Run Sql %d No Tx Success,Number of rows affected :%d", i, rows))
			}
		}
		exitCode = 0
	}()
	return pr
}

func writeLog(pw *io.PipeWriter, log string) {
	now := time.Now().Local().Format("2006-01-02 15:04:05: ")
	pw.Write([]byte(fmt.Sprintf("\n%s %s", now, log)))
}
