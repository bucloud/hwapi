package hwapi

// Version info
type Version struct {
	server string
	mod    string
}

const versionMod = "0.1.5"

// GetVersion current hwapi version
func GetVersion() string {
	return versionMod
}

// GetServerVersion Check current API verion
func (api *HWApi) GetServerVersion() (string, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			URL:    "/version",
		},
	)
	if e != nil {
		return "", e
	}

	return r.Headers.Get("X-Cdnws-Version"), nil
}
