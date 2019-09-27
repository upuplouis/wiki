package render

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

type SeldonTemplateInfo struct {

}

//打印日志
func CheckErrors(err error, message string, logSuccess bool)  {
	if err != nil {
		logrus.Error(fmt.Sprintf("fail: %s, detail: %s", message, err))
	} else if logSuccess {
		logrus.Info(fmt.Sprintf("success: %s", message))
	}
}

//渲染模板
func UpdateTemplate(applicationName string, chartPrefix string)  {
	seldonTemplateInfo := SeldonTemplateInfo{}
	templateFile := fmt.Sprintf("%s/values.template", chartPrefix)
	valueYamlFile := fmt.Sprintf("%s/value.yaml", chartPrefix)
	valueBytes := graphRender(templateFile, seldonTemplateInfo)
	f, _ := os.OpenFile(valueYamlFile, os.O_RDWR|os.O_CREATE, 0666)
	_, err := io.WriteString(f, string(valueBytes.Bytes()))
	CheckErrors(err, "", false)
}

//解析模板
func graphRender(templateFile string, seldonTemplateInfo SeldonTemplateInfo) *bytes.Buffer {
	templateFileByte, err := ioutil.ReadFile(templateFile)
	templateInfo := template.New("")
	_, err = templateInfo.Parse(string(templateFileByte))
	var byteBuf bytes.Buffer
	err = templateInfo.Execute(&byteBuf, seldonTemplateInfo)
	CheckErrors(err, "", false)
	return &byteBuf
}