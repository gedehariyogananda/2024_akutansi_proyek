package Dto

type (
	SetPushNotificationShopeeRequest struct {
		PartnerID         string `json:"partner_id"`
		PartnerKey        string `json:"partner_key"`
		BlockedShopIdList []int  `json:"blocked_shop_id_list"`
		CallbackUrl       string `json:"callback_url"`
		SetPushConfigOff  []int  `json:"set_push_config_off"`
		SetPushConfigOn   []int  `json:"set_push_config_on"`
	}
)
