package growthpush

import (
	"net/url"
)

// RegisterNewClient クライアントデバイスを新規に登録
/**
 * token       クライアントデバイスのデバイストークン
 * os          OS名 (ios/android)
 * environment 環境設定 (development/production)
 */
func RegisterNewClient(pushToken string, os string) ([]byte, error) {
	// clients api
	client := NewClient("clients")

	// params
	values := url.Values{}
	values.Add("applicationId", client.GrowthPushConfig.ApplicationID)
	values.Add("secret", client.GrowthPushConfig.SecretKey)
	values.Add("token", pushToken)
	values.Add("os", os)
	if client.GrowthPushConfig.Production {
		values.Add("environment", "production")
	} else {
		values.Add("environment", "development")
	}

	// do request
	return client.Post(values)
}

// PutTagClient クライアントデバイスにタグ付け。既に同一名のタグが登録されている場合は、そのvalueを更新する
/**
 * token クライアントを識別するデバイストークン
 * name  タグ名
 */
func PutTagClient(pushToken string, name string) ([]byte, error) {
	// tags api
	client := NewClient("tags")

	// params
	values := url.Values{}
	values.Add("applicationId", client.GrowthPushConfig.ApplicationID)
	values.Add("secret", client.GrowthPushConfig.SecretKey)
	values.Add("token", pushToken)
	values.Add("name", name)

	// do request
	return client.Post(values)
}
