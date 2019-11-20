package global

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/starichat/jianshu/util"
)

var App = &app{}

type app struct {
	Name    string
	Version string
	Date    time.Time

	// 项目根目录
	RootDir string
	// 模板根目录
	TemplateDir string

	// 启动时间
	LaunchTime time.Time
	Uptime     time.Duration

	Domain string
	SEO    map[string]string

	Build struct {
		GitCommitLog string
		BuildTime    string
		GitRelease   string
		GoVersion    string
		EchoVersion  string
	}

	locker sync.Mutex
}

func (a *app) SetUptime() {
	a.locker.Lock()
	defer a.locker.Unlock()
	a.Uptime = time.Now().Sub(a.LaunchTime)
}

func (a *app) FillBuildInfo(gitCommitLog, buildTime, gitRelease string) {
	a.Build.GitCommitLog = gitCommitLog
	a.Build.BuildTime = buildTime

	pos := strings.Index(gitRelease, "/")
	if pos >= -1 {
		a.Build.GitRelease = gitRelease[pos+1:]
	}

	fmt.Println(a)
}

func (a *app) fillOtherField() {
	a.Name = viper.GetString("name")
	a.Domain = viper.GetString("domain")
	a.SEO = viper.GetStringMapString("seo")
}

func (a *app) String() string {
	return "Build Info:" +
		"\nGit Commit Log: " + a.Build.GitCommitLog +
		"\nGit Release Info: " + a.Build.GitRelease +
		"\nBuild Time: " + a.Build.BuildTime +
		"\nGo Version: " + a.Build.GoVersion
}

// inferRootDir 递归推导项目根目录
func inferRootDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 这里的技巧可以学习下：先定义一个函数变量，然后在赋值函数时，里面使用了这个函数变量
	var infer func(d string) string
	infer = func(d string) string {
		if d == "/" {
			panic("请确保在项目根目录或子目录下运行程序，当前在：" + cwd)
		}

		if util.Exist(d + "/config") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	return infer(cwd)
}

func init() {
	App.Version = "V1.0"
	App.LaunchTime = time.Now()

	// 默认在项目根目录运行程序
	App.RootDir = "."

	// 用来处理单元测试或当前目录不在根目录，获取项目根目录
	if !viper.InConfig("http.port") {
		App.RootDir = inferRootDir()
	}
	App.TemplateDir = App.RootDir + "/template/"

	fileInfo, err := os.Stat(os.Args[0])
	if err != nil {
		panic(err)
	}

	App.Date = fileInfo.ModTime()
	App.Build.GoVersion = runtime.Version()
	App.Build.EchoVersion = echo.Version
}
