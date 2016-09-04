package authentication

import (
	"../provider"
)

func AuthenticateRequest(at string) (int64, error) {
	id, err := provider.RedisClient().Get("access_token:" + at).Int64()
	if err != nil {
		if err.Error() != "redis: nil" {
			return -1, err
		} else {
			return -1, nil
		}
	}

	return id, nil
}
