package hwapi

import (
	"encoding/json"
	"fmt"
)

//Notification list
type NotificationList struct {
	List []*Notification `json:"list"` //list
}

//Notification
type Notification struct {
	Id          int       `json:"id"`          //Id
	CreatedDate string    `json:"createdDate"` //Created date
	Services    []*string `json:"services"`    //Services
	Subject     string    `json:"subject"`     //Notification subject
	Subtitle    string    `json:"subtitle"`    //Notification Subtitle
}

//Get notification list
//Path /api/v1/accounts/{account_hash}/notifications
func (api *HWApi) GetNotifications(accountHash string, includeMessage bool, startDate string, endDate string) (*NotificationList, error) {
	im := "false"
	if includeMessage {
		im = "true"
	}
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/notifications", accountHash),
			Query: map[string]string{
				"includeMessage": im,
				"startDate":      startDate,
				"endDate":        endDate,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	al := &NotificationList{}
	return al, json.Unmarshal(r.body, al)
}

//Get notification
//Path /api/v1/accounts/{account_hash}/notifications/{notification_id}
//Get notification
func (api *HWApi) GetNotification(accountHash string, notificationID int) (*Notification, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/notifications/%d", accountHash, notificationID),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Notification{}
	return al, json.Unmarshal(r.body, al)
}
