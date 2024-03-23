package codex

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestErrToCode(t *testing.T) {
	url := "https://xy.people.cn/rbbj_api/api/gzarticle_pass_generate"
	method := "GET"

	payload := strings.NewReader(`{"name":"公司","kind":2,"phone":"13369389074","avatar":"1","address":"","domain":"11","status":1}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", " application/json, text/plain, */*")
	req.Header.Add("Accept-Encoding", " gzip, deflate, br")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9")
	req.Header.Add("Authorization", " Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiNjliMDI1MmYxYjI1M2Q5NDI3ODRjM2JjYzVjM2E4YTcyZmQ3NWM3MGY5NWMwN2IzZWU2ZTA3ZTNhODU5Y2Q5ZjA3ZmZkNDUzZjY0MDUxZTkiLCJpYXQiOjE3MDEyMzgxMzkuOTI3ODIzLCJuYmYiOjE3MDEyMzgxMzkuOTI3ODI3LCJleHAiOjE3MDEzMjQ1MzkuODYyNDE2LCJzdWIiOiI1ODEwNyIsInNjb3BlcyI6W119.X3r9_O3KZ6van_dXGIx4UMC3ksftxPJabV30t8N41aE5SKMutjBNOM4HdC1EWvcdpItNcLP7NmVbohP-9JxWLA")
	req.Header.Add("Connection", " keep-alive")
	req.Header.Add("Cookie", " __jsluid_s=a37796c7bb12e88ea971342ad9a3bb5e; x-clockwork=%7B%22requestId%22%3A%221701238142-2452-1826778106%22%2C%22version%22%3A%225.1.12%22%2C%22path%22%3A%22%5C%2F__clockwork%5C%2F%22%2C%22webPath%22%3A%22%5C%2Fclockwork%5C%2Fapp%22%2C%22token%22%3A%22f2b7a5fd%22%2C%22metrics%22%3Atrue%2C%22toolbar%22%3Atrue%7D; __jsluid_s=bb82ce6c24945475b73a191274df5c83")
	req.Header.Add("Host", " xy.people.cn")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
