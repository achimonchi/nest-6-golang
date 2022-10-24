package httpclient

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type JSONTypicode struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func TestGet(t *testing.T) {
	client := NewHttpClient("https://jsonplaceholder.typicode.com/posts")
	resp, err := client.Get("/")
	assert.Nil(t, err)

	var datas []JSONTypicode

	err = json.Unmarshal(resp, &datas)

	fmt.Println(datas[0])
}
func TestPost(t *testing.T) {
	client := NewHttpClient("https://jsonplaceholder.typicode.com/posts")
	resp, err := client.Post("/", map[string]interface{}{
		"userId": 1,
		"title":  "Percobaan",
		"body":   "ini percobaan",
	})
	assert.Nil(t, err)

	var datas JSONTypicode

	err = json.Unmarshal(resp, &datas)
	fmt.Println(err)
	fmt.Printf("%+v\n", datas)
}
