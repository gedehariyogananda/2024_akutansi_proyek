package Dto

import (
	"2024_akutansi_project/Helper"
	"encoding/json"
)

type IntegratedProfile struct {
	UserID string       `json:"user_id" bson:"user_id"`
	Shopee *Integration `json:"shopee,omitempty" bson:"shopee,omitempty"`
	Tiktok *Integration `json:"tiktok,omitempty" bson:"tiktok,omitempty"`
}

type Integration struct {
	Integrated bool `json:"integrated" bson:"integrated"`
	Credential any  `json:"credential" bson:"credential"`
}

type ShopeeIntegrateRequest struct {
	PartnerID  string `json:"partner_id"`
	PartnerKey string `json:"partner_key"`
}

type TiktokIntegrateRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (i *IntegratedProfile) EncryptTokenData() {
	if i.Shopee != nil {
		byteShopeeToken, _ := json.Marshal(i.Shopee.Credential)
		encryptedData, _ := Helper.EncryptData(byteShopeeToken)
		i.Shopee = &Integration{
			Integrated: i.Shopee.Integrated,
			Credential: encryptedData,
		}
	}

	if i.Tiktok != nil {
		byteTiktokToken, _ := json.Marshal(i.Tiktok.Credential)
		encryptedData, _ := Helper.EncryptData(byteTiktokToken)
		i.Tiktok = &Integration{
			Integrated: i.Tiktok.Integrated,
			Credential: encryptedData,
		}
	}
}

func (i *IntegratedProfile) DecryptTokenData() {
	if i.Shopee != nil {
		decryptedData, _ := Helper.DecryptData(i.Shopee.Credential.(string))
		var tokenData map[string]string
		_ = json.Unmarshal(decryptedData, &tokenData)
		i.Shopee = &Integration{
			Integrated: i.Shopee.Integrated,
			Credential: tokenData,
		}
	}

	if i.Tiktok != nil {
		decryptedData, _ := Helper.DecryptData(i.Tiktok.Credential.(string))
		var tokenData map[string]string
		_ = json.Unmarshal(decryptedData, &tokenData)
		i.Tiktok = &Integration{
			Integrated: i.Tiktok.Integrated,
			Credential: tokenData,
		}
	}
}
