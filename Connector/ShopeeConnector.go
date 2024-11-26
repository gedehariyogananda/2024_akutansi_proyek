package Connector

import (
	"2024_akutansi_project/Models/Dto"
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type (
	IShopeeConnector interface {
		SetPushNotification(ctx context.Context, request *Dto.SetPushNotificationShopeeRequest) error
	}

	ShopeeConnector struct {
	}
)

func ShopeeConnectorProvider() *ShopeeConnector {
	return &ShopeeConnector{}
}

// TODO :: adjust the integrate shopee
func (c *ShopeeConnector) SetPushNotification(ctx context.Context, request *Dto.SetPushNotificationShopeeRequest) error {
	params := url.Values{}

	partnerId := request.PartnerID
	params.Set("partner_id", partnerId)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params.Set("timestamp", timestamp)

	path := "/api/v2/push/set_app_push_config"

	partnerKey := request.PartnerKey
	baseString := fmt.Sprintf("%s%s%s", partnerId, path, timestamp)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))
	params.Set("sign", sign)

	redirectUrl := "https://www.baidu.com/"
	params.Set("redirect", redirectUrl)

	urlParam := params.Encode()

	shopeeBaseURL := os.Getenv("SHOPEE_BASE_URL")
	endpoint := fmt.Sprintf("%s%s?%s", shopeeBaseURL, path, urlParam)

	input, err := json.Marshal(struct {
		BlockedShopIdList []int  `json:"blocked_shop_id_list"`
		CallbackUrl       string `json:"callback_url"`
		SetPushConfigOff  []int  `json:"set_push_config_off"`
		SetPushConfigOn   []int  `json:"set_push_config_on"`
	}{
		BlockedShopIdList: request.BlockedShopIdList,
		CallbackUrl:       request.CallbackUrl,
		SetPushConfigOff:  request.SetPushConfigOff,
		SetPushConfigOn:   request.SetPushConfigOn,
	})
	if err != nil {
		return err
	}

	body := bytes.NewReader(input)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Re-usable response body for logging
	rawBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close() // must immediately close
	resp.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("set push notification to shopee got %v, response : %v", resp.StatusCode, string(rawBody))
		fmt.Println(fmt.Sprintf("set push notif got err = %v", err))
		return err
	}

	fmt.Println(fmt.Sprintf("set push notification to shopee got %v, response : %v", resp.StatusCode, string(rawBody)))

	return nil
}
