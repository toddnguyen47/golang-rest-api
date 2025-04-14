package restapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type RestApi interface {
	Calc(input int) []byte
}

type Config struct {
	Url string
}

type impl struct {
	Config
}

func NewRestApi(config Config) RestApi {
	return &impl{
		Config: config,
	}
}

func (i1 *impl) Calc(input int) []byte {
	path := fmt.Sprintf("%s/calc", i1.Url)
	baseUrl, err := url.Parse(path)
	if err != nil {
		// TODO: Handle error gracefully
		panic(err)
	}

	params := baseUrl.Query()
	params.Add("input", strconv.Itoa(input))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		// TODO: Handle error gracefully
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		// TODO: Handle error gracefully
		panic(fmt.Errorf("status is NOT ok; status: %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO: Handle error gracefully
		panic(err)
	}

	return body
}
