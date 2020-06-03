package hwapi

type Version struct {
	server string
	mod    string
}

const versionMod = "0.0.1"

func GetVersion() string {
	return versionMod
}

//Check current API verion
func (api *hwapi) GetServerVersion() (string, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    "/version",
		},
	)
	if e != nil {
		return "", e
	}

	return r.Headers.Get("X-Cdnws-Version"), nil
}
