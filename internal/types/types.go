package types

type Config struct {
	Endpoint      string `json:"endpoint"`
	Authorization string `json:"authorization"`
	Proxy         string `json:"proxy"`
}
