package models

import (
	"gorm.io/gorm"
	"time"
)

type Finance struct {
	gorm.Model
	FinanceSourceId        uint
	FinanceSourceType      int // 1 订单 2运单
	FinanceSourceCode      string
	FinanceSourceIsSign    int // 0 未签收 1已签收
	FinanceType            int // 1 订单运费 2订单返款 3提车费 4承运费 5代收款 6送车费 7 其他费用
	FinanceTypeCn          string
	FinancePayMoney        float64 //应付金额
	FinanceReceivableMoney float64 //应收金额

	FinanceSettleCrop         string //结算单位
	FinanceSettlePersonName   string //结算人
	FinanceSettlePersonMobile string //结算人联系电话
	FinanceSettleStatus       int    //结算状态 0 待结算 1 部分结算 2 已结算
	FinanceSettleType         int    // 0到付 1回单结算 2回单月结 3账期N+1 4账期N+2 5账期N+3
	FinanceRemark             string //备注
	OrderId                   uint
	OrderCarType              string // 商品车、二手车、私家车，巡展车，事故车
	OrderCarName              string
	OrderCarCode              string  //车架号
	OrderCarWorth             float64 //车辆估值

	OrderStartProvinceName string
	OrderStartProvinceCode string
	OrderStartCityName     string
	OrderStartCityCode     string
	OrderStartAreaName     string
	OrderStartAreaCode     string

	OrderArriveProvinceName string
	OrderArriveProvinceCode string
	OrderArriveCityName     string
	OrderArriveityCode      string
	OrderArriveAreaName     string
	OrderArriveAreaCode     string

	OrderIsGuardDeliver int //是否控车

	FinancePaymentMoney    float64 //已付金额
	FinancePaymentUserId   int
	FinancePaymentUserName string
	FinancePaymentTime     time.Time
	FinanceIsUpdate        int //0 未被财务修改 1被财务修改
	CreatedUserId          uint
	CreatedUserName        string
	CreatedUserMobile      string
	CreatedUserOfficeId    uint
	CreatedUserOfficeName  string
	CreatedCompanyId       uint
	CreatedCompanyName     string
}
