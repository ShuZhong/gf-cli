package get

import (
	"fmt"
	"os"

	"github.com/ShuZhong/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
)

func Help() {
	mlog.Print(gstr.TrimLeft(`
USAGE    
    gf get PACKAGE

ARGUMENT 
    PACKAGE  remote golang package path, eg: github.com/ShuZhong/gf

EXAMPLES
    gf get github.com/ShuZhong/gf
    gf get github.com/ShuZhong/gf@latest
    gf get github.com/ShuZhong/gf@master
    gf get golang.org/x/sys
`))
}

func Run() {
	if len(os.Args) > 2 {
		gproc.ShellRun(fmt.Sprintf(`go get -u %s`, gstr.Join(os.Args[2:], " ")))
	} else {
		mlog.Fatal("please input the package path for get")
	}
}
