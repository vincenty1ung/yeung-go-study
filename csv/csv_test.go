package csv

import (
	"testing"
	"time"

	"github.com/vincenty1ung/yeung-go-study/csv/clan_pb"
	"github.com/vincenty1ung/yeung-go-study/csv/i18n"
)

func getMembership(code string, types clan_pb.ClanAdminType) string {
	switch types {
	case clan_pb.ClanAdminType_CAT_GUEST:
		return ""
	case clan_pb.ClanAdminType_CAT_NORMAL:
		return normalMap[code]
	case clan_pb.ClanAdminType_CAT_ANCHOR:
		return specialMap[code]
	case clan_pb.ClanAdminType_CAT_ADMIN:
		return adminMap[code]
	case clan_pb.ClanAdminType_CAT_HIGH_ADMIN:
		return bigAdminMap[code]
	case clan_pb.ClanAdminType_CAT_OWNER:
		return bigMasterMap[code]
	case clan_pb.ClanAdminType_CAT_BIG_FAMILY:
		return bigMasterMap[code]
	}
	return ""
}

// go test -bench=. -benchmem

func BenchmarkNew1(b *testing.B) {
	members := make([]*ClanMemberExtend, 0)
	c1 := &clan_pb.ClanMember{
		ChairTimeToday: 3600,
	}
	duration := time.Duration(c1.ChairTimeToday) * time.Second
	members = append(
		members, &ClanMemberExtend{
			Id:               1453272,
			ImGoldLastWeek:   99,
			Name:             "yangbo",
			AdminType:        clan_pb.ClanAdminType_CAT_NORMAL,
			RoomGoldLastWeek: 10,
			ChairTimeToday:   3600,
			ChairTimeWeekStr: duration.String(),
		},
	)
	/*members = append(
		members, &clan_pb.ClanMember{
			Id:               1453271,
			ImGoldLastWeek:   20,
			Name:             "wanc",
			AdminType:        clan_pb.ClanAdminType_CAT_OWNER,
			RoomGoldLastWeek: 18,
		},
	)*/

	// WriterCSV(mapHeader[i18n.DefaultLanguageCode], members, true)

	// key, key1 := "ImGoldLastWeek", "IMGoldLastWeek"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WriterCSV(mapHeaderV2[i18n.LanguageCodeZHHans], members, true)
	}
}

func BenchmarkNew2(b *testing.B) {
	var csva ClanMemberExtends
	members := make([]*ClanMemberExtend, 0)
	c1 := &clan_pb.ClanMember{
		ChairTimeToday: 3600,
	}
	duration := time.Duration(c1.ChairTimeToday) * time.Second
	members = append(
		members, &ClanMemberExtend{
			Id:               1453272,
			ImGoldLastWeek:   99,
			Name:             "yangbo",
			AdminType:        clan_pb.ClanAdminType_CAT_NORMAL,
			RoomGoldLastWeek: 10,
			ChairTimeToday:   3600,
			ChairTimeWeekStr: duration.String(),
		},
	)
	csva = members
	/*members = append(
		members, &clan_pb.ClanMember{
			Id:               1453271,
			ImGoldLastWeek:   20,
			Name:             "wanc",
			AdminType:        clan_pb.ClanAdminType_CAT_OWNER,
			RoomGoldLastWeek: 18,
		},
	)*/

	// WriterCSV(mapHeader[i18n.DefaultLanguageCode], members, true)

	// key, key1 := "ImGoldLastWeek", "IMGoldLastWeek"
	strings := mapHeaderV2[i18n.LanguageCodeZHHans]
	keys := make([]string, 0)
	head := make([]string, 0)
	for i := 1; i < len(strings); i = i + 2 {
		keys = append(keys, strings[i-1])
		head = append(head, strings[i])
	}
	b.ResetTimer()
	// generate :=
	for i := 0; i < b.N; i++ {
		WriterCSVV2(head, keys, csva.Generate(keys), true)
	}
}
func BenchmarkNew3(b *testing.B) {
	var csva ClanMemberExtends
	members := make([]*ClanMemberExtend, 0)
	c1 := &clan_pb.ClanMember{
		ChairTimeToday: 3600,
	}
	duration := time.Duration(c1.ChairTimeToday) * time.Second
	members = append(
		members, &ClanMemberExtend{
			Id:               1453272,
			ImGoldLastWeek:   99,
			Name:             "yangbo",
			AdminType:        clan_pb.ClanAdminType_CAT_NORMAL,
			RoomGoldLastWeek: 10,
			ChairTimeToday:   3600,
			ChairTimeWeekStr: duration.String(),
		},
	)
	csva = members
	/*members = append(
		members, &clan_pb.ClanMember{
			Id:               1453271,
			ImGoldLastWeek:   20,
			Name:             "wanc",
			AdminType:        clan_pb.ClanAdminType_CAT_OWNER,
			RoomGoldLastWeek: 18,
		},
	)*/

	// WriterCSV(mapHeader[i18n.DefaultLanguageCode], members, true)

	// key, key1 := "ImGoldLastWeek", "IMGoldLastWeek"
	strings := mapHeaderV2[i18n.LanguageCodeZHHans]
	keys := make([]string, 0)
	head := make([]string, 0)
	for i := 1; i < len(strings); i = i + 2 {
		keys = append(keys, strings[i-1])
		head = append(head, strings[i])
	}
	b.ResetTimer()
	// generate :=
	for i := 0; i < b.N; i++ {
		writer := csva.GenerateV2(mapHeader[i18n.LanguageCodeZHHans])
		_ = writer
	}
}

func TestNew2(t *testing.T) {
	var csva ClanMemberExtends
	members := make([]*ClanMemberExtend, 0)
	c1 := &clan_pb.ClanMember{
		ChairTimeToday: 3600,
	}
	duration := time.Duration(c1.ChairTimeToday) * time.Second
	members = append(
		members, &ClanMemberExtend{
			Id:               1453272,
			ImGoldLastWeek:   99,
			Name:             "yangbo",
			AdminType:        clan_pb.ClanAdminType_CAT_NORMAL,
			RoomGoldLastWeek: 10,
			ChairTimeToday:   3600,
			ChairTimeWeekStr: duration.String(),
		},
	)
	csva = members
	/*members = append(
		members, &clan_pb.ClanMember{
			Id:               1453271,
			ImGoldLastWeek:   20,
			Name:             "wanc",
			AdminType:        clan_pb.ClanAdminType_CAT_OWNER,
			RoomGoldLastWeek: 18,
		},
	)*/

	// WriterCSV(mapHeader[i18n.DefaultLanguageCode], members, true)

	// key, key1 := "ImGoldLastWeek", "IMGoldLastWeek"
	strings := mapHeaderV2[i18n.LanguageCodeZHHans]
	keys := make([]string, 0)
	head := make([]string, 0)
	for i := 1; i < len(strings); i = i + 2 {
		keys = append(keys, strings[i-1])
		head = append(head, strings[i])
	}
	// b.ResetTimer()
	// generate :=
	WriterCSVV2(head, keys, csva.Generate(keys), true)
}
