package logic

import "bluebell/dao/mysql"

// GetCommunityList tip : 查询所有的社区名称
func GetCommunityList() ([]string, error) {
	communityNameList, err := mysql.QueryAllCommunity()
	if err != nil {
		return communityNameList, err
	}
	return communityNameList, err
}

// GetCommunityDetail tip : 根据社区ID获取社区详情
func GetCommunityDetail(oId int) (string, error) {
	// 调用DAO层查询详情
	detail, err := mysql.QueryCommunityDetailByCommunityID(oId)
	if err != nil {
		return detail, err
	}
	return detail, err
}
