package capybara

import (
	"encoding/json"
	"os"
)

type Response[OutputDataType interface{}] struct {
	OutputDataPath string         `json:"-"`
	OutputData     OutputDataType `json:"output_data"`
	Errors         []error        `json:"-"`
	ErrorsString   []string       `json:"errors"`
}

func (r *Response[OutputDataType]) SaveResponse() {

	r.ErrorsString = make([]string, 0)
	for _, v := range r.Errors {
		r.ErrorsString = append(r.ErrorsString, v.Error())
	}

	JsonB, _ := json.Marshal(r)

	err := os.WriteFile(r.OutputDataPath, []byte(JsonB), 0755)
	if err != nil {
		CapyLogger.Infoln(r.OutputDataPath)
		CapyLogger.Errorln(err)
	}

}

func (r *Response[OutputDataType]) SetResponseDataFilePath(filePath string) {
	r.OutputDataPath = filePath
}

func (r *Response[OutputDataType]) SetResponseData(data OutputDataType) {
	r.OutputData = data
}

func (r *Response[OutputDataType]) AppendError(err error) {
	CapyLogger.Errorln(err)
	r.Errors = append(r.Errors, err)
}

func NewResponse[OutputData interface{}]() *Response[OutputData] {
	return &Response[OutputData]{}
}
