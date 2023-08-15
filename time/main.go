package main

import (
	"fmt"
	"time"
)

const (
	TimeYyyyMmDd       = "2006-01-02"
	TimeYyyyMmDdHhMmSs = "2006-01-02 15:04:05.000"
)

// 获取当前时间距离明天0点的时间差
func GetCurrentTime2TomorrowSubUnix() (timeDifference int64) {
	now := time.Now()
	timeStr := now.Format(TimeYyyyMmDd)
	t, err := time.ParseInLocation(TimeYyyyMmDd, timeStr, time.Local)
	if err != nil {
		return 0
	}
	tomorrow := t.Add(time.Hour * 24)
	return tomorrow.Unix() - now.Unix()
}
func GetCurrentTime2TomorrowSubTime() (timeDifference time.Duration) {
	now := time.Now()
	timeStr := now.Format(TimeYyyyMmDd)
	t, err := time.ParseInLocation(TimeYyyyMmDd, timeStr, time.Local)
	if err != nil {
		return 0
	}
	tomorrow := t.Add(time.Hour * 24)
	return tomorrow.Sub(now)
}

func GetTomorrowTime() (timeDifference time.Time) {
	timeStr := time.Now().Format(TimeYyyyMmDd)
	t, err := time.ParseInLocation(TimeYyyyMmDd, timeStr, time.Local)
	if err != nil {
		return time.Time{}
	}
	return t.Add(time.Hour * 24)
}

// 获取当前时间所在月1号的00：00：00
func GetCurrentFirstDateOfMonth() time.Time {
	now := time.Now()
	return GetFirstDateOfMonth(now)
}
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetTime2ZeroTime(d)
}

// 获取当前时间所在月31/30号的00：00：00
func GetCurrentLastDateOfMonth() time.Time {
	return GetCurrentFirstDateOfMonth().AddDate(0, 1, -1)
}

func GetLastDateOfMonth(time2 time.Time) time.Time {
	return GetFirstDateOfMonth(time2).AddDate(0, 1, -1)
}

func FormatYyyyMmDdHhMmSsTime2String(time time.Time) string {
	return time.Format(TimeYyyyMmDdHhMmSs)
}

func GetPreDayTime() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// 获取某一天的0点时间
func GetTime2ZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetCurrentZeroTime() time.Time {
	d := time.Now()
	return GetTime2ZeroTime(d)
}

func main() {
	/*fmt.Println(GetCurrentTime2TomorrowSubUnix())
	fmt.Println(GetCurrentZeroTime())
	fmt.Println(GetPreDayTime())
	fmt.Println(GetTime2ZeroTime(GetPreDayTime()))
	fmt.Println(GetCurrentLastDateOfMonth())
	fmt.Println(GetCurrentFirstDateOfMonth())

	fmt.Println(406 & 32)

	timeD := 16848000
	oneDay := 60 * 60 * 24

	// oneHour := 60 * 60
	// "24H15M13S"  13+ (15*60) + (24*60*60)
	// "7M13S"  13+ (7*60) =443s
	fmt.Println(timeD & oneDay)
	// fmt.Println(timeD / oneDay)

	duration := time.Duration(443 * time.Second)
	fmt.Println(duration.String())*/
	int64s := make([]int64, 0, 0)
	int64s = append(int64s, 0, 0, 1, 0, 3, 2)
	conve(int64s)

	var a *name
	var namea nameAble
	var nameb nameAble = (*name)(nil)
	namea = a
	fmt.Println(a)
	fmt.Println(namea)
	fmt.Println(nameb)
	fmt.Println(a == namea)
	// nameb.name()
	// namea.name()
	defer func() {
		fmt.Println(1)
	}()
	defer add(3 + 4)
	defer add(5 + 6)
	defer func() {
		fmt.Println(2)
	}()

}
func add(c int) {
	fmt.Println(c)
}

var _ nameAble = (*name)(nil)

// var _ nameAble = *name
var _ nameAble = name{}

type name struct {
}

func (receiver name) name() {

}

type nameAble interface {
	name()
}

func conve(nums []int64) {
	right := 0
	left := 0

	for left < len(nums) {
		if nums[left] != 0 {
			tmp := nums[left]
			nums[left] = nums[right]
			nums[right] = tmp
			right++
		}
		left++
	}

	fmt.Println(nums)
}
