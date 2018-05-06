package logger

import (
	"os"
	"time"

	"fmt"

	"bitbucket.org/insamo/mvc/web/bootstrap"
)

func ConfigureApplicationLogger(b *bootstrap.Bootstrapper) {
	filename := ""
	if b.Environment.Core().GetString("log.level") == "debug" {
		filename = b.Environment.Core().GetString("log.storage") + "/debug.app.log"
	} else {
		filename = b.Environment.Core().GetString("log.storage") + "/" + time.Now().Format("2006-01-02") + ".app.log"
	}
	// open an output file, this will append to the today's file if server restarted.
	var err error
	b.ApplicationLogFile, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		if err != nil {
			fmt.Errorf("Failed open application log file: %s \n", err)
		}
	}
	b.Logger().SetOutput(b.ApplicationLogFile)
	b.Logger().SetLevel(b.Environment.Core().GetString("log.level"))
}
