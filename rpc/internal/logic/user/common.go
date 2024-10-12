package user

import "github.com/Sanagiig/fox-admin-core/rpc/ent"

func GetRoleIds(data []*ent.Role) []uint64 {
	var ids []uint64
	for _, v := range data {
		ids = append(ids, v.ID)
	}
	return ids
}

func GetRoleNames(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Name)
	}
	return codes
}

func GetRoleCodes(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Code)
	}
	return codes
}

func GetPositionIds(data []*ent.Position) []uint64 {
	var ids []uint64
	for _, v := range data {
		ids = append(ids, v.ID)
	}
	return ids
}
