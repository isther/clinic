package conf

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	filePerm = 0644
	dirPerm  = 0755
)

var Server *server

type server struct {
	Dev      bool   `yaml:"dev"`
	LogLevel string `yaml:"log_level"`
	BaseURL  string `yaml:"base_url"`
	Listen   string `yaml:"listen"`
	DB       struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		SSLMode  string `yaml:"ssl_mode"`
		TimeZone string `yaml:"time_zone"`
	} `yaml:"db"`
	ImgDir string `yaml:"img_dir"`
}

func (s *server) DSN() string {
	return "host=" + s.DB.Host + " port=" + strconv.Itoa(s.DB.Port) + " user=" + s.DB.User + " password=" + s.DB.Password +
		" dbname=" + s.DB.DBName + " sslmode=" + s.DB.SSLMode + " TimeZone=" + s.DB.TimeZone
}

func init() {
	var err error

	Server = new(server)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	err = setupLogDir()
	if err != nil {
		logrus.Fatalln(err)
	}

	err = setupLogOutput()
	if err != nil {
		logrus.Fatalln(err)
	}

	yamlFileBytes, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		logrus.Fatalln(err)
	}

	err = yaml.Unmarshal(yamlFileBytes, Server)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = setupLogLevel(Server.LogLevel)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = setupGinLog()
	if err != nil {
		logrus.Fatalln(err)
	}

	if !Server.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
}

func setupLogDir() error {
	var err error
	if _, err = os.Stat("./logs/"); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir("./logs/", dirPerm)
	}
	return err
}

func setupLogOutput() error {
	var err error

	logFileName := time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile("./logs/"+logFileName+"-app-all.log", syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, filePerm)
	if err != nil {
		return err
	}
	logOut := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(logOut)

	return err
}

func setupLogLevel(pLevel string) error {
	switch pLevel {
	case "":
		logrus.SetLevel(logrus.DebugLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		return errors.New("unknown log level: " + pLevel)
	}
	return nil
}

func setupGinLog() error {
	var err error
	logErrorFileName := time.Now().Format("2006-01-02")
	logErrorFile, err := os.OpenFile("./logs/"+logErrorFileName+"-gin-error.log", syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, filePerm)
	if err != nil {
		return err
	}
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, logErrorFile)
	logInfoFileName := time.Now().Format("2006-01-02")
	logInfoFile, err := os.OpenFile("./logs/"+logInfoFileName+"-gin-info.log", syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, filePerm)
	if err != nil {
		return err
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logInfoFile)
	return err
}
