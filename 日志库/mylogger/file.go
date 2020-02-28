package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 往文件中打印日志
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件名称
	fileMaxSize int64  // 文件最大大小
	fileObj     *os.File
	errFileObj  *os.File
	// errorFileName string // 错误日志文件名称
}

// NewFileLog 构造结构体函数
func NewFileLog(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		fileMaxSize: maxSize,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}
func (f *FileLogger) initFile() error {
	fileObj, err := os.OpenFile(f.filePath+f.fileName+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Open LogFile failed, err:", err)
		return err
	}
	errFileObj, err := os.OpenFile(f.filePath+f.fileName+".err.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Open LogFile failed, err:", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel > f.Level
}
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Get file info failed")
	}
	// 如果传的文件大小>最大文件大小，返回true
	return fileInfo.Size() > f.fileMaxSize
}
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		now := time.Now()
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFileObj, err := f.splitFile(f.fileObj)
			if err != nil {
				fmt.Println("Split file failed")
				return
			}
			f.fileObj = newFileObj
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFileObj, err := f.splitFile(f.errFileObj)
				if err != nil {
					fmt.Println("Split file failed")
					return
				}
				f.errFileObj = newFileObj
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}
func (f *FileLogger) splitFile(fileObj *os.File) (*os.File, error) {
	// 需要切割日志文件
	// 1.关闭当前日志文件
	fileObj.Close()
	// 2.备份一下
	nowFileName := path.Join(f.filePath + f.fileName + ".log")
	nowStr := time.Now().Format("20060102150405000")
	newFileName := fmt.Sprintf("%s%s.%s.log", f.filePath, f.fileName, nowStr)
	os.Rename(nowFileName, newFileName)
	// 3.打开一个新的日志文件
	newFileObj, err := os.OpenFile(nowFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Open LogFile failed, err:", err)
		return nil, err
	}
	// 4.将打开的日志文件赋值给f.fileObj
	return newFileObj, nil
}

// Debug debug级别的日志记录
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Trace trace级别的日志记录
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

// Info info级别的日志记录
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning warning级别的日志记录
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error error级别的日志记录
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal fatal级别的日志记录
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// Close 关闭记录日志的文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
