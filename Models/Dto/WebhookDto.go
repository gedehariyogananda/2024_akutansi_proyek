package Dto

import (
	"2024_akutansi_project/Helper"
	"encoding/json"
)

type IntegratedProfile struct {
	UserID      string  `json:"user_id" bson:"user_id"`
	ShopeeToken *string `json:"shopee_token,omitempty" bson:"shopee_token,omitempty"`
	TiktokToken *string `json:"tiktok_token,omitempty" bson:"tiktok_token,omitempty"`
}

type ShopeeIntegrateRequest struct {
}

type TiktokIntegrateRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (i *IntegratedProfile) EncryptTokenData() {
	if i.ShopeeToken != nil {
		byteShopeeToken, _ := json.Marshal(i.ShopeeToken)
		encryptedData, _ := Helper.EncryptData(byteShopeeToken)
		i.ShopeeToken = &encryptedData
	}

	if i.TiktokToken != nil {
		byteTiktokToken, _ := json.Marshal(i.TiktokToken)
		encryptedData, _ := Helper.EncryptData(byteTiktokToken)
		i.TiktokToken = &encryptedData
	}
}

func (i *IntegratedProfile) DecryptTokenData() {
	if i.ShopeeToken != nil {
		decryptedData, _ := Helper.DecryptData(*i.ShopeeToken)
		var tokenData string
		_ = json.Unmarshal(decryptedData, &tokenData)
		i.ShopeeToken = &tokenData
	}

	if i.TiktokToken != nil {
		decryptedData, _ := Helper.DecryptData(*i.TiktokToken)
		var tokenData string
		_ = json.Unmarshal(decryptedData, &tokenData)
		i.TiktokToken = &tokenData
	}
}
