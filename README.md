![ITRLog](https://user-images.githubusercontent.com/58651329/80480060-96624d80-8982-11ea-994f-153f4f987fbe.png)

The simplified usage of [Zap](https://github.com/uber-go/zap) and [Lumberjack](https://github.com/natefinch/lumberjack) logging systems in Go for easier usage of the library.

# Installation
```go
go get -u github.com/itrepablik/itrlog
```

# Usage
This is how you can use the simplified initializations for Zap and Lumberjack logging libraries.
```go
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
	
	// Similar to the standard Go's fmt or log formatting
	itrlog.Infof("started API on port: %s", port)
	itrlog.Errorf("server started at: %s", time.Now().Format(time.RFC3339))
	itrlog.Warnf("fetched data results: %v", results)
	itrlog.Debugf("debugging starts here: %v", err)
	
	// Structured context as loosely typed key-value pairs.
	itrlog.Infow("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	itrlog.Errorw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	itrlog.Warnw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
	itrlog.Debugw("New User", "username", "admin", "email", "email@site.com", "log_time", time.Now().Format(itrlog.LogTimeFormat))
}
```
It will auto-create the **logs** folder in the root directory of your project, but of course, it's up to you which directory to point your log files to. The **test_log_** is the initial log filename. The logging output is in a **JSON** format for easy integration with any of your existing software to keep track of the log information.

## Used By

The following projects use this project:

https://github.com/itrepablik/itrlog/network/dependents

# Subscribe to Maharlikans Code Youtube Channel:
Please consider subscribing to my YouTube channel to recognize my work on any of my tutorial series. Thank you so much for your support!
https://www.youtube.com/c/MaharlikansCode?sub_confirmation=1

# License
I've shared the code under the MIT license, feel free to use it in your proprietary projects as well.
