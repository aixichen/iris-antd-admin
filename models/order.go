package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OrderCode         string
	OrderReceiveType  int    //提车方式 1 自送网点 2 上门提车
	OrderDeliverType  int    //交车方式 1 网点自提 2 送车上门
	OrderOfficeId     uint   //当前处理网点
	OrderOfficeName   string //当前处理网点名字
	OrderStatus       int    // 1处理中 10运输中 20 待交车 30 已完成
	OrderSettleStatus int    //结算状态 1 待结算 2 部分结算 3 已结算
	OrderSettleType   int    // 1到付 2回单结算 3回单月结 3账期N+1 4账期N+2 5账期N+3
	OrderCarType      string // 商品车、二手车、私家车，巡展车，事故车
	OrderCarName      string
	OrderCarCode      string  //车架号
	OrderCarWorth     float64 //车辆估值

	OrderClientCorp         string
	OrderClientPersonName   string
	OrderClientPersonMobile string

	OrderArriveCorp         string
	OrderArrivePersonName   string
	OrderArrivePersonMobile string

	OrderSettleCorp         string
	OrderSettlePersonName   string
	OrderSettlePersonMobile string

	OrderStartProvinceName string
	OrderStartProvinceCode string
	OrderStartCityName     string
	OrderStartCityCode     string
	OrderStartAreaName     string
	OrderStartAreaCode     string
	OrderStartAddress      string

	OrderArriveProvinceName string
	OrderArriveProvinceCode string
	OrderArriveCityName     string
	OrderArriveCityCode     string
	OrderArriveAreaName     string
	OrderArriveAreaCode     string
	OrderArriveAddress      string

	//当前位置
	OrderCurrentProvinceName string
	OrderCurrentProvinceCode string
	OrderCurrentCityName     string
	OrderCurrentCityCode     string
	OrderCurrentAreaName     string
	OrderCurrentAreaCode     string
	OrderCurrentTime         time.Time //到达当前位置时间

	OrderFreightMoney float64 //总运费
	OrderRefundMoney  float64 //返款
	//返款结算人
	OrderRefundCorp         string
	OrderRefundPersonName   string
	OrderRefundPersonMobile string
	OrderAmountMoney        float64 //应收合计
	OrderIsInsure           int     //是否包含保险  1是 0 否
	OrderIsInvoice          int     //是否开票 1是 0 否
	OrderInvoice            string  //开票信息
	OrderRemark             string  //订单备注

	OrderIsGuardDeliver int //是否控车

	CreatedUserId         uint
	CreatedUserName       string
	CreatedUserMobile     string
	CreatedUserOfficeId   uint
	CreatedUserOfficeName string
	CreatedCompanyId      uint
	CreatedCompanyName    string
}
