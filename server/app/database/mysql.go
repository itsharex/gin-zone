package database

import (
	"fmt"
	"log"
	"time"

	mylog "gitee.com/jiang-xia/gin-zone/server/pkg/log"
	"gitee.com/jiang-xia/gin-zone/server/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Mysql *gorm.DB                 // gorm数据库实例
var file = utils.GetLogFile("sql") // 获取记录gorm错误的日志文件

var newLogger = logger.New(
	log.New(file, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	logger.Config{
		SlowThreshold:             time.Second, // 慢 SQL 阈值
		LogLevel:                  logger.Info, // 日志级别
		IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,       // 彩色打印
	},
)

/* 使用 logrus 记录gorm日志 */
//定义自己的Writer
type MyWriter struct {
	mlog *logrus.Logger
}

// 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.mlog.Info(logstr)
}

// 自定义记录器
func NewMyWriter() *MyWriter {
	log := mylog.Logger
	return &MyWriter{mlog: log}
}

/*
*
数据库初始化
*/
func DatabaseSetup() {
	// 根据自定义Writer，新建一个logger
	slowLogger := logger.New(
		//设置Logger
		NewMyWriter(),
		logger.Config{
			//慢SQL阈值
			SlowThreshold: time.Second,
			//设置日志级别，只有Warn以上才会打印sql
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	// dsn := "root:88888888@/elk-blog?charset=utf8&parseTime=True&loc=Local"

	var (
		err                                                     error
		dbType, dbName, user, password, host, port, tablePrefix string
	)
	// 读取配置
	sec := utils.Config.Section("database")

	dbType = sec.Key("type").String()
	dbName = sec.Key("dbname").String()
	user = sec.Key("user").String()
	password = sec.Key("password").String()
	host = sec.Key("host").String()
	port = sec.Key("port").String()
	tablePrefix = sec.Key("table_prefix").String()
	fmt.Println(dbType, dbName, user, password, host, tablePrefix)
	// dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 自定义表名前缀
			SingularTable: true,        // 全局禁用表名复数

		},
		Logger: slowLogger,
		// Logger: logger.Default.LogMode(logger.Info),
	})
	// 需要把当前成功连接的实例赋值给全局变量Mysql(不然没法操作数据库)
	Mysql = conn
	sqlDB, err := Mysql.DB()
	sqlDB.SetMaxIdleConns(25)                 // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100)                // 设置最大连接数
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // 设置每个链接的过期时间

	fmt.Println("数据库连接成功")
	logrus.Info("数据库连接成功")
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	// 调试单个操作，显示此操作的详细日志
	// db.Debug().Where("name = ?", "jinzhu").First()
}

func CloseDB() {
	defer func() {
		sqlDB, err := Mysql.DB()
		sqlDB.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}()
}
