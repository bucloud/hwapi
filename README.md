# hwapi
This package supports all APIs of Highwinds/Striketracker
All API are under hwapi
# Usage
You can init api object like
api := hwapi.Init(http.Transport{
		MaxIdleConns:    3,
		IdleConnTimeout: 60,
	})

Note, Init function accept http.Transport as default httpClient config.

Get auth token, just use
api.auth(your_username,your_password) or api.SetToken(your_token)

Always keep in mind that,
1. You can use any APIs after auth/SetToken, Authinfo would automatic append to all API requests.
2. All request return same struct as API doc showed.
