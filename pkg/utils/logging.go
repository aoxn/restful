package utils

import (
    "github.com/hhkbp2/go-logging"
    "fmt"
    //"yy.com/container/apiserver"
    //"github.com/GoogleCloudPlatform/kubernetes/pkg/auth/handlers"
    "strconv"
    "os"
    "strings"
)

const (
    LOGGING_PROPERTY_FILE = "logging.file"
    LOGGING_PROPERTY_ROTATE_SIZE = "logging.rotate_size"
    LOGGING_PROPERTY_LEVEL = "logging_level"
)


type Logger struct {
    logger logging.Logger
}

// Fatal formats using the default formats for its operands and
// logs a message with severity "LevelFatal".
func (l *Logger) Fatal(args ...interface{}) {
    l.logger.Fatal(args...)
}
// Error formats using the default formats for its operands and
// logs a message with severity "LevelError".
func (l *Logger) Error(args ...interface{}) {
    l.logger.Error(args...)
}
// Warn formats using the default formats for its operands and
// logs a message with severity "LevelWarn".
func (l *Logger) Warn(args ...interface{}) {
    l.logger.Warn(args...)
}
// Info formats using the default formats for its operands and
// logs a message with severity "LevelInfo".
func (l *Logger) Info(args ...interface{}) {
    l.logger.Info(args...)
}
// Debug formats using the default formats for its operands and
// logs a message with severity "LevelDebug".
func (l *Logger) Debug(args ...interface{}) {
    l.logger.Debug(args...)
}
// Fatalf formats according to a format specifier and
// logs a message with severity "LevelFatal".
func (l *Logger) Fatalf(format string, args ...interface{}) {
    l.logger.Fatalf(format, args...)
}
// Errorf formats according to a format specifier and
// logs a message with severity "LevelError".
func (l *Logger) Errorf(format string, args ...interface{}) {
    l.logger.Errorf(format, args...)
}
// Warnf formats according to a format specifier and
// logs a message with severity "LevelWarn".
func (l *Logger) Warnf(format string, args ...interface{}) {
    l.logger.Warnf(format, args...)
}
// Infof formats according to a format specifier and
// logs a message with severity "LevelInfo".
func (l *Logger) Infof(format string, args ...interface{}) {
    l.logger.Infof(format, args...)
}
// Debugf formats according to a format specifier and
// logs a message with severity "LevelDebug".
func (l *Logger) Debugf(format string, args ...interface{}) {
    l.logger.Debugf(format, args...)
}

const (
    LOG_MODULE_APISVR = "apisvr"
    LOG_MODULE_SCHED = "sched"
    LOG_MODULE_ARCH = "arch"
    LOG_MODULE_MASTERCORE = "mastercore"
    LOG_MODULE_SLAVECORE = "slavecore"
    LOG_MODULE_SLAVERT = "slavert"
)





func GetLogger(module string) *Logger {
    l := &Logger{}
    l.logger = logging.GetLogger(module)
    return l

    /*
    fmt.Printf("[GetLogger] start\r\n")

    handler := logging.NewStdoutHandler()
    l.logger.AddHandler(handler)
    l.logger.SetLevel(logging.LevelDebug)

    fmt.Println("GetLevel()", l.logger.GetLevel())
    fmt.Println("handler.GetLevel()", handler.GetLevel())

    l.logger.Infof("[GetLogger] mode:%s,message: %s %d", module, "Hello", 2015)
    fmt.Printf("[GetLogger] ok\r\n")
    */


}



func setLogger(strName string,strFile string,strSize string){

    var uData uint64
    var error error
    //func ParseUint(s string, base int, bitSize int) (n uint64, err error)
    uData,error=strconv.ParseUint(strSize, 10, 64)
    if error != nil{
        fmt.Println("[setLogger] strconv.ParseUint  failed,size:%s",strSize)
    }

    filePath :=strFile
    fileMode := os.O_APPEND
    // set the maximum size of every file to 100 M bytes
    fileMaxBytes := uData
    // keep 9 backup at most(including the current using one,
    // there could be 10 log file at most)
    backupCount := uint32(10)
    // create a handler(which represents a log message destination)
    handler := logging.MustNewRotatingFileHandler(
    filePath, fileMode, fileMaxBytes, backupCount)

    // the format for the whole log message
    format := "%(asctime)s %(levelname)s (%(filename)s:%(lineno)d) " +
    "%(name)s %(message)s"

    // the format for the time part
    dateFormat := "%Y-%m-%d %H:%M:%S.%3n"
    // create a formatter(which controls how log messages are formatted)
    formatter := logging.NewStandardFormatter(format, dateFormat)
    // set formatter for handler
    handler.SetFormatter(formatter)

    // create a logger(which represents a log message source)
    logger := logging.GetLogger(strName)
    logger.SetLevel(logging.LevelDebug)
    logger.AddHandler(handler)


    fmt.Printf("[setLogger] name:%s,file:%s,size:%s ok\r\n", strName,strFile,strSize)


}

type logmodule struct {
    bIsFirst bool
    strFile string
    strSize string
    vecModule []string
}




func (lm *logmodule) OnValueChange(key string, oldval interface{}, newval interface{}) bool {
    var bIsOk bool
    var strOld string
    var strNew string

    if oldval!=nil {
        strOld, bIsOk=oldval.(string)
        if !bIsOk {
            fmt.Printf("[OnValueChange] key:%s,bIsOk=oldval.(string) failed \r\n", key)
            return false
        }
    }

    if newval!=nil {
        strNew, bIsOk=newval.(string)
        if !bIsOk {
            fmt.Printf("[OnValueChange] key:%s,bIsOk=newval.(string) failed \r\n", key)
            return false
        }
    }

    fmt.Printf("[OnValueChange] key:%s,oldval:%s,newval:%s start\r\n", key, strOld, strNew)



    if key==LOGGING_PROPERTY_FILE {
        lm.strFile=strNew
    }

    if key==LOGGING_PROPERTY_ROTATE_SIZE {
        lm.strSize=strNew
    }


    if lm.bIsFirst==true && lm.strFile!="" && lm.strSize!="" {//日志第一次初始化
        lm.bIsFirst=false
        for _, strName := range lm.vecModule {
            setLogger(strName, lm.strFile, lm.strSize)
            fmt.Printf("[OnValueChange] key:%s,oldval:%s,newval:%s,module:%s ok\r\n", key, strOld, strNew, strName)
        }
    }


    if lm.bIsFirst==false && strings.HasSuffix(key, LOGGING_PROPERTY_LEVEL) {//初始化后才去设置level,判断是否有这个后缀
        mapTemp := map[string]logging.LogLevelType{
            "debug": logging.LevelDebug,
            "info": logging.LevelInfo,
            "warn":logging.LevelWarn,
            "error":logging.LevelError,
            "fatal":logging.LevelFatal,
            "critical":logging.LevelCritical,
        }
        _, bIsok := mapTemp[strNew] // 如果key1存在则ok == true，否在ok为false

        if !bIsok {
            fmt.Printf("[OnValueChange] key:%s,oldval:%s,newval:%s failed\r\n", key, strOld, strNew)
            return false
        }

        var vecData []string
        vecData=strings.Split(key, ".")//格式是模块名.LOGGING_PROPERTY_LEVEL
        if len(vecData)>0 {
            bIsFound := false
            for _, strName := range lm.vecModule {
                if strName==vecData[0] {
                    objLog := logging.GetLogger(vecData[0])
                    objLog.SetLevel(mapTemp[strNew])
                    fmt.Printf("[OnValueChange] key:%s,oldval:%s,newval:%s ok\r\n", key, strOld, strNew)
                    bIsFound=true
                    break
                }

            }
            return bIsFound
        } else {
             return false
        }
    }

    return true


}


var g_logmodule logmodule

func InitLogging(cm ConfMgr) {
    g_logmodule = logmodule{
        bIsFirst:true,
        strFile:"",
        strSize:"",
        vecModule:[]string{LOG_MODULE_APISVR,
            LOG_MODULE_SCHED,LOG_MODULE_ARCH,LOG_MODULE_MASTERCORE,LOG_MODULE_SLAVECORE,LOG_MODULE_SLAVERT},
    }

    cm.RegisterWatcher(LOGGING_PROPERTY_FILE, &g_logmodule)
    cm.RegisterWatcher(LOG_MODULE_APISVR+"."+LOGGING_PROPERTY_LEVEL, &g_logmodule)
    cm.RegisterWatcher(LOGGING_PROPERTY_ROTATE_SIZE, &g_logmodule)
    cm.Refresh()
}
