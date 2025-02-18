package mysql

// QueryAllCommunity tip : 查询所有的社区名称
func QueryAllCommunity() ([]string, error) {
	sqlStr := "select community_name from community"
	var communityNameList []string
	if err := db.Select(&communityNameList, sqlStr); err != nil {
		return communityNameList, err
	}
	return communityNameList, nil
}

// QueryCommunityDetailByCommunityID tip : 根据社区ID查询社区详情
func QueryCommunityDetailByCommunityID(oID int) (string, error) {
	sqlStr := "select introduction from community where community_id = ?"
	var detail string
	if err := db.Get(&detail, sqlStr, oID); err != nil {
		return detail, err
	}
	return detail, nil
}
