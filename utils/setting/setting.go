package setting

import (
	"embed"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	QrCodePath string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	ServerIp     string
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
	Prefix   string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host     string
	Password string
	Db       int
	Prefix   string
}

var AppPath string

var RedisSetting = &Redis{}

var f embed.FS

func Setup() {

	var iniPath string
	env := GetEnv()
	if env == "" {
		iniPath = "./conf/app.ini"
	} else {
		iniPath = fmt.Sprintf("./conf/app.%s.ini", env)
	}
	iniData, _ := f.ReadFile(iniPath)
	Cfg, err := ini.Load(iniData)
	if err != nil {
		log.Fatal(2, "load ini error:%v", err)
	}
	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err:%v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)

	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err:%v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)

	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err:%v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)

	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err:%v", err)
	}
}

func GetEnv() string {
	env := os.Getenv("APPLICATION_PATH")
	if env == "" || env == "production" {
		return ""
	}
	return env
}
