package pool_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer([]byte(""))
	},
}

func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)

	bb := bufPool.Get().(*bytes.Buffer)
	w.Write(bb.Bytes())
}

func TestLog(t *testing.T) {
	log(os.Stdout, "path", "/search?q=flowers")
}

func TestGormDB(t *testing.T) {
	var dbPool = sync.Pool{
		New: func() interface{} {
			db, err := gorm.Open("mysql", "username:password@tcp(ip:port)/db?charset=utf8&parseTime=True&loc=Local")
			if err != nil {
				panic(err)
			}
			return db
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()

		db := dbPool.Get().(*gorm.DB)
		pc := []string{}
		err := db.Table("t_xxx").Pluck("sn", &pc).Error
		dbPool.Put(db)
		fmt.Println("t_xxx", err, pc)
	}()

	db := dbPool.Get().(*gorm.DB)
	sn := []string{}
	err := db.Table("t_xxx").Pluck("sn", &sn).Error
	dbPool.Put(db)
	fmt.Println("t_xxx", err, sn)

	<-ctx.Done()
	fmt.Println("finish")
}
