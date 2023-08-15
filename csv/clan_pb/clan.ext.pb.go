package clan_pb

import (
	"fmt"

	"google.golang.org/protobuf/runtime/protoimpl"
)

// 公会管理员类型
type ClanAdminType int32

const (
	ClanAdminType_CAT_GUEST      ClanAdminType = 0  // 占位
	ClanAdminType_CAT_NORMAL     ClanAdminType = 10 // 普通成员
	ClanAdminType_CAT_ANCHOR     ClanAdminType = 11 // 排挡成员
	ClanAdminType_CAT_ADMIN      ClanAdminType = 20 // 管理员
	ClanAdminType_CAT_HIGH_ADMIN ClanAdminType = 21 // 高级管理员
	ClanAdminType_CAT_OWNER      ClanAdminType = 30 // 会长
	ClanAdminType_CAT_BIG_FAMILY ClanAdminType = 40 // 大家族长
)

// Enum value maps for ClanAdminType.
var (
	ClanAdminType_name = map[int32]string{
		0:  "CAT_GUEST",
		10: "CAT_NORMAL",
		11: "CAT_ANCHOR",
		20: "CAT_ADMIN",
		21: "CAT_HIGH_ADMIN",
		30: "CAT_OWNER",
		40: "CAT_BIG_FAMILY",
	}
	ClanAdminType_value = map[string]int32{
		"CAT_GUEST":      0,
		"CAT_NORMAL":     10,
		"CAT_ANCHOR":     11,
		"CAT_ADMIN":      20,
		"CAT_HIGH_ADMIN": 21,
		"CAT_OWNER":      30,
		"CAT_BIG_FAMILY": 40,
	}
)

func (x ClanAdminType) Enum() *ClanAdminType {
	p := new(ClanAdminType)
	*p = x
	return p
}

type ClanMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                          int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`   // playerId
	Id2                         int64             `protobuf:"varint,2,opt,name=id2,proto3" json:"id2,omitempty"` // 靓号
	Name                        string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Icon                        string            `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`                                                                            // 头像
	Sex                         int32             `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`                                                                             // 性别 参考 common.SexType , 0未知/1男/2女
	RoomGoldWeek                int32             `protobuf:"varint,6,opt,name=room_gold_week,json=roomGoldWeek,proto3" json:"room_gold_week,omitempty"`                                     // 本周房间流水
	PeopleGoldWeek              int32             `protobuf:"varint,7,opt,name=people_gold_week,json=peopleGoldWeek,proto3" json:"people_gold_week,omitempty"`                               // 本周个人流水
	Ticket                      int32             `protobuf:"varint,8,opt,name=ticket,proto3" json:"ticket,omitempty"`                                                                       // 贡献值钻石，（公会收益界面使用）
	BillGoldWeek                int32             `protobuf:"varint,9,opt,name=bill_gold_week,json=billGoldWeek,proto3" json:"bill_gold_week,omitempty"`                                     // 本周点单流水
	PersonalPatternGoldWeek     int32             `protobuf:"varint,10,opt,name=personal_pattern_gold_week,json=personalPatternGoldWeek,proto3" json:"personal_pattern_gold_week,omitempty"` // 本周个人模式收益
	IsBillAdmin                 bool              `protobuf:"varint,11,opt,name=is_bill_admin,json=isBillAdmin,proto3" json:"is_bill_admin,omitempty"`                                       // 是否点单管理员
	ImGiftGoldWeek              int32             `protobuf:"varint,12,opt,name=im_gift_gold_week,json=imGiftGoldWeek,proto3" json:"im_gift_gold_week,omitempty"`                            // 本周IM礼物流水
	IntimateGoldWeek            int32             `protobuf:"varint,13,opt,name=intimate_gold_week,json=intimateGoldWeek,proto3" json:"intimate_gold_week,omitempty"`                        // 本周亲密模式流水
	SongOrderGoldWeek           int32             `protobuf:"varint,14,opt,name=song_order_gold_week,json=songOrderGoldWeek,proto3" json:"song_order_gold_week,omitempty"`                   // 本周点单流水
	AdminType                   ClanAdminType     `protobuf:"varint,15,opt,name=admin_type,json=adminType,proto3,enum=clan_pb.ClanAdminType" json:"admin_type,omitempty"`                    // 会员类型
	ChairTimeToday              int32             `protobuf:"varint,16,opt,name=chair_time_today,json=chairTimeToday,proto3" json:"chair_time_today,omitempty"`                              // 今天上麦时长，单位：秒
	ChairTimeWeek               int32             `protobuf:"varint,17,opt,name=chair_time_week,json=chairTimeWeek,proto3" json:"chair_time_week,omitempty"`                                 // 本周上麦时长，单位：秒
	LastWeekReceiveNum          int32             `protobuf:"varint,18,opt,name=last_week_receive_num,json=lastWeekReceiveNum,proto3" json:"last_week_receive_num,omitempty"`                // 上周接待数
	WeekReceiveNum              int32             `protobuf:"varint,19,opt,name=week_receive_num,json=weekReceiveNum,proto3" json:"week_receive_num,omitempty"`                              // 本周接待数
	OnlyFansGoldWeek            int32             `protobuf:"varint,20,opt,name=only_fans_gold_week,json=onlyFansGoldWeek,proto3" json:"only_fans_gold_week,omitempty"`                      // 本周粉丝订阅流水
	PhoneGoldWeek               int32             `protobuf:"varint,21,opt,name=phone_gold_week,json=phoneGoldWeek,proto3" json:"phone_gold_week,omitempty"`                                 // 本周心动8分钟（连麦）流水
	NewUserRate                 int32             `protobuf:"varint,22,opt,name=new_user_rate,json=newUserRate,proto3" json:"new_user_rate,omitempty"`                                       // 新用户回复率
	PlayerRate                  int32             `protobuf:"varint,23,opt,name=player_rate,json=playerRate,proto3" json:"player_rate,omitempty"`                                            // 主播回复率
	AuctionGoldWeek             int32             `protobuf:"varint,24,opt,name=auction_gold_week,json=auctionGoldWeek,proto3" json:"auction_gold_week,omitempty"`                           // 本周拍拍房流水
	AdminTypeCsvMap             map[string]string `protobuf:"bytes,25,rep,name=admin_type_csv_map,json=adminTypeCsvMap,proto3" json:"admin_type_csv_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AdminTypeCsv                string            `protobuf:"bytes,26,opt,name=admin_type_csv,json=adminTypeCsv,proto3" json:"admin_type_csv,omitempty"`
	MemberType                  int32             `protobuf:"varint,27,opt,name=member_type,json=memberType,proto3" json:"member_type,omitempty"`                                                          // 会员类型
	RoomGoldLastWeek            int32             `protobuf:"varint,28,opt,name=room_gold_last_week,json=roomGoldLastWeek,proto3" json:"room_gold_last_week,omitempty"`                                    // 家族房间流水+",
	PeopleGoldLastWeek          int32             `protobuf:"varint,29,opt,name=people_gold_last_week,json=peopleGoldLastWeek,proto3" json:"people_gold_last_week,omitempty"`                              // 个人房间流水+",
	PersonalPatternGoldLastWeek int32             `protobuf:"varint,30,opt,name=personal_pattern_gold_last_week,json=personalPatternGoldLastWeek,proto3" json:"personal_pattern_gold_last_week,omitempty"` // 直播模式流水+",
	ImGoldLastWeek              int32             `protobuf:"varint,31,opt,name=im_gold_last_week,json=imGoldLastWeek,proto3" json:"im_gold_last_week,omitempty"`                                          // IM礼物流水+",
	IntimateGoldLastWeek        int32             `protobuf:"varint,32,opt,name=intimate_gold_last_week,json=intimateGoldLastWeek,proto3" json:"intimate_gold_last_week,omitempty"`                        // 亲密礼物流水+",
	OnlyFansGoldLastWeek        int32             `protobuf:"varint,33,opt,name=only_fans_gold_last_week,json=onlyFansGoldLastWeek,proto3" json:"only_fans_gold_last_week,omitempty"`                      // 粉丝订阅收益+",
	PhoneGoldLastWeek           int32             `protobuf:"varint,34,opt,name=phone_gold_last_week,json=phoneGoldLastWeek,proto3" json:"phone_gold_last_week,omitempty"`                                 // 心动8分钟收益+",
}

func (x *ClanMember) String() string {
	return fmt.Sprintf("%+v", x)
}

func (*ClanMember) ProtoMessage() {}

func (x *ClanMember) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClanMember) GetId2() int64 {
	if x != nil {
		return x.Id2
	}
	return 0
}

func (x *ClanMember) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClanMember) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ClanMember) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *ClanMember) GetRoomGoldWeek() int32 {
	if x != nil {
		return x.RoomGoldWeek
	}
	return 0
}

func (x *ClanMember) GetPeopleGoldWeek() int32 {
	if x != nil {
		return x.PeopleGoldWeek
	}
	return 0
}

func (x *ClanMember) GetTicket() int32 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

func (x *ClanMember) GetBillGoldWeek() int32 {
	if x != nil {
		return x.BillGoldWeek
	}
	return 0
}

func (x *ClanMember) GetPersonalPatternGoldWeek() int32 {
	if x != nil {
		return x.PersonalPatternGoldWeek
	}
	return 0
}

func (x *ClanMember) GetIsBillAdmin() bool {
	if x != nil {
		return x.IsBillAdmin
	}
	return false
}

func (x *ClanMember) GetImGiftGoldWeek() int32 {
	if x != nil {
		return x.ImGiftGoldWeek
	}
	return 0
}

func (x *ClanMember) GetIntimateGoldWeek() int32 {
	if x != nil {
		return x.IntimateGoldWeek
	}
	return 0
}

func (x *ClanMember) GetSongOrderGoldWeek() int32 {
	if x != nil {
		return x.SongOrderGoldWeek
	}
	return 0
}

func (x *ClanMember) GetAdminType() ClanAdminType {
	if x != nil {
		return x.AdminType
	}
	return ClanAdminType_CAT_GUEST
}

func (x *ClanMember) GetChairTimeToday() int32 {
	if x != nil {
		return x.ChairTimeToday
	}
	return 0
}

func (x *ClanMember) GetChairTimeWeek() int32 {
	if x != nil {
		return x.ChairTimeWeek
	}
	return 0
}

func (x *ClanMember) GetLastWeekReceiveNum() int32 {
	if x != nil {
		return x.LastWeekReceiveNum
	}
	return 0
}

func (x *ClanMember) GetWeekReceiveNum() int32 {
	if x != nil {
		return x.WeekReceiveNum
	}
	return 0
}

func (x *ClanMember) GetOnlyFansGoldWeek() int32 {
	if x != nil {
		return x.OnlyFansGoldWeek
	}
	return 0
}

func (x *ClanMember) GetPhoneGoldWeek() int32 {
	if x != nil {
		return x.PhoneGoldWeek
	}
	return 0
}

func (x *ClanMember) GetNewUserRate() int32 {
	if x != nil {
		return x.NewUserRate
	}
	return 0
}

func (x *ClanMember) GetPlayerRate() int32 {
	if x != nil {
		return x.PlayerRate
	}
	return 0
}

func (x *ClanMember) GetAuctionGoldWeek() int32 {
	if x != nil {
		return x.AuctionGoldWeek
	}
	return 0
}

func (x *ClanMember) GetAdminTypeCsvMap() map[string]string {
	if x != nil {
		return x.AdminTypeCsvMap
	}
	return nil
}

func (x *ClanMember) GetAdminTypeCsv() string {
	if x != nil {
		return x.AdminTypeCsv
	}
	return ""
}

func (x *ClanMember) GetMemberType() int32 {
	if x != nil {
		return x.MemberType
	}
	return 0
}

func (x *ClanMember) GetRoomGoldLastWeek() int32 {
	if x != nil {
		return x.RoomGoldLastWeek
	}
	return 0
}

func (x *ClanMember) GetPeopleGoldLastWeek() int32 {
	if x != nil {
		return x.PeopleGoldLastWeek
	}
	return 0
}

func (x *ClanMember) GetPersonalPatternGoldLastWeek() int32 {
	if x != nil {
		return x.PersonalPatternGoldLastWeek
	}
	return 0
}

func (x *ClanMember) GetImGoldLastWeek() int32 {
	if x != nil {
		return x.ImGoldLastWeek
	}
	return 0
}

func (x *ClanMember) GetIntimateGoldLastWeek() int32 {
	if x != nil {
		return x.IntimateGoldLastWeek
	}
	return 0
}

func (x *ClanMember) GetOnlyFansGoldLastWeek() int32 {
	if x != nil {
		return x.OnlyFansGoldLastWeek
	}
	return 0
}

func (x *ClanMember) GetPhoneGoldLastWeek() int32 {
	if x != nil {
		return x.PhoneGoldLastWeek
	}
	return 0
}
