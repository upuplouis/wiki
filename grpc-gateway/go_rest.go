package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"protos"
	"testing"
)

func TestGo(t *testing.T)  {
	test := make(map[string]interface{})
	byteData, _ := json.Marshal(test)
	reader := bytes.NewReader(byteData)
	client := &http.Client{}
	request, _  := http.NewRequest("POST", "", reader)
	res, _ := client.Do(request)
	respose, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(respose, "")
	_ = res.Body.Close()

	result := &protos.Null{}
	f, _ := ioutil.ReadFile("")
	_ = json.Unmarshal(f, result)
	_ = proto.UnmarshalText(string(f), result)
}

func TestProto(t *testing.T)  {
	result := &protos.Request{}
	p := result.GetPparam().P
	m := result.GetMparam().M
	result.Params = &protos.Request_Pparam{
		Pparam: result.GetPparam(),
	}
	result.Params = &protos.Request_Mparam{
		Mparam: result.GetMparam(),
	}
	fmt.Println(p, m)
}
