package main

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	fmt.Println("你好")

	_, err := process.NewProcess(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	mbInt := uint64(1024) * 1024
	gbInt := uint64(1024) * 1024 * 1024

	machineMemory, err := mem.VirtualMemoryWithContext(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(
		fmt.Sprintf(
			"当前电脑内存信息：总内存【%dGB】，可用内存「%dGB」，已使内存「%dGB」", machineMemory.Total/gbInt,
			machineMemory.Available/gbInt, machineMemory.Used/gbInt,
		),
	)

	withContext, err := process.ProcessesWithContext(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, pv := range withContext {
		n, err := pv.Name()
		fmt.Println(n)
		if err == nil && n == "etcd" {
			fmt.Println(fmt.Sprintf("LosslessSwitcher Pid:%d", pv.Pid))
			username, _ := pv.Username()
			fmt.Println(username)
			cwd, _ := pv.Cwd()
			fmt.Println(cwd)
			info, _ := pv.MemoryInfo()
			fmt.Println(fmt.Sprintf("内存信息：物理内存【%dMB】，虚拟内存「%dGB」", info.RSS/mbInt, info.VMS/gbInt))
			name, _ := pv.Name()
			fmt.Println(name)
			threads, _ := pv.NumThreads()
			fmt.Println(threads)

			cPUPercent, _ := pv.CPUPercent()
			fmt.Println(cPUPercent)

			cPUTimes, _ := pv.Times()
			fmt.Println(cPUTimes)
			connections, _ := pv.Connections()
			fmt.Println(connections)
			memoryPercent, _ := pv.MemoryPercent()
			fmt.Println(memoryPercent)

			exe, _ := pv.Exe()
			fmt.Println(fmt.Sprintf("exe：【%s】", exe))
			cmdline, _ := pv.Cmdline()
			fmt.Println(fmt.Sprintf("cmdline：【%s】", cmdline))

			iOCounters, _ := pv.IOCounters()
			fmt.Println(
				fmt.Sprintf(
					"iOCounters：读取总数【%d】，写入总数【%d】，读取量【%dMB】，写入量【%d】，", iOCounters.ReadBytes,
					iOCounters.WriteCount, iOCounters.ReadBytes/mbInt, iOCounters.WriteBytes/mbInt,
				),
			)
		}

	}
}
