# itrlog

# Usage
This is how you can use the simplified initializations for Zap and Lamberjack logging libraries.
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
	Sugar = itrlog.InitLog(50, 28, "logs", "_test_log_")
	Sugar.Info("hello test sugar info log...")
}

func main() {
	Sugar.Error("error type of log...")
}
```
