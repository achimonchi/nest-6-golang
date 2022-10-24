package adaptor

import (
	"encoding/json"
	"sesi4/pkg/httpclient"
)

type TypicodeAdaptor struct {
	client *httpclient.Client
}

type JSONTypicode struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func NewTypicodeAdaptor(baseUrl string) *TypicodeAdaptor {
	client := httpclient.NewHttpClient(baseUrl)

	return &TypicodeAdaptor{
		client: client,
	}
}

func (t *TypicodeAdaptor) GetAllTypicode() (*[]JSONTypicode, error) {
	data, err := t.client.Get("/")
	if err != nil {
		return nil, err
	}
	var datas []JSONTypicode

	err = json.Unmarshal(data, &datas)
	if err != nil {
		return nil, err
	}

	return &datas, err

}
func (t *TypicodeAdaptor) CreateTypicode() {
	t.client.Post("/", nil)
}
