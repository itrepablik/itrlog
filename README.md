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
)

func init() {
	// Custom settings to initialize the itrlog.
	itrlog.SetLogInit(50, 90, "logs", "test_log_")
}

func main() {
	// Basic logging message
	itrlog.Info("basic info log output...")
	itrlog.Error("basic error log output...")
	itrlog.Warn("basic warning log output...")
	itrlog.Debug("basic debug log output...")
	
	// Structured context as loosely typed key-value pairs.
	itrlog.Infow("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	itrlog.Errorw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	itrlog.Warnw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	itrlog.Debugw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
}
```

## To use the **itrlog** from the other packages, for example:
```
package app

import (
	"github.com/itrepablik/itrlog"
)

func CreateUser() {
	itrlog.Info("User has been created successfully!")
}
```

It will auto-create the **logs** folder in the root directory of your project, but, of course, it's up to you which directory to point your log files and the **test_log_** is the initial log filename.  The logging output it's a **JSON** format for easy integration with any of your existing software to keep track of the logs information.

# License
Code is distributed under MIT license, feel free to use it in your proprietary projects as well.
