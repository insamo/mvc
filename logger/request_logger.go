package logger

import (
	"os"
	"strings"
	"time"

	"fmt"

	"github.com/insamo/mvc/web/bootstrap"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
)

const deleteFileOnExit = false

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

//func ConfigureRequestLogger(b *bootstrap.Bootstrapper) (h context.Handler, close func() error) {
func ConfigureRequestLogger(b *bootstrap.Bootstrapper) {
	if b.Environment.Core().GetString("request.log.level") == "disable" {
		return
	}

	var err error
	//close = func() error { return nil }

	c := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}

	filename := ""

	if b.Environment.Core().GetString("log.level") == "debug" {
		filename = b.Environment.Core().GetString("request.log.storage") + "/debug.request.log"
	} else {
		filename = b.Environment.Core().GetString("request.log.storage") + "/" + time.Now().Format("2006-01-02") + ".request.log"
	}
	// open an output file, this will append to the today's file if server restarted.
	b.RequestLogFile, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		if err != nil {
			if err != nil {
				fmt.Errorf("Failed open request log file: %s \n", err)
			}
		}
	}

	//close = func() error {
	//	err := logFile.Close()
	//	if deleteFileOnExit {
	//		err = os.Remove(logFile.Name())
	//	}
	//	return err
	//}

	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		line := fmt.Sprintf("%s | %v | %4v | %s | %s | %s", now.String(), status, latency, ip, method, path)
		if message != nil {
			line += fmt.Sprintf(" | %v", message)
		}

		if headerMessage != nil {
			line += fmt.Sprintf(" | %v", headerMessage)
		}

		b.RequestLogFile.Write([]byte(line + "\n"))
	}

	//	we don't want to use the logger
	// to log requests to assets and etc
	c.AddSkipper(func(ctx context.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})

	h := logger.New(c)
	b.UseGlobal(h)
	//return
}
