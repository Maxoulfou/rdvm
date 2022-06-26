/*
Copyright © 2022 BROCHIER MAXENCE maxence@brochier.xyz

*/
package execution

import (
	"bytes"
	"fmt"
	"github.com/Maxoulfou/rdvm/config"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

func Execute(action string, json config.JsonConfig) []string {
	url := json.MasterUrl
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("key", json.Key)
	_ = writer.WriteField("hash", json.Hash)
	_ = writer.WriteField("action", action)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)

		return nil
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)

		return nil
	}
	req.Header.Add("key", json.Key)
	req.Header.Add("hash", json.Hash)
	req.Header.Add("action", action)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

		return nil
	}

	//fmt.Println("Body: ", string(body))

	final := strings.Split(string(body), `>`)

	var inter []string
	for i, elt := range final {
		if i%2 == 1 {
			inter2 := strings.Split(elt, `</`)
			if i%2 == 1 {
				inter = append(inter, inter2[0])
			}
		}
	}

	/*for i, elt := range inter {
		fmt.Printf("i: %+v / elt: %+v\n", i, elt)
	}*/

	return inter
}
