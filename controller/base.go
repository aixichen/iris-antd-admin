package controller

type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	ShowType     int         `json:"showType"` // error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
	TraceId      string      `json:"traceId"`
}

type Pagination struct {
}

type List struct {
	Data     interface{} `json:"data"`
	Success  bool        `json:"success"`
	Current  int         `json:"current"`  //当前页
	PageSize int         `json:"pageSize"` //每页多少条
	Total    int64       `json:"total"`    //数据总条数
}

func ApiResource(success bool, data interface{}, errorCode string, errorMessage string, ShowType int, traceId string) (r *Response) {
	r = &Response{Success: success, Data: data, ErrorCode: errorCode, ErrorMessage: errorMessage, ShowType: ShowType, TraceId: traceId}
	return
}
