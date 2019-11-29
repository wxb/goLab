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

// 临时对象池的Put方法总会先试图把新的临时对象，存储到对应的本地池的private字段中，以便在后面获取临时对象的时候，可以快速地拿到一个可用的值。
// 只有当这个private字段已经存有某个值时，该方法才会去访问本地池的shared字段。
// 相应的，临时对象池的Get方法，总会先试图从对应的本地池的private字段处获取一个临时对象。只有当这个private字段的值为nil时，它才会去访问本地池的shared字段。
// 一个本地池的shared字段原则上可以被任何 goroutine 中的代码访问到，不论这个 goroutine 关联的是哪一个 P。这也是我把它叫做共享临时对象列表的原因。
// 相比之下，一个本地池的private字段，只可能被与之对应的那个 P 所关联的 goroutine 中的代码访问到，所以可以说，它是 P 级私有的。
// 以临时对象池的Put方法为例，它一旦发现对应的本地池的private字段已存有值，就会去访问这个本地池的shared字段。当然，由于shared字段是共享的，所以此时必须受到互斥锁的保护。
func TestGormDB(t *testing.T) {
	var dbPool = sync.Pool{
		New: func() interface{} {
			db, err := gorm.Open("mysql", "username:password@tcp(ip:port)/db?charset=utf8&parseTime=True&loc=Local")
			if err != nil {
				panic(err)
			}
			fmt.Println("execute New()")
			return db
		},
	}

	db := dbPool.Get().(*gorm.DB)
	// 临时对象池的Put方法总会先试图把新的临时对象，存储到对应的本地池的private字段中，以便在后面获取临时对象的时候，可以快速地拿到一个可用的值。
	dbPool.Put(db)
	// 只有当这个private字段已经存有某个值时，该方法才会去访问本地池的shared字段
	dbPool.Put(db)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()

		// 临时对象池的Get方法，总会先试图从对应的本地池的private字段处获取一个临时对象。
		// 只有当这个private字段的值为nil时，它才会去访问本地池的shared字段
		db := dbPool.Get().(*gorm.DB)
		pc := []string{}
		err := db.Table("t_xx").Pluck("sn", &pc).Error
		dbPool.Put(db)
		fmt.Println("t_xx", err, pc)
	}()

	sn := []string{}
	err := db.Table("t_xx").Pluck("sn", &sn).Error
	dbPool.Put(db)
	fmt.Println("t_xx", err, sn)

	<-ctx.Done()
	fmt.Println("finish")
}
