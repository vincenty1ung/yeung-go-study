package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/uncleyeung/yeung-go-study/csv/clan_pb"
	"github.com/uncleyeung/yeung-go-study/csv/i18n"
)

var (
	mapHeader = map[string][]string{
		i18n.DefaultLanguageCode: zh_Hant_header,
		i18n.LanguageCodeZHHans:  zh_Hans_header,
		i18n.LanguageCodeMs:      ms_header,
		i18n.LanguageCodeEn:      en_header,
		"id":                     id_header,
	}
	// 中文 zh-CN
	zh_Hans_header = []string{
		"Num", "序号",
		"Name", "昵称",
		"Id", "ID",
		"AdminTypeCsv", "成员类型",
		"RoomGoldLastWeek", "家族房间流水",
		"PeopleGoldLastWeek", "个人房间流水",
		"PersonalPatternGoldLastWeek", "直播模式流水",
		"ImGoldLastWeek", "IM礼物流水",
		"IntimateGoldLastWeek", "亲密礼物流水",
		"OnlyFansGoldLastWeek", "粉丝订阅收益",
		"PhoneGoldLastWeek", "心动8分钟收益",
		"LastWeekReceiveNum", "上周总接待人数",
		"PlayerRate", "上周主播回复率",
		"NewUserRate", "上周新人回复率",
		"ChairTimeWeekStr", "上周排档时长",
	}
	// 繁体 zh-HK
	zh_Hant_header = []string{
		"Num", "序號",
		"Name", "暱稱",
		"Id", "ID",
		"AdminTypeCsv", "成員類型",
		"RoomGoldLastWeek", "家族房間流水",
		"PeopleGoldLastWeek", "個人房間流水",
		"PersonalPatternGoldLastWeek", "直播模式流水",
		"ImGoldLastWeek", "IM禮物流水",
		"IntimateGoldLastWeek", "親密禮物流水",
		"OnlyFansGoldLastWeek", "粉絲訂閱收益",
		"PhoneGoldLastWeek", "心動8分鐘收益",
		"LastWeekReceiveNum", "上周總接待人數",
		"PlayerRate", "上周主播回覆率",
		"NewUserRate", "上周新人回覆率",
		"ChairTimeWeekStr", "上周排檔時長",
	}

	// 英文 en-US
	en_header = []string{
		"Num", "Serial number",
		"Name", "name",
		"Id", "ID",
		"AdminTypeCsv", "Membership",
		"RoomGoldLastWeek", "earnings of family room",
		"PeopleGoldLastWeek", "earnings of personal room",
		"PersonalPatternGoldLastWeek", "Live Mode",
		"ImGoldLastWeek", "Earnings of IM Gift",
		"IntimateGoldLastWeek", "Earnings of Intimate mode",
		"OnlyFansGoldLastWeek", "Fan Subscriptions Earnings",
		"PhoneGoldLastWeek", "8 Minute Heartbeat Earnings",
		"LastWeekReceiveNum", "Reception last week",
		"PlayerRate", "Host response rate",
		"NewUserRate", "New response rate",
		"ChairTimeWeekStr", "Last Week Reception Number",
	}
	// 马来 ms-MY
	ms_header = []string{
		"Num", "Nombor siri",
		"Name", "pengguna",
		"Id", "ID",
		"AdminTypeCsv", "Jenis ahli",
		"RoomGoldLastWeek", "Pendapatan bilik keluarga",
		"PeopleGoldLastWeek", "Pendapatan bilik peribadi",
		"PersonalPatternGoldLastWeek", "Mod Live",
		"ImGoldLastWeek", "Pendapatan hadiah IM",
		"IntimateGoldLastWeek", "Pendapatan hadiah akrab",
		"OnlyFansGoldLastWeek", "Pendapatan Langganan Fan",
		"PhoneGoldLastWeek", "Pendapatan Tangkapmu 8 minit",
		"LastWeekReceiveNum", "Minggu ini",
		"PlayerRate", "Kadar balasan hos",
		"NewUserRate", "Kadar balas pengguna baru",
		"ChairTimeWeekStr", "Jumlah Sambutan Minggu Lalu",
	}
	// 印尼 id-ID
	id_header = []string{
		"Num", "Nomor seri:",
		"Name", "nama panggilan",
		"Id", "ID",
		"AdminTypeCsv", "Tipe Anggota",
		"RoomGoldLastWeek", "Penghasilan ruangan Family",
		"PeopleGoldLastWeek", "Penghasilan ruangan saya",
		"PersonalPatternGoldLastWeek", "Mod Siaran Langsung",
		"ImGoldLastWeek", "Penghasilan hadiah IM",
		"IntimateGoldLastWeek", "Penghasilandari Hadiah Kedekatan",
		"OnlyFansGoldLastWeek", "Perhasilkan Subscribe Pengemar",
		"PhoneGoldLastWeek", "Aliran 8 Menit Tangkapmu",
		"LastWeekReceiveNum", "Jumlah total orang yang diterima minggu yg lalu",
		"PlayerRate", "Persentase balasan dari host",
		"NewUserRate", "Persentase balasan pengguna baru",
		"ChairTimeWeekStr", "Waktu Shift",
	}
	mapHeaderV2 = map[string][]string{
		i18n.DefaultLanguageCode: zh_Hant_headerV2,
		i18n.LanguageCodeZHHans:  zh_Hans_headerV2,
		i18n.LanguageCodeMs:      ms_headerV2,
		i18n.LanguageCodeEn:      en_headerV2,
		"id":                     id_headerV2,
	}
	// 中文 zh-CN
	zh_Hans_headerV2 = []string{
		"num", "序号",
		"name", "昵称",
		"id", "ID",
		"admin_type_csv", "成员类型",
		"room_gold_last_week", "家族房间流水",
		"people_gold_last_week", "个人房间流水",
		"personal_pattern_gold_last_week", "直播模式流水",
		"im_gold_last_week", "IM礼物流水",
		"intimate_gold_last_week", "亲密礼物流水",
		"only_fans_gold_last_week", "粉丝订阅收益",
		"phone_gold_last_week", "心动8分钟收益",
		"last_week_receive_num", "上周总接待人数",
		"player_rate", "上周主播回复率",
		"new_user_rate", "上周新人回复率",
		"chair_time_week_str", "上周排档时长",
	}
	// 繁体 zh-HK
	zh_Hant_headerV2 = []string{
		"Num", "序號",
		"Name", "暱稱",
		"Id", "ID",
		"AdminTypeCsv", "成員類型",
		"RoomGoldLastWeek", "家族房間流水",
		"PeopleGoldLastWeek", "個人房間流水",
		"PersonalPatternGoldLastWeek", "直播模式流水",
		"ImGoldLastWeek", "IM禮物流水",
		"IntimateGoldLastWeek", "親密禮物流水",
		"OnlyFansGoldLastWeek", "粉絲訂閱收益",
		"PhoneGoldLastWeek", "心動8分鐘收益",
		"LastWeekReceiveNum", "上周總接待人數",
		"PlayerRate", "上周主播回覆率",
		"NewUserRate", "上周新人回覆率",
		"ChairTimeWeekStr", "上周排檔時長",
	}

	// 英文 en-US
	en_headerV2 = []string{
		"Num", "Serial number",
		"Name", "name",
		"Id", "ID",
		"AdminTypeCsv", "Membership",
		"RoomGoldLastWeek", "earnings of family room",
		"PeopleGoldLastWeek", "earnings of personal room",
		"PersonalPatternGoldLastWeek", "Live Mode",
		"ImGoldLastWeek", "Earnings of IM Gift",
		"IntimateGoldLastWeek", "Earnings of Intimate mode",
		"OnlyFansGoldLastWeek", "Fan Subscriptions Earnings",
		"PhoneGoldLastWeek", "8 Minute Heartbeat Earnings",
		"LastWeekReceiveNum", "Reception last week",
		"PlayerRate", "Host response rate",
		"NewUserRate", "New response rate",
		"ChairTimeWeekStr", "Last Week Reception Number",
	}
	// 马来 ms-MY
	ms_headerV2 = []string{
		"Num", "Nombor siri",
		"Name", "pengguna",
		"Id", "ID",
		"AdminTypeCsv", "Jenis ahli",
		"RoomGoldLastWeek", "Pendapatan bilik keluarga",
		"PeopleGoldLastWeek", "Pendapatan bilik peribadi",
		"PersonalPatternGoldLastWeek", "Mod Live",
		"ImGoldLastWeek", "Pendapatan hadiah IM",
		"IntimateGoldLastWeek", "Pendapatan hadiah akrab",
		"OnlyFansGoldLastWeek", "Pendapatan Langganan Fan",
		"PhoneGoldLastWeek", "Pendapatan Tangkapmu 8 minit",
		"LastWeekReceiveNum", "Minggu ini",
		"PlayerRate", "Kadar balasan hos",
		"NewUserRate", "Kadar balas pengguna baru",
		"ChairTimeWeekStr", "Jumlah Sambutan Minggu Lalu",
	}
	// 印尼 id-ID
	id_headerV2 = []string{
		"Num", "Nomor seri:",
		"Name", "nama panggilan",
		"Id", "ID",
		"AdminTypeCsv", "Tipe Anggota",
		"RoomGoldLastWeek", "Penghasilan ruangan Family",
		"PeopleGoldLastWeek", "Penghasilan ruangan saya",
		"PersonalPatternGoldLastWeek", "Mod Siaran Langsung",
		"ImGoldLastWeek", "Penghasilan hadiah IM",
		"IntimateGoldLastWeek", "Penghasilandari Hadiah Kedekatan",
		"OnlyFansGoldLastWeek", "Perhasilkan Subscribe Pengemar",
		"PhoneGoldLastWeek", "Aliran 8 Menit Tangkapmu",
		"LastWeekReceiveNum", "Jumlah total orang yang diterima minggu yg lalu",
		"PlayerRate", "Persentase balasan dari host",
		"NewUserRate", "Persentase balasan pengguna baru",
		"ChairTimeWeekStr", "Waktu Shift",
	}
	bigMasterMap = map[string]string{
		i18n.DefaultLanguageCode: "家族长",
		i18n.LanguageCodeZHHans:  "家族长",
		i18n.LanguageCodeMs:      "Ketua Keluarga",
		i18n.LanguageCodeEn:      "Family Leader",
		"id":                     "Kepala Family",
	}
	bigAdminMap = map[string]string{
		i18n.DefaultLanguageCode: "高級管理員",
		i18n.LanguageCodeZHHans:  "高级管理员",
		i18n.LanguageCodeMs:      "Admin Kanan",
		i18n.LanguageCodeEn:      "Senior Admini",
		"id":                     "Admin Tingkat Tinggi",
	}
	adminMap = map[string]string{
		i18n.DefaultLanguageCode: "管理員",
		i18n.LanguageCodeZHHans:  "管理员",
		i18n.LanguageCodeMs:      "Admin",
		i18n.LanguageCodeEn:      "Admini",
		"id":                     "Admin",
	}
	specialMap = map[string]string{
		i18n.DefaultLanguageCode: "排檔成員",
		i18n.LanguageCodeZHHans:  "排档成员",
		i18n.LanguageCodeMs:      "Ahli Host",
		i18n.LanguageCodeEn:      "Host Member",
		"id":                     "Anggota shift",
	}
	normalMap = map[string]string{
		i18n.DefaultLanguageCode: "普通成員",
		i18n.LanguageCodeZHHans:  "普通成员",
		i18n.LanguageCodeMs:      "Ahli Normal",
		i18n.LanguageCodeEn:      "Regular Member",
		"id":                     "Anggota Biasa",
	}
)

type ClanMemberExtends []*ClanMemberExtend

// 扩展
type ClanMemberExtend struct {
	Id                          int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`   // playerId
	Id2                         int64                 `protobuf:"varint,2,opt,name=id2,proto3" json:"id2,omitempty"` // 靓号
	Name                        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Icon                        string                `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`                                                                            // 头像
	Sex                         int32                 `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`                                                                             // 性别 参考 common.SexType , 0未知/1男/2女
	RoomGoldWeek                int32                 `protobuf:"varint,6,opt,name=room_gold_week,json=roomGoldWeek,proto3" json:"room_gold_week,omitempty"`                                     // 本周房间流水
	PeopleGoldWeek              int32                 `protobuf:"varint,7,opt,name=people_gold_week,json=peopleGoldWeek,proto3" json:"people_gold_week,omitempty"`                               // 本周个人流水
	Ticket                      int32                 `protobuf:"varint,8,opt,name=ticket,proto3" json:"ticket,omitempty"`                                                                       // 贡献值钻石，（公会收益界面使用）
	BillGoldWeek                int32                 `protobuf:"varint,9,opt,name=bill_gold_week,json=billGoldWeek,proto3" json:"bill_gold_week,omitempty"`                                     // 本周点单流水
	PersonalPatternGoldWeek     int32                 `protobuf:"varint,10,opt,name=personal_pattern_gold_week,json=personalPatternGoldWeek,proto3" json:"personal_pattern_gold_week,omitempty"` // 本周个人模式收益
	IsBillAdmin                 bool                  `protobuf:"varint,11,opt,name=is_bill_admin,json=isBillAdmin,proto3" json:"is_bill_admin,omitempty"`                                       // 是否点单管理员
	ImGiftGoldWeek              int32                 `protobuf:"varint,12,opt,name=im_gift_gold_week,json=imGiftGoldWeek,proto3" json:"im_gift_gold_week,omitempty"`                            // 本周IM礼物流水
	IntimateGoldWeek            int32                 `protobuf:"varint,13,opt,name=intimate_gold_week,json=intimateGoldWeek,proto3" json:"intimate_gold_week,omitempty"`                        // 本周亲密模式流水
	SongOrderGoldWeek           int32                 `protobuf:"varint,14,opt,name=song_order_gold_week,json=songOrderGoldWeek,proto3" json:"song_order_gold_week,omitempty"`                   // 本周点单流水
	AdminType                   clan_pb.ClanAdminType `protobuf:"varint,15,opt,name=admin_type,json=adminType,proto3,enum=clan_pb.ClanAdminType" json:"admin_type,omitempty"`                    // 会员类型
	ChairTimeToday              int32                 `protobuf:"varint,16,opt,name=chair_time_today,json=chairTimeToday,proto3" json:"chair_time_today,omitempty"`                              // 今天上麦时长，单位：秒
	ChairTimeWeek               int32                 `protobuf:"varint,17,opt,name=chair_time_week,json=chairTimeWeek,proto3" json:"chair_time_week,omitempty"`                                 // 本周上麦时长，单位：秒
	LastWeekReceiveNum          int32                 `protobuf:"varint,18,opt,name=last_week_receive_num,json=lastWeekReceiveNum,proto3" json:"last_week_receive_num,omitempty"`                // 上周接待数
	WeekReceiveNum              int32                 `protobuf:"varint,19,opt,name=week_receive_num,json=weekReceiveNum,proto3" json:"week_receive_num,omitempty"`                              // 本周接待数
	OnlyFansGoldWeek            int32                 `protobuf:"varint,20,opt,name=only_fans_gold_week,json=onlyFansGoldWeek,proto3" json:"only_fans_gold_week,omitempty"`                      // 本周粉丝订阅流水
	PhoneGoldWeek               int32                 `protobuf:"varint,21,opt,name=phone_gold_week,json=phoneGoldWeek,proto3" json:"phone_gold_week,omitempty"`                                 // 本周心动8分钟（连麦）流水
	NewUserRate                 int32                 `protobuf:"varint,22,opt,name=new_user_rate,json=newUserRate,proto3" json:"new_user_rate,omitempty"`                                       // 新用户回复率
	PlayerRate                  int32                 `protobuf:"varint,23,opt,name=player_rate,json=playerRate,proto3" json:"player_rate,omitempty"`                                            // 主播回复率
	AuctionGoldWeek             int32                 `protobuf:"varint,24,opt,name=auction_gold_week,json=auctionGoldWeek,proto3" json:"auction_gold_week,omitempty"`                           // 本周拍拍房流水
	AdminTypeCsvMap             map[string]string     `protobuf:"bytes,25,rep,name=admin_type_csv_map,json=adminTypeCsvMap,proto3" json:"admin_type_csv_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AdminTypeCsv                string                `protobuf:"bytes,26,opt,name=admin_type_csv,json=adminTypeCsv,proto3" json:"admin_type_csv,omitempty"`
	MemberType                  int32                 `protobuf:"varint,27,opt,name=member_type,json=memberType,proto3" json:"member_type,omitempty"`                                                          // 会员类型
	RoomGoldLastWeek            int32                 `protobuf:"varint,28,opt,name=room_gold_last_week,json=roomGoldLastWeek,proto3" json:"room_gold_last_week,omitempty"`                                    // 家族房间流水+",
	PeopleGoldLastWeek          int32                 `protobuf:"varint,29,opt,name=people_gold_last_week,json=peopleGoldLastWeek,proto3" json:"people_gold_last_week,omitempty"`                              // 个人房间流水+",
	PersonalPatternGoldLastWeek int32                 `protobuf:"varint,30,opt,name=personal_pattern_gold_last_week,json=personalPatternGoldLastWeek,proto3" json:"personal_pattern_gold_last_week,omitempty"` // 直播模式流水+",
	ImGoldLastWeek              int32                 `protobuf:"varint,31,opt,name=im_gold_last_week,json=imGoldLastWeek,proto3" json:"im_gold_last_week,omitempty"`                                          // IM礼物流水+",
	IntimateGoldLastWeek        int32                 `protobuf:"varint,32,opt,name=intimate_gold_last_week,json=intimateGoldLastWeek,proto3" json:"intimate_gold_last_week,omitempty"`                        // 亲密礼物流水+",
	OnlyFansGoldLastWeek        int32                 `protobuf:"varint,33,opt,name=only_fans_gold_last_week,json=onlyFansGoldLastWeek,proto3" json:"only_fans_gold_last_week,omitempty"`                      // 粉丝订阅收益+",
	PhoneGoldLastWeek           int32                 `protobuf:"varint,34,opt,name=phone_gold_last_week,json=phoneGoldLastWeek,proto3" json:"phone_gold_last_week,omitempty"`                                 // 心动8分钟收益+",
	ChairTimeWeekStr            string                `json:"chair_time_week_str"`
	Num                         uint32                `json:"num"`
}

type CsvData interface {
	Len() int
	Generate(header []string) []map[string]string
	GenerateV2(header []string) *csv.Writer
}

func (c ClanMemberExtends) Len() int {
	return len(c)
}

func (c ClanMemberExtends) Generate(header []string) []map[string]string {
	headerMap := make(map[string]int)
	for i, s := range header {
		headerMap[s] = i
	}
	res := make([]map[string]string, 0, c.Len())
	for _, extend := range c {
		m := make(map[string]string)
		of := reflect.TypeOf(extend)
		// valueOf := reflect.ValueOf(extend)
		// switch valueOf.Kind() {
		switch of.Kind() {
		case reflect.Pointer:
			elem := of.Elem()
			// value := valueOf.Elem()
			// for j := 0; j < value.NumField()-1; j++ {
			for j := 0; j < elem.NumField()-1; j++ {
				// field := valueOf.Field(j)
				structField := elem.Field(j)
				if v, ok := structField.Tag.Lookup("json"); ok {
					split := strings.Split(v, ",")
					if len(split) > 1 {
						v = split[0]
					}
					if _, ok := headerMap[v]; ok {
						u := uintptr(unsafe.Pointer(extend)) + structField.Offset
						var s any
						switch structField.Type.Kind() {
						case reflect.Int64:
							s = *(*int64)(unsafe.Pointer(u))
						case reflect.Int32:
							s = *(*int32)(unsafe.Pointer(u))
						case reflect.String:
							s = *(*string)(unsafe.Pointer(u))
						}
						// m[v] = fmt.Sprint(value.Field(j))
						m[v] = fmt.Sprint(s)
						// break
					}

				}
			}
		case reflect.Struct:

		}
		res = append(res, m)
	}
	return res
}

func (c ClanMemberExtends) GenerateV2(header []string) *csv.Writer {
	buf := new(bytes.Buffer)

	// 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码 并没有用
	buf.Write([]byte{byte(0xef), byte(0xbb), byte(0xbf)})

	w := csv.NewWriter(buf)
	keys := make([]string, 0)
	head := make([]string, 0)
	headMap := make(map[string]struct{}, len(header))
	for i := 1; i < len(header); i = i + 2 {
		headMap[header[i-1]] = struct{}{}
		keys = append(keys, header[i-1])
		head = append(head, header[i])
	}
	if true {
		w.Write(head)
	}

	// res := make([]map[string]string, 0, c.Len())
	for _, extend := range c {
		// m := make(map[string]string)
		of := reflect.TypeOf(extend)
		// valueOf := reflect.ValueOf(extend)
		// switch valueOf.Kind() {
		row := make([]string, 0)
		switch of.Kind() {
		case reflect.Pointer:
			elem := of.Elem()
			// value := valueOf.Elem()
			// for j := 0; j < value.NumField()-1; j++ {

			for j := 0; j < elem.NumField()-1; j++ {
				// field := valueOf.Field(j)
				structField := elem.Field(j)
				if _, ok := headMap[structField.Name]; ok {
					u := uintptr(unsafe.Pointer(extend)) + structField.Offset
					var s any
					switch structField.Type.Kind() {
					case reflect.Int64:
						s = *(*int64)(unsafe.Pointer(u))
					case reflect.Int32:
						s = *(*int32)(unsafe.Pointer(u))
					case reflect.String:
						s = *(*string)(unsafe.Pointer(u))
					}
					// m[v] = fmt.Sprint(value.Field(j))
					row = append(row, fmt.Sprint(s))
					// break
				}

			}
		case reflect.Struct:

		}
		// res = append(res, m)
		w.Write(row)
	}
	return w
}

func WriterCSVV2(head, keys []string, mapV []map[string]string, first bool) []byte {
	buf := new(bytes.Buffer)

	// 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码 并没有用
	buf.Write([]byte{byte(0xef), byte(0xbb), byte(0xbf)})

	w := csv.NewWriter(buf)
	/*keys := make([]string, 0)
	head := make([]string, 0)
	for i := 1; i < len(header); i = i + 2 {
		keys = append(keys, header[i-1])
		head = append(head, header[i])
	}*/
	if first {
		w.Write(head)
	}
	// generateMap := able.Generate(keys)
	for _, m := range mapV {
		row := make([]string, 0)
		for _, key := range keys {
			if v, ok := m[key]; ok {
				row = append(row, v)
			}
		}
		w.Write(row)
	}

	w.Flush()
	return buf.Bytes()
}

// header: []string{导出的字段名，表头，导出的字段名，表头...}
// data: 一个slice，slice的数据可以是自定义的struct结构，也可以是一个map结构，map的key应该为string类型
// -- 导出的字段名 必须与struct的字段名一致，或map的key值一直
func WriterCSV(header []string, data interface{}, first bool) []byte {
	buf := new(bytes.Buffer)

	// 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码 并没有用
	buf.Write([]byte{byte(0xef), byte(0xbb), byte(0xbf)})

	w := csv.NewWriter(buf)
	keys := make([]string, 0)
	head := make([]string, 0)
	headMap := make(map[string]struct{}, len(header))
	for i := 1; i < len(header); i = i + 2 {
		headMap[header[i-1]] = struct{}{}
		keys = append(keys, header[i-1])
		head = append(head, header[i])
	}
	if first {
		w.Write(head)
	}

	value := reflect.ValueOf(data)
	l := value.Len()
	for i := 0; i < l; i++ {
		item := value.Index(i)
		if item.IsValid() {
			row := make([]string, 0)
			if item.Kind() == reflect.Ptr || item.Kind() == reflect.Struct {
				elem := item.Elem()
				for _, k := range keys {
					fv := elem.FieldByName(k)
					if fv.IsValid() {
						// fmt.Println(fmt.Sprint(fv))
						// fmt.Println(fv.Int())
						row = append(row, fmt.Sprint(fv))
					} else {
						row = append(row, "")
					}
				}
			} else if item.Kind() == reflect.Map {
				for _, k := range keys {
					fv := item.MapIndex(reflect.ValueOf(k))
					if fv.IsValid() {
						row = append(row, fmt.Sprint(fv))
					} else {
						row = append(row, "")
					}
				}
			}
			w.Write(row)
		}
	}

	w.Flush()
	return buf.Bytes()
}
