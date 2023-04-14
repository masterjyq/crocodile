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
			now := time.Now().Local().Format("2006-01-02 15:04:05: ")
			pw.Write([]byte(fmt.Sprintf("\n%sRun Finished,Return Code:%5d", now, exitCode))) // write exitCode,total 5 byte
			// pw.Write([]byte(fmt.Sprintf("%3d", exitCode))) // write exitCode,total 3 byte
		}()
		// 初始化连接
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s)", da.User, da.PassWord, da.Host, da.Port, da.DataBase))
		if err != nil {
			log.Error("Open Data Base failed", zap.Error(err))
			return
		}
		defer db.Close()

		if da.Opentx {
			// start tx
			tx, err := db.Begin()
			if err != nil {
				log.Error("Open Tx failed", zap.Error(err))
				return
			}
			for i, sqlString := range da.SqlData {
				sqlRes, err := tx.Exec(sqlString)
				if err != nil {
					log.Error("Exec SQL And Tx failed ", zap.Error(err))
					return
				}
				now := time.Now().Local().Format("2006-01-02 15:04:05: ")
				pw.Write([]byte(fmt.Sprintf("\n%sRun Sql And Tx %d Success,Return :%s", now, i, sqlRes)))
			}
			err = tx.Commit()
			if err != nil {
				log.Error("Tx Commit failed", zap.Error(err))
				return
			}
		} else {
			for i, sqlString := range da.SqlData {
				sqlRes, err := db.Exec(sqlString)
				if err != nil {
					log.Error("Exec SQL No Tx failed", zap.Error(err))
					return
				}
				now := time.Now().Local().Format("2006-01-02 15:04:05: ")
				pw.Write([]byte(fmt.Sprintf("\n%sRun Sql No Tx %d Success,Return :%s", now, i, sqlRes)))
			}
		}
		exitCode = 0
	}()
	return pr
}
