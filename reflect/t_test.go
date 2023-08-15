package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func init() {
	// runtime.GOMAXPROCS(1)
}

func TestName(t *testing.T) {
	config := Config{}
	reflect.TypeOf(config).Name()
	fmt.Println(reflect.TypeOf(config).Name())
	fmt.Println(reflect.TypeOf(config).Kind())
	works := EsMultimediaMusicWorks{MusicUrl: "sdasdasd", CreateTime: time.Now()}
	updateMap := works.GUpdateMapJson()
	fmt.Println(updateMap)
	map3 := works.GUpdateMap3()
	works.GUpdateMapJsoniterJson()
	fmt.Println(map3)
	//	fmt.Println(reflect.TypeOf(config).NumOut())
}
func TestNameMap(t *testing.T) {
	m := make(map[string]string, 1024)
	fmt.Println(len(m))
}

func BenchmarkNew(b *testing.B) {
	var config *Config
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config = new(Config)
	}
	_ = config
}

func BenchmarkReflectNew(b *testing.B) {
	var config *Config
	typ := reflect.TypeOf(Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config, _ = reflect.New(typ).Interface().(*Config)
	}
	_ = config
}

// ---------------------------------------------//
func BenchmarkSet(b *testing.B) {
	config := new(Config)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config.Name = "name"
		config.IP = "ip"
		config.URL = "url"
		config.Timeout = "timeout"
	}
}

func BenchmarkReflect_FieldSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Field(0).SetString("name")
		ins.Field(1).SetString("ip")
		ins.Field(2).SetString("url")
		ins.Field(3).SetString("timeout")
	}
}

func BenchmarkReflect_FieldByNameSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.FieldByName("Name").SetString("name")
		ins.FieldByName("IP").SetString("ip")
		ins.FieldByName("URL").SetString("url")
		ins.FieldByName("Timeout").SetString("timeout")
	}
}

// ---------------------------------------------//

func BenchmarkReflect_EsMultimediaMusicWorks_Reflect(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		works := EsMultimediaMusicWorks{MusicUrl: "sdasdasd", CreateTime: time.Now()}
		updateMap := works.GUpdateMap()
		_ = updateMap
	}
}
func BenchmarkReflect_EsMultimediaMusicWorks_Set(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		works := EsMultimediaMusicWorks{MusicUrl: "sdasdasd", CreateTime: time.Now()}
		updateMap := works.GUpdateMapSet()
		_ = updateMap
	}
}
func BenchmarkReflect_EsMultimediaMusicWorks_JsoniterJson(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		works := EsMultimediaMusicWorks{MusicUrl: "sdasdasd", CreateTime: time.Now()}
		updateMap := works.GUpdateMapJsoniterJson()
		_ = updateMap
	}
}
func BenchmarkReflect_EsMultimediaMusicWorks_Json(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		works := EsMultimediaMusicWorks{MusicUrl: "sdasdasd", CreateTime: time.Now()}
		updateMap := works.GUpdateMapJson()
		_ = updateMap
	}
}
func BenchmarkReflect_EsMultimediaMusicWorks_FatihStructsMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		works := EsMultimediaMusicWorks{MusicUrl: "sdasdasd", CreateTime: time.Now()}
		updateMap := works.GUpdateMap3()
		_ = updateMap
	}
}
