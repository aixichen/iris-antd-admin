package validates

import "time"

type (
	OrderRequest struct {
		Id               uint    `json:"id" comment:"订单ID"`
		OrderCode        string  `json:"order_code" comment:"订单编号"`
		OrderReceiveType int     `json:"order_receive_type" comment:"提车方式"`
		OrderDeliverType int     `json:"order_deliver_type" comment:"交车方式"`
		OrderOfficeId    uint    `json:"order_office_id" validate:"required" comment:"当前处理网点"`
		OrderOfficeName  string  `json:"order_office_name" validate:"required" comment:"当前处理网点"`
		OrderSettleType  int     `json:"order_settle_type" validate:"required" comment:"结算方式"`
		OrderCarType     string  `json:"order_car_type" validate:"required" comment:"车辆类型"`
		OrderCarName     string  `json:"order_car_name" validate:"required" comment:"车辆品牌"`
		OrderCarCode     string  `json:"order_car_code" comment:"车架号"`
		OrderCarWorth    float64 `json:"order_car_worth" comment:"车辆估值"`

		OrderClientCorp         string `json:"order_client_corp" comment:"委托单位"`
		OrderClientPersonName   string `json:"order_client_person_name" validate:"required" comment:"委托人"`
		OrderClientPersonMobile string `json:"order_client_person_mobile" validate:"required，gte=11,lte=11" comment:"委托联系方式"`

		OrderArriveCorp         string `json:"order_arrive_corp" comment:"收货单位"`
		OrderArrivePersonName   string `json:"order_arrive_person_name" validate:"required" comment:"收货人"`
		OrderArrivePersonMobile string `json:"order_arrive_person_mobile" validate:"required，gte=11,lte=11" comment:"收货联系方式"`

		OrderSettleCorp         string `json:"order_settle_corp" comment:"结算单位"`
		OrderSettlePersonName   string `json:"order_settle_person_name" validate:"required" comment:"结算人"`
		OrderSettlePersonMobile string `json:"order_settle_person_mobile" validate:"required，gte=11,lte=11" comment:"结算联系方式"`

		OrderStartProvinceName string `json:"order_start_province_name" validate:"required" comment:"起始地:省"`
		OrderStartProvinceCode string `json:"order_start_province_code" validate:"required" comment:"起始地:省CODE"`
		OrderStartCityName     string `json:"order_start_city_name" validate:"required" comment:"起始地:市"`
		OrderStartCityCode     string `json:"order_start_city_code" validate:"required" comment:"起始地:市CODE`
		OrderStartAreaName     string `json:"order_start_area_name" comment:"起始地:区"`
		OrderStartAreaCode     string `json:"order_start_area_code" comment:"起始地:区CODE"`
		OrderStartAddress      string `json:"order_start_address" comment:"起始地:详细地址"`

		OrderArriveProvinceName string `json:"order_arrive_province_name" validate:"required" comment:"目的地:省"`
		OrderArriveProvinceCode string `json:"order_arrive_province_code" validate:"required" comment:"目的地:省CODE"`
		OrderArriveCityName     string `json:"order_arrive_city_name" validate:"required" comment:"目的地:市"`
		OrderArriveCityCode     string `json:"order_arrive_city_code" validate:"required" comment:"目的地:市CODE`
		OrderArriveAreaName     string `json:"order_arrive_area_name" comment:"目的地:区"`
		OrderArriveAreaCode     string `json:"order_arrive_area_code" comment:"目的地:区CODE"`
		OrderArriveAddress      string `json:"order_arrive_address" comment:"目的地:详细地址"`

		//当前位置
		OrderCurrentProvinceName string    `json:"order_current_province_name" comment:"当前位置:省"`
		OrderCurrentProvinceCode string    `json:"order_current_province_code" comment:"当前位置:省CODE"`
		OrderCurrentCityName     string    `json:"order_current_city_name" comment:"当前位置:市"`
		OrderCurrentCityCode     string    `json:"order_current_city_code" comment:"当前位置:市CODE`
		OrderCurrentAreaName     string    `json:"order_current_area_name" comment:"当前位置:区"`
		OrderCurrentAreaCode     string    `json:"order_current_area_code" comment:"当前位置:区CODE"`
		OrderCurrentTime         time.Time `json:"order_current_time" comment:"当前位置:时间"` //到达当前位置时间

		OrderFreightMoney float64 `json:"order_freight_money" comment:"总运费"` //总运费
		OrderRefundMoney  float64 `json:"order_refund_money" comment:"返款"`   //返款
		//返款结算人
		OrderRefundCorp         string  `json:"order_refund_corp" comment:"返款结算公司"`
		OrderRefundPersonName   string  `json:"order_refund_person_name" comment:"反馈结算人"`
		OrderRefundPersonMobile string  `json:"order_refund_person_mobile" comment:"返款结算人联系方式"`
		OrderAmountMoney        float64 `json:"order_amount_money" comment:"合计金额"` //应收合计
		OrderIsInsure           int     `json:"order_is_insure"`                   //是否包含保险  1是 2 否
		OrderIsInvoice          int     `json:"order_is_invoice"`                  //是否开票 1是 2 否
		OrderInvoice            string  `json:"order_invoice"`                     //开票信息
		OrderRemark             string  `json:"order_remark"`                      //订单备注

		OrderIsGuardDeliver int `json:"order_is_guard_deliver"` //是否控车

	}
)
