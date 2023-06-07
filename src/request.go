package capybara

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/gofrs/uuid"
)

type User struct {
	Name string    `json:"name"`
	Link uuid.UUID `json:"link"`
}

type Request[InputDataType any] struct {
	InputDataPath  string
	OutputDataPath string
	InputData      InputDataType `json:"input_data"`
	User           User          `json:"user"`
}

func (r *Request[InputDataType]) readInputData() {
	b, err := os.ReadFile(r.InputDataPath)
	if err != nil {
		CapyLogger.Errorln(err)
		return
	}
	errJson := json.Unmarshal(b, &r)
	if errJson != nil {
		CapyLogger.Errorln(errJson)
	}
}

func (r *Request[InputDataType]) initDataPaths() {
	flag.StringVar(&r.InputDataPath, "input_data_path", "", "path of file within intput data")
	flag.StringVar(&r.OutputDataPath, "output_data_path", "", "path of file within output data")
	flag.Parse()
}

func NewRequest[InputDataType any]() *Request[InputDataType] {
	req := &Request[InputDataType]{}
	req.initDataPaths()
	req.readInputData()
	return req
}
