package common

type successRres struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRres {
	return &successRres{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccesResponse(data interface{}) *successRres {
	return NewSuccessResponse(data, nil, nil)
}
