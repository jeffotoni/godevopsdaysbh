package logrus

import (
	"bytes"
	"fmt"
	logrusof "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// capture stdout
var old, os_r, os_w *os.File

// new logrus
var Log = logrusof.New()

// fields Type map[string]interfaces{}
type Filelds logrusof.Fields

type LogrusEntry struct {
	F Filelds
}

// struct splitter
type OutputSplitter struct{}

func init() {

	Log.Out = os.Stdout
	Log.SetFormatter(&logrusof.JSONFormatter{})
	Log.SetReportCaller(true)
	//Log.SetOutput(os.Stdout)

	// to capture stdout custom
	Log.SetOutput(&OutputSplitter{}) // this is capture stdout
	// Only log the warning severity or above.
	// Log.SetLevel(logrusof.WarnLevel)
}

func Fileld(f Filelds) map[string]interface{} {
	return f
}

func (splitter *OutputSplitter) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("level=error")) {
		return os.Stderr.Write(p)
	}
	return os.Stdout.Write(p)
}

//var Info = &LogrusEntry{}
func (info *LogrusEntry) Info(msg string) {
	// start
	//StdoutStart()
	// logrus func
	Log.WithFields(Fileld(info.F)).Info(msg)
	// end
	//StdoutEnd()
}

//var Info = &LogrusEntry{}
func (e *LogrusEntry) Error(msg string) {
	//StdoutStart()
	Log.WithFields(Fileld(e.F)).Error(msg)
	//StdoutEnd()
}

func WithFields(fields Filelds) *LogrusEntry {
	return &LogrusEntry{
		F: fields,
	}
}

func StdoutStart() {
	old = os.Stdout
	os_r, os_w, _ = os.Pipe()
	os.Stdout = os_w
}

func StdoutEnd() {
	os_w.Close()
	out, _ := ioutil.ReadAll(os_r)
	os.Stdout = old

	// asynchronous, publish receive string
	// reativar here....
	// go gologs.Publish(string(out)).Rpc()
	////////////////////////////////////////

	fmt.Printf("\n%s", out)
	// go gologs.Publish(string(out)).Tcp()
}
