package main

import (
	"bytes"
	"fmt"
	"net"
	"sort"
	"sync"

	"github.com/oschwald/geoip2-golang"
)

var (
	ip1 = net.ParseIP("216.14.49.184")
	ip2 = net.ParseIP("216.14.49.191")
)

func check(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		fmt.Printf("%v is not an IPv4 address\n", trial)
		return false
	}
	if bytes.Compare(trial, ip1) >= 0 && bytes.Compare(trial, ip2) <= 0 {
		fmt.Printf("%v is between %v and %v\n", trial, ip1, ip2)
		return true
	}
	fmt.Printf("%v is NOT between %v and %v\n", trial, ip1, ip2)
	return false
}

var GeoDb *geoip2.Reader
var once sync.Once

// var geoDbLock sync.Mutex

func GeoInit(dir string) {
	var err error
	once.Do(
		func() {
			GeoDb, err = geoip2.Open(dir)
			if err != nil {
				panic(err)
			}
		},
	)

}
func GeoDbGetInfo(ipStr string) string {

	ip := net.ParseIP(ipStr)
	record, _ := GeoDb.City(ip)

	fmt.Println(ipStr)
	fmt.Println(record.Country.IsoCode)
	/*fmt.Println(record.Country.Names["zh-CN"])
	fmt.Println(ipStr)*/
	return record.Country.IsoCode

}

type man struct {
	ip   string
	aeg  int
	long int
	name string
}

func (m man) String() string {
	return fmt.Sprintf("name:%s ,age:%d ,ip:%s ,long:%d;", m.name, m.aeg, m.ip, m.long)
}

func main() {
	check("1.2.3.4")
	check("216.14.49.185")
	check("1::16")

	GeoInit("./test/net/GeoLite2-City.mmdb")
	GeoDbGetInfo("216.14.49.185")
	GeoDbGetInfo("216.14.59.185")
	GeoDbGetInfo("156.249.14.100")

	mens := make([]man, 0, 10)
	mens = append(
		mens, man{
			name: "yang c",
			aeg:  10,
			long: 7,

			ip: "216.14.49.185",
		},
	)
	mens = append(
		mens, man{
			name: "yang1",
			aeg:  12,
			long: 1,
			ip:   "216.14.59.185",
		},
	)
	mens = append(
		mens, man{
			aeg:  22,
			long: 8,
			name: "wang",
			ip:   "216.14.49.185",
		},
	)
	mens = append(
		mens, man{
			name: "ang",
			aeg:  32,
			long: 12,
			ip:   "156.249.14.100",
		},
	)
	mens = append(
		mens, man{
			name: "89ang",
			aeg:  32,
			long: 12,
			ip:   "133.249.14.1",
		},
	)
	mens = append(
		mens, man{
			name: "yang c",
			aeg:  10,
			long: 6,

			ip: "216.14.49.185",
		},
	)
	mens = append(
		mens, man{
			name: "yang2",
			aeg:  12,
			long: 6,
			ip:   "156.249.14.100",
		},
	)
	mens = append(
		mens, man{
			name: "yang909",
			aeg:  12,
			long: 6,
			ip:   "110.249.14.100",
		},
	)
	mens = append(
		mens, man{
			name: "yang99",
			aeg:  78,
			long: 90,
			ip:   "120.230.80.28",
		},
	)

	fmt.Println("==================")
	fmt.Println(mens)

	arg := ""
	sort.Slice(
		mens, func(i, j int) bool {
			if GeoDbGetInfo(mens[i].ip) == GeoDbGetInfo(mens[j].ip) ||
				(GeoDbGetInfo(mens[i].ip) != arg && GeoDbGetInfo(mens[j].ip) != arg) {
				if mens[i].aeg == mens[j].aeg {
					return mens[i].long > mens[j].long
				}
				return mens[i].aeg > mens[j].aeg
			}

			return GeoDbGetInfo(mens[i].ip) == arg
		},
	)
	fmt.Println("==================")
	fmt.Println(mens)

	i := int64(1647255317000)
	fmt.Println(i)

	i2 := i >> 3
	fmt.Println(i2)
}
