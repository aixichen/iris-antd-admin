package models

import (
	"gorm.io/gorm"
)

type Waybill struct {
	gorm.Model
	WaybillCode          string
	WaybillType          int // 0 提车运单 1承运运单 2送车运单
	WaybillStatus        int // 0 待处理 1运输中 2已完成
	WaybillSettleStatus  int //结算状态 0 待结算 1 部分结算 2 已结算
	WaybillSettleType    int // 0到付 1回单结算 2回单月结 3账期N+1 4账期N+2 5账期N+3
	WaybillOrderCarCount int
	WaybillOrderCarCode  string
	WaybillOrderCode     string
	WaybillFreightMoney  float64 //总运费
	WaybillPregetMoney   float64 //代收款合计
	WaybillAmountMoney   float64 //应付合计

	WaybillPlateCarType             int    //0自营运输 2外协运输
	WaybillPlateCarIdCode           string //车牌
	WaybillPlateCarCrop             string //外协公司
	WaybillPlateCarCropPersonName   string //外协公司联系人
	WaybillPlateCarCropPersonMobile string //外协公司联系人电话
	WaybillPlateCarPersonName       string //司机姓名
	WaybillPlateCarPersonMobile     string //司机电话
	WaybillPlateCarPersonOther      string //司机其他信息

	WaybillSettleCrop         string //结算单位
	WaybillSettlePersonName   string //结算人
	WaybillSettlePersonMobile string //结算人联系电话

	WaybillPaymentMoney float64 //支付现金金额
	WaybillRemark       string  //运单备注

	WaybillOrderIsHaveGuardDeliver int //是否有控车车辆

	CreatedUserId         uint
	CreatedUserName       string
	CreatedUserMobile     string
	CreatedUserOfficeId   uint
	CreatedUserOfficeName string
	CreatedCompanyId      uint
	CreatedCompanyName    string
}
