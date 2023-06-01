package capybara

type CapybaraSession[InputDataType, OutputDataType any] struct {
	Request  *Request[InputDataType]
	Response *Response[OutputDataType]
}

func (cs *CapybaraSession[InputDataType, OutputDataType]) AppendError(err error) {
	cs.Response.AppendError(err)
}

func (cs *CapybaraSession[InputDataType, OutputDataType]) SetOutputData(data OutputDataType) {
	cs.Response.SetResponseData(data)
}

func (cs *CapybaraSession[InputDataType, OutputDataType]) GetInputData() InputDataType {
	return cs.Request.InputData
}

func (cs *CapybaraSession[InputDataType, OutputDataType]) SaveResponse() {
	cs.Response.SaveResponse()
}

func NewCapybaraSession[InputDataType, OutputDataType any]() *CapybaraSession[InputDataType, OutputDataType] {
	req := NewRequest[InputDataType]()
	resp := NewResponse[OutputDataType]()
	resp.OutputDataPath = req.OutputDataPath
	return &CapybaraSession[InputDataType, OutputDataType]{
		Request:  req,
		Response: resp,
	}
}
