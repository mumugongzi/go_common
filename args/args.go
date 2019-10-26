package args

import (
	"errors"
	"flag"
	"os"
)

var Args = &args{}

// example
type args struct {
	rawStartTimeStr *string
	rawEndTimeStr   *string
	rawMsIDStr      *string
	help            *bool

	UTCStartTimeStr string
	UTCEndTimeStr   string
	MSNodeIDList    []int64
}

func (a *args) Init() {
	a.rawStartTimeStr = flag.String("s", "", "指定开始时间(北京时间): 2018-11-11 11:11:11")
	a.rawEndTimeStr = flag.String("e", "", "指定结束时间(北京时间): 2018-11-11 11:11:11")
	a.rawMsIDStr = flag.String("m", "", "需要统计的MS节点")
	a.help = flag.Bool("h", false, "显示ms节点ID与名称对应关系")
	flag.Parse()
}

func (a *args) validate() error {
	if *a.rawStartTimeStr == "" {
		return errors.New("start time is empty")
	}

	if *a.rawEndTimeStr == "" {
		return errors.New("end time is empty")
	}

	if *a.rawEndTimeStr == "" {
		return errors.New("ms id is empty")
	}
	return nil
}

func (a *args) Parse() error {
	a.Init()

	if *a.help {
		a.usage()
		os.Exit(0)
	}
	return nil
}

func (a *args) usage() {
	flag.Usage()
}
