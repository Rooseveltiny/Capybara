package capybara

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/smartystreets/goconvey/convey"
)

type TestInputData struct {
	FeetQnty  int    `json:"feet_qnty"`
	NameOfDog string `json:"name_of_dog"`
}

type TestInputData2 struct {
	Name    string `json:"name"`
	Sername string `json:"sername"`
}

type TestOutputData struct {
	SpongeColor     string   `json:"sponge_color"`
	ColorsAvailable []string `json:"colors_available"`
}

type TestOutputDataResponse[OutputDataType any] struct {
	Errors     []string       `json:"errors"`
	OutputData OutputDataType `json:"output_data"`
}

var InputFilePath string = "/Users/saveliytrifanenkov/Desktop/Проекты/Capybara/test_data/test_input_data.json"
var InputFilePath2 string = "/Users/saveliytrifanenkov/Desktop/Проекты/Capybara/test_data/test_input_data2.json"
var OutputFilePath string = "/Users/saveliytrifanenkov/Desktop/Проекты/Capybara/test_data/test_output_data.json"
var InputFilePathFlag string = fmt.Sprintf("-input_data_path=%s", InputFilePath)
var InputFilePathFlag2 string = fmt.Sprintf("-input_data_path=%s", InputFilePath2)
var OutputFilePathFlag string = fmt.Sprintf("-output_data_path=%s", OutputFilePath)

func TestCapyLogger(t *testing.T) {
	convey.Convey("test logger", t, func() {
		fi, err := os.Stat("system_logging.log")
		convey.So(err, convey.ShouldBeNil)
		size := fi.Size()
		CapyLogger.Infoln("test 1")
		fi, err = os.Stat("system_logging.log")
		convey.So(err, convey.ShouldBeNil)
		convey.So(size, convey.ShouldBeLessThan, fi.Size())
	})
}

func TestRequst2(t *testing.T) {
	convey.Convey("test req 2", t, func() {
		os.Args = append(os.Args, InputFilePathFlag2)
		os.Args = append(os.Args, OutputFilePathFlag)

		req := NewRequest[TestInputData2]()
		fmt.Println(req)
	})
}

func TestRequest(t *testing.T) {
	convey.Convey("test request init", t, func() {

		os.Args = append(os.Args, InputFilePathFlag)
		os.Args = append(os.Args, OutputFilePathFlag)

		req := NewRequest[TestInputData]()

		var InputDataTemplate TestInputData = TestInputData{
			FeetQnty:  10,
			NameOfDog: "Testy boy",
		}

		userUUID, _ := uuid.FromString("9e98d431-73f1-44b0-bbfd-408876abdb94")
		var User User = User{
			Name: "Test user",
			Link: userUUID,
		}

		convey.Convey("get data paths", func() {
			// req.initDataPaths()
			convey.So(req.InputDataPath, convey.ShouldEqual, InputFilePath)
			convey.So(req.OutputDataPath, convey.ShouldEqual, OutputFilePath)
			convey.Convey("read input data from file", func() {
				// req.readInputData()
				convey.So(InputDataTemplate, convey.ShouldResemble, req.InputData)
				convey.So(User, convey.ShouldResemble, req.User)
			})
		})
	})
}

func TestResponse(t *testing.T) {
	convey.Convey("test response init", t, func() {
		resp := NewResponse[TestOutputData]()
		convey.So(resp, convey.ShouldNotBeNil)
		convey.Convey("test set file path of output data", func() {
			convey.So(resp.OutputDataPath, convey.ShouldBeEmpty)
			resp.SetResponseDataFilePath(OutputFilePath)
			convey.So(resp.OutputDataPath, convey.ShouldEqual, OutputFilePath)
			convey.Convey("test set response data", func() {
				od := TestOutputData{
					SpongeColor:     "red",
					ColorsAvailable: []string{"blue", "greend", "red"},
				}
				resp.SetResponseData(od)
				convey.So(resp.OutputData, convey.ShouldResemble, od)
				convey.Convey("test appending an error into response", func() {
					resp.AppendError(errors.New("error occured! code of an error: 1456"))
					resp.AppendError(errors.New("error occured! code of an error: 8888"))
					resp.AppendError(errors.New("error occured! code of an error: 3264"))
					convey.So(len(resp.Errors), convey.ShouldEqual, 3)
					convey.Convey("test save response method", func() {
						resp.SaveResponse()
						f, _ := os.Open(OutputFilePath)
						defer f.Close()
						b, _ := io.ReadAll(f)
						output_data := &TestOutputDataResponse[TestOutputData]{}
						json.Unmarshal(b, output_data)
						convey.So(output_data.OutputData.SpongeColor, convey.ShouldEqual, "red")
					})
				})
			})
		})
	})
}

func TestCapybaraSession(t *testing.T) {

	os.Args = append(os.Args, InputFilePathFlag)
	os.Args = append(os.Args, OutputFilePathFlag)

	convey.Convey("test capybara session object", t, func() {
		capybara := NewCapybaraSession[TestInputData, TestOutputData]()
		convey.So(capybara.GetInputData().FeetQnty, convey.ShouldEqual, 10)
		convey.So(capybara.Request.InputDataPath, convey.ShouldEqual, InputFilePath)
		convey.So(capybara.Response.OutputDataPath, convey.ShouldEqual, OutputFilePath)
		convey.Convey("test setting output data", func() {
			od := TestOutputData{
				SpongeColor:     "blue",
				ColorsAvailable: []string{"orange", "brown", "violet"},
			}
			capybara.SetOutputData(od)
			convey.So(capybara.Response.OutputData, convey.ShouldResemble, od)
			capybara.AppendError(errors.New("new error occured"))
			convey.So(len(capybara.Response.Errors), convey.ShouldEqual, 1)
			convey.Convey("test saving response", func() {
				capybara.SaveResponse()
				f, _ := os.Open(OutputFilePath)
				defer f.Close()
				b, _ := io.ReadAll(f)
				output_data := &TestOutputDataResponse[TestOutputData]{}
				json.Unmarshal(b, output_data)
				convey.So(output_data.OutputData.SpongeColor, convey.ShouldEqual, "blue")
			})
		})
	})
}
