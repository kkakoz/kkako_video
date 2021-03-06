package mysqlx

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

type Options struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func New(viper *viper.Viper) (*gorm.DB, error) {
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.name", "test")
	o := &Options{}
	var err error
	if err := viper.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "viper unmarshal失败")
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?"+
		"charset=utf8mb4&parseTime=True&loc=Local",
		o.User, o.Password,
		o.Host, o.Port,
		o.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	config := &gorm.Config{
		Logger: newLogger,
	}
	db, err = gorm.Open(mysql.Open(dns), config)
	return db, errors.Wrap(err, "打开mysql连接失败")
}

// 可以看 https://github.com/win5do/go-microservice-demo/blob/main/docs/sections/gorm.md
type ctxTransactionKey struct{}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxTransactionKey{}, tx)
}

func Transaction(ctx context.Context, fc func(txctx context.Context) error) error {
	db := db.WithContext(ctx)

	return db.Transaction(func(tx *gorm.DB) error {
		txctx := CtxWithTransaction(ctx, tx)
		return fc(txctx)
	})
}

func Begin(ctx context.Context, opts ...*sql.TxOptions) (context.Context, CheckError) {
	tx := db.Begin(opts...)
	return context.WithValue(ctx, ctxTransactionKey{}, tx), func(err error) error {
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Commit().Error
		return err
	}
}

func GetDB(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if ok {
			return tx
		}
	}

	return db.WithContext(ctx)
}

func CheckErr(database *gorm.DB, err error) error {
	if err != nil {
		curErr := database.Rollback().Error
		if curErr != nil {
			return curErr
		}
		return err
	}
	return database.Commit().Error
}

type Model struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FlushDB() {

	var tables []string
	err := db.Table("information_schema.tables").Where("table_schema = ?", "kkako_video_test").Pluck("table_name", &tables).Error
	if err != nil {
		log.Fatalln(err)
	}

	for _, table := range tables {
		db.Table(table).Exec(fmt.Sprintf("truncate table `%s`", table))
	}

}

type CheckError func(err error) error