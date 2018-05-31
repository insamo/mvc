package logger

import (
	"os"
	"time"

	"fmt"

	"github.com/insamo/mvc/web/bootstrap"
)

func ConfigureDatabaseLogger(b *bootstrap.Bootstrapper) {
	filename := ""
	if b.Environment.Core().GetString("log.level") == "debug" {
		filename = b.Environment.Core().GetString("log.storage") + "/debug.database.log"
	} else {
		filename = b.Environment.Core().GetString("log.storage") + "/" + time.Now().Format("2006-01-02") + ".database.log"
	}
	// open an output file, this will append to the today's file if server restarted.
	var err error
	b.DatabaseLogFile, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		if err != nil {
			fmt.Errorf("Failed open application log file: %s \n", err)
		}
	}
}
