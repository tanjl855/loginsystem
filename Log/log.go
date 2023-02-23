package Log

import (
	"io"
	"log"
	"os"
)

var ErrorLog *log.Logger //Err
var Info *log.Logger     //info

func InitErrLog() *log.Logger {
	f, err := os.Create("./Log/err.log")
	if err != nil {
		panic(err)
	}
	out := io.MultiWriter(f, os.Stdout)
	logger := log.New(out, "loginsystem:", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

func InitLogLog() *log.Logger {
	f, err := os.Create("./Log/log.log")
	if err != nil {
		panic(err)
	}
	out := io.MultiWriter(f, os.Stdout)
	logger := log.New(out, "loginsystem:", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
