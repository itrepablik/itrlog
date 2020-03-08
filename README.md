![ITRLog](https://user-images.githubusercontent.com/58651329/75599806-682bd300-5ae3-11ea-9d9b-adb175589143.png)

The simplified usage of [Zap](https://github.com/uber-go/zap) and [Lumberjack](https://github.com/natefinch/lumberjack) logging systems in Go for easier usage of the library.

# Installation
```
go get -u github.com/itrepablik/itrlog
```

# Usage
This is how you can use the simplified initializations for Zap and Lumberjack logging libraries.
```
package main

import (
	"github.com/itrepablik/itrlog"
	"go.uber.org/zap"
)

// Sugar zap and lumberjack logger initialization simplified by ITRepablik for easy usage
// and its declarations to be global throughout your Go' project.
var Sugar *zap.SugaredLogger

func init() {
	Sugar = itrlog.InitLog(50, 28, "logs", "test_log_")
	Sugar.Info("hello test sugar info log...")
}

func main() {
	// Basic logging message
	Sugar.Info("basic info log output...")
	Sugar.Error("basic error log output...")
	Sugar.Warn("basic warning log output...")
	Sugar.Debug("basic debug log output...")
	
	// Structured context as loosely typed key-value pairs.
	Sugar.Infow("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	Sugar.Errorw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	Sugar.Warnw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	Sugar.Debugw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
}
```

It will auto-create the **logs** folder in the root directory of your project, but, of course, it's up to you which directory to point your log files and the **test_log_** is the initial log filename.  The logging output it's a **JSON** format for easy integration with any of your existing software to keep track of the logs information.

# License
Code is distributed under MIT license, feel free to use it in your proprietary projects as well.
