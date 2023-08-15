package ylog

import (
	"fmt"
	"os"
	"path"
)

func init() {
	logPath := "." + "/" + "test" + ".log"
	if err := InitAppLog(
		ProcessName(path.Base(os.Args[0])+"_rd"),
		LogPath(logPath),
		TestEnv(false),
	); err != nil {
		println("InitAppLog : ", err)
	}

	fmt.Println("Initialized ulog on path " + logPath)

	// xconfig.OnConfigChange(InitLogLevel)
	// tars.RegisterAdmin("tars.setloglevel", notifySetLogLevel)
}

func InitLogLevel() {
	err := SetLogLevel("debug")
	if err != nil {
		fmt.Println(err)
	}
}

// func notifySetLogLevel(command string) (string, error) {
//	cmd := strings.Split(command, " ")
//	err := SetLogLevel(cmd[1])
//	if err != nil {
//		println("notifySetLogLevel - SetLogLevel: ", err)
//		return "", nil
//	}
//	ret := fmt.Sprintf("notifySetLogLevel - SetLogLevel to %s success", cmd[1])
//	fmt.Println(ret)
//	return ret, nil
// }
