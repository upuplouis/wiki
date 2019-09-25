package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

func FileUtil()  {
	f, err := os.OpenFile("", os.O_CREATE|os.O_RDWR, os.ModePerm)
	f, _ = os.Open("")
	os.IsNotExist(err)
	_, err = f.Write([]byte{})
	err = os.Setenv("", "")
	os.Getenv("")
	err = os.Mkdir("", os.ModePerm)
	err = os.Remove("")
	_, err = ioutil.ReadFile("")
	_, err = ioutil.ReadAll(bytes.NewReader([]byte{}))
	_, err = io.Copy(f, f)
	_, err = io.WriteString(f, "")
}
