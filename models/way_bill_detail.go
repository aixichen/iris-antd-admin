package models

import (
	"gorm.io/gorm"
)

type WaybillDetail struct {
	gorm.Model
	WaybillDetailType int // 1 承运 2 中转 3直送客户 4 中途装车
	WaybillId         uint
	WaybillCode       string
	WaybillType       int // 0 提车运单 1承运运单 2送车运单

	WaybillDetailFreightMoney float64 //运费
	WaybillDetailPregetMoney  float64 //代收款
	WaybillSettleCrop         string  //结算单位
	WaybillSettlePersonName   string  //结算人
	WaybillSettlePersonMobile string  //结算人联系电话
	WaybillSettleStatus       int     //结算状态 0 待结算 1 部分结算 2 已结算
	WaybillSettleType         int     // 0到付 1回单结算 2回单月结 3账期N+1 4账期N+2 5账期N+3
	WaybillDetailRemark       string  //备注

	WaybillPaymentMoney     float64 //支付现金金额
	OrderId                 string
	OrderCarType            string // 商品车、二手车、私家车，巡展车，事故车
	OrderCarName            string
	OrderCarCode            string  //车架号
	OrderCarWorth           float64 //车辆估值
	OrderClientCorp         string
	OrderClientPersonName   string
	OrderClientPersonMobile string

	OrderArriveCorp         string
	OrderArrivePersonName   string
	OrderArrivePersonMobile string

	//起始位置
	WaybillDetailStartProvinceName string
	WaybillDetailStartProvinceCode string
	WaybillDetailStartCityName     string
	WaybillDetailStartCityCode     string
	WaybillDetailStartAreaName     string
	WaybillDetailStartAreaCode     string
	WaybillDetailStartAddress      string

	WaybillDetailArriveProvinceName string
	WaybillDetailArriveProvinceCode string
	WaybillDetailArriveCityName     string
	WaybillDetailArriveityCode      string
	WaybillDetailArriveAreaName     string
	WaybillDetailArriveAreaCode     string
	WaybillDetailArriveAddress      string

	OrderIsGuardDeliver int //是否控车

	WaybillDetailIsSign         int //0 未签收,1 已签收
	WaybillDetailSignType       int //0 转发运,1 转交车
	WaybillDetailSignUserId     string
	WaybillDetailSignUserName   string
	WaybillDetailSignOfficeId   string
	WaybillDetailSignOfficeName string
	WaybillDetailSignRemark     string

	CreatedUserId         uint
	CreatedUserName       string
	CreatedUserMobile     string
	CreatedUserOfficeId   uint
	CreatedUserOfficeName string
	CreatedCompanyId      uint
	CreatedCompanyName    string
}
