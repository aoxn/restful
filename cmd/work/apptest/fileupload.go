package apptest

import (
	"fmt"
	"mime/multipart"
	"bytes"
	"os"
	"io"
	"net/http"
	"yy.com/container/common"
)

func Upload() (err error) {

	// Create buffer
	buf := new(bytes.Buffer) // caveat IMO dont use this for large files, \
	// create a tmpfile and assemble your multipart from there (not tested)
	w := multipart.NewWriter(buf)
	// Create file field
	wdir,err:= common.GetWWWPath()
	fmt.Println(wdir)
	w.WriteField("user","user1")
	w.WriteField("sessionid","e6dbd8f6-c1bc-4f26-b28d-771ccb2d17c5")
	w.WriteField("filepath","file")

	fw, err := w.CreateFormFile("file", "build.go") //这里的file很重要，必须和服务器端的FormFile一致
	if err != nil {
		fmt.Println("c")
		return err
	}
	fd, err := os.Open("../../unsafe/unsafe.go")
	if err != nil {
		fmt.Println("d")
		return err
	}
	defer fd.Close()
	// Write file field from file to upload
	_, err = io.Copy(fw, fd)
	if err != nil {
		fmt.Println("e")
		return err
	}
	// Important if you do not close the multipart writer you will not have a
	// terminating boundry
	w.Close()
	req, err := http.NewRequest("POST","http://localhost:8080/container/v1/pkgmaker/pkg/56/files", buf)
	if err != nil {
		fmt.Println("f")
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("g",err)
		return err
	}
	io.Copy(os.Stderr, res.Body) // Replace this with Status.Code check
	fmt.Println("h",res)
	return err
}
