package tonapi

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type SDK struct {
	BaseURL string // https://toncenter.com/api/v2/
}

func NewSDK(baseURl string) SDK {
	return SDK{BaseURL: baseURl}
}

func (s *SDK) GetAddressInformation(address string) (result AddressInformationResponse, err error) {
	client := resty.New()

	resp, err := client.R().SetQueryParam("address", address).Get(
		fmt.Sprintf("%s%s", s.BaseURL, getAddressInformation),
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return result, err
	}
	if result.Error != "" {
		return result, fmt.Errorf(result.Error)
	}
	return
}

func (s *SDK) GetAddressBalance(address string) (result float64, err error) {
	client := resty.New()

	resp, err := client.R().SetQueryParam("address", address).Get(
		fmt.Sprintf("%s%s", s.BaseURL, getAddressBalance),
	)
	if err != nil {
		return result, err
	}
	if v := gjson.GetBytes(resp.Body(), "error").String(); v != "" {
		return result, fmt.Errorf(v)
	}

	return gjson.GetBytes(resp.Body(), "result").Float(), nil
}
