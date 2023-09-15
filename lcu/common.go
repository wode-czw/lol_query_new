package lcu

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
	"github.com/real-web-world/hh-lol-prophet/pkg/windows/process"
)

const (
	lolUxProcessName = "LeagueClientUx.exe"
)

//common这个函数是用来找进程的
//但是他写的v3这个版本很奇怪，到了某一步我查不下去了

var (
	lolCommandlineReg     = regexp.MustCompile(`--remoting-auth-token=(.+?)" "--app-port=(\d+)"`)
	ErrLolProcessNotFound = errors.New("未找到lol进程")
)

func GetLolClientApiInfo() (int, string, error) {
	return GetLolClientApiInfoV3()
}

// 从pid的列表里面获取LeagueClientUx.exe的信息，根据其中的命令参数，返还令牌和port
func GetLolClientApiInfoV3() (port int, token string, err error) {
	cmdline, err := process.GetProcessCommand(lolUxProcessName) //这个cmdline是个string类型，得到的除了命令名字，还有参数命令

	/*
		regex := regexp.MustCompile(`"([^"]+)"`)
		matches := regex.FindAllString(cmdline, -1)

		for _, arg := range matches {
			fmt.Println(arg)
		}
		fmt.Println("===============================================================================================")
		fmt.Println("===============================================================================================")

	*/

	if err != nil {
		err = ErrLolProcessNotFound
		return
	}

	btsChunk := lolCommandlineReg.FindSubmatch([]byte(cmdline)) //用正则表达式从这一段参数里面找到port和token,放在了一个[][]uint8的数组
	if len(btsChunk) < 3 {
		return port, token, ErrLolProcessNotFound
	}

	token = string(btsChunk[1])
	port, err = strconv.Atoi(string(btsChunk[2]))
	return
}
