package response

type K3Response struct {
	Result Result `json:"Result"`
}

type Result struct {
	Status         K3Status         `json:"ResponseStatus"`
	ID             int64            `json:"Id"`
	Number         string           `json:"Number"`
	NeedReturnData []NeedReturnData `json:"NeedReturnData"`
	Result         interface{}      `json:"Result"` //查看接口可用
}

type K3Response2 struct {
	Result Result2 `json:"Result"`
}

type Result2 struct {
	Status         K3Status         `json:"ResponseStatus"`
	ID             string           `json:"Id"`
	Number         string           `json:"Number"`
	NeedReturnData []NeedReturnData `json:"NeedReturnData"`
	Result         interface{}      `json:"Result"` //查看接口可用
}

type NeedReturnData struct {
	FVOUCHERGROUPNO string `json:"FVOUCHERGROUPNO"`
	FVOUCHERID      string `json:"FVOUCHERID"`
}

type K3Status struct {
	ErrorCode int  `json:"ErrorCode"`
	IsSuccess bool `json:"IsSuccess"`
	Errors    []struct {
		FieldName interface{} `json:"FieldName"`
		Message   string      `json:"Message"`
		DIndex    int         `json:"DIndex"`
	} `json:"Errors"`
	SuccessEntitys  []SuccessEntities `json:"SuccessEntitys"`
	SuccessMessages []interface{}     `json:"SuccessMessages"`
	MsgCode         int               `json:"MsgCode"`
}

type SuccessEntities struct {
	ID     int    `json:"Id"`
	Number string `json:"Number"`
	DIndex int    `json:"DIndex"`
}
