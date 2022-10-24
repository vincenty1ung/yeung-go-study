package reflect

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/structs"
	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}

func getConfig() *Config {
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		typ.NumOut()
		if v, ok := f.Tag.Lookup("json"); ok {
			tmp := strings.ReplaceAll(strings.ToUpper(v), "-", "_")
			key := fmt.Sprintf("CONFIG_%s", tmp)
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}

func ReflectDemo() {
	os.Setenv("CONFIG_SERVER_NAME", "Go-Server")
	os.Setenv("CONFIG_SERVER_IP", "10.10.10.11")
	os.Setenv("CONFIG_SERVER_URL", "chende.com")
	cfg := getConfig()
	fmt.Printf("%+v", cfg)
}

type (
	EsMultimediaMusicWorks struct {
		// 基础区
		Anonymous   int32     `gorm:"column:anonymous" json:"anonymous"`       // 是否匿名作品0否，1是
		AuditStatus int32     `gorm:"column:audit_status" json:"audit_status"` // 1:待审核,2:审核失败,3:审核成功
		CommentNum  int32     `gorm:"column:comment_num" json:"comment_num"`   // 评论数量
		CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`   // 创建时间
		Deleted     int32     `gorm:"column:deleted" json:"deleted"`           // 是否被删除, 1:已删除
		LikeNum     int32     `gorm:"column:like_num" json:"like_num"`         // 点赞数量
		MusicName   string    `gorm:"column:music_name" json:"music_name"`     // 歌曲名称
		MusicType   int32     `gorm:"column:music_type" json:"music_type"`     // 歌曲类型1:原唱,2:翻唱
		UploadType  int32     `gorm:"column:upload_type" json:"upload_type"`   // 上传类型1:用户上传,2:后台上传
		// NoticePlayers   string    `gorm:"column:notice_players" json:"notice_players"`     // @通知的人列表
		Rejection       string    `gorm:"column:rejection" json:"rejection"`               // 拒绝理由
		PlayerCount     int32     `gorm:"column:player_count" json:"player_count"`         // 播放次数
		PlayerId        string    `gorm:"column:player_id" json:"player_id"`               // 歌手id
		RecommendStatus int32     `gorm:"column:recommend_status" json:"recommend_status"` // 0:不推荐, 1:推荐
		RelayNum        int32     `gorm:"column:relay_num" json:"relay_num"`               // 转发数量
		ShareNum        int32     `gorm:"column:share_num" json:"share_num"`               // 分享数量
		ShelfStatus     int32     `gorm:"column:shelf_status" json:"shelf_status"`         // 0:下架, 1:上架
		Sort            int32     `gorm:"column:sort" json:"sort"`                         // 排序
		Status          int32     `gorm:"column:status" json:"status"`                     // 0:正常, 1:疑似涉黄
		Tags            string    `gorm:"column:tags" json:"tags"`                         // 标签id列表
		TopStatus       int32     `gorm:"column:top_status" json:"top_status"`             // 0:不置顶1:置顶
		UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`           // 更新时间
		WorksName       string    `gorm:"column:works_name" json:"works_name"`             // 作品名称
		Id              int64     `gorm:"column:id" json:"id"`                             // 作品Id
		PlayNum         int32     `gorm:"column:play_num" json:"play_num"`                 // 完播次数
		OffShelfReason  string    `gorm:"column:off_shelf_reason" json:"off_shelf_reason"` // 下架原因
		ReportDetailId  int64     `gorm:"column:report_detail_id" json:"report_detail_id"` // 下架上报原因id
		CollectNum      int32     `gorm:"column:collect_num" json:"collect_num"`           // 收藏数
		MusicHot        int64     `gorm:"column:music_hot" json:"music_hot"`               // 音乐热度

		// Media区
		CoverUrl    string `gorm:"column:cover_url" json:"cover_url"`       // 封面url
		Duration    int32  `gorm:"column:duration" json:"duration"`         // 持续时间
		ExtraUrl1   string `gorm:"column:extra_url1" json:"extra_url1"`     // 额外url
		GreenRate   int64  `gorm:"column:green_rate" json:"green_rate"`     // 鉴黄分数
		GreenScenes string `gorm:"column:green_scenes" json:"green_scenes"` // 鉴黄场景
		GreenStatus int64  `gorm:"column:green_status" json:"green_status"` // 鉴黄状态
		LyricUrl    string `gorm:"column:lyric_url" json:"lyric_url"`       // 歌词地址
		Media       string `gorm:"column:media" json:"media"`               // 资源信息json序列化
		MusicUrl    string `gorm:"column:music_url" json:"music_url"`       // 音乐url
	}
)

func (receiver EsMultimediaMusicWorks) GUpdateMap() map[string]interface{} {
	receiverType := reflect.TypeOf(receiver)
	receiverValue := reflect.ValueOf(receiver)
	m := make(map[string]interface{}, receiverType.NumField())
	for i := 0; i < receiverType.NumField(); i++ {
		field := receiverType.Field(i)
		value := receiverValue.Field(i)
		if value.IsValid() {
			m[field.Tag.Get("json")] = fmt.Sprint(value)
		} else {
			m[field.Tag.Get("json")] = ""
		}
	}
	return m
}

func (receiver EsMultimediaMusicWorks) GUpdateMapSet() map[string]interface{} {
	m := make(map[string]interface{}, 1<<8-1)
	m["anonymous"] = receiver.Anonymous
	m["audit_status"] = receiver.AuditStatus
	m["comment_num"] = receiver.CommentNum
	m["create_time"] = receiver.CreateTime
	m["Deleted"] = receiver.Deleted
	m["LikeNum"] = receiver.LikeNum
	m["MusicName"] = receiver.MusicName
	m["MusicType"] = receiver.MusicType
	m["UploadType"] = receiver.UploadType
	m["Rejection"] = receiver.Rejection
	m["PlayerCount"] = receiver.PlayerCount
	m["PlayerId"] = receiver.PlayerId
	m["RecommendStatus"] = receiver.RecommendStatus
	m["RelayNum"] = receiver.RelayNum
	m["ShareNum"] = receiver.ShareNum
	m["ShelfStatus"] = receiver.ShelfStatus
	m["Sort"] = receiver.Sort
	m["Tags"] = receiver.Tags
	m["Status"] = receiver.Status
	m["Id"] = receiver.Id
	m["TopStatus"] = receiver.TopStatus
	m["UpdateTime"] = receiver.UpdateTime
	m["PlayNum"] = receiver.PlayNum
	m["OffShelfReason"] = receiver.OffShelfReason
	m["ReportDetailId"] = receiver.ReportDetailId
	m["CollectNum"] = receiver.CollectNum
	m["MusicHot"] = receiver.MusicHot
	m["CoverUrl"] = receiver.CoverUrl
	m["Duration"] = receiver.Duration
	m["ExtraUrl1"] = receiver.ExtraUrl1
	m["GreenRate"] = receiver.GreenRate
	m["GreenScenes"] = receiver.GreenScenes
	m["LyricUrl"] = receiver.LyricUrl
	m["Media"] = receiver.Media
	m["MusicUrl"] = receiver.MusicUrl
	return m
}

func (receiver EsMultimediaMusicWorks) GUpdateMapJsoniterJson() map[string]interface{} {
	m := make(map[string]interface{}, 1<<8-1)
	data, _ := jsoniter.Marshal(&receiver)
	err := jsoniter.Unmarshal(data, &m)
	if err != nil {
		return nil
	}
	return m
}
func (receiver EsMultimediaMusicWorks) GUpdateMapJson() map[string]interface{} {
	m := make(map[string]interface{}, 1<<8-1)
	data, _ := json.Marshal(&receiver)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil
	}
	return m
}

func (receiver EsMultimediaMusicWorks) GUpdateMap3() map[string]interface{} {
	// structs.Map(receiver)
	return structs.Map(receiver)
}
