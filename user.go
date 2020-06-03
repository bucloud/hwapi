package hwapi

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id                       int         `json:"id"`          //The user id
	UserName                 string      `json:"username"`    //The username of the user
	Status                   string      `json:"status"`      //The status of the user
	OldPassword              string      `json:"oldPassword"` //The user's previous password, required to change another user's password
	Password                 string      `json:"password"`    //The password of the user
	Roles                    Roles       `json:"roles"`
	Preferences              Preferences `json:"preferences"`
	CreatedDate              string      `json:"createdDate"`              //createdDate
	UpdatedDate              string      `json:"updatedDate"`              //updatedDate
	LastLogin                string      `json:"lastLogin"`                //lastLogin
	AuthorizedSupportContact bool        `json:"authorizedSupportContact"` //authorizedSupportContact
	FirstName                string      `json:"firstName"`                //firstName
	LastName                 string      `json:"lastName"`                 //lastName
	Email                    string      `json:"email"`                    //email
	Phone                    string      `json:"phone"`                    //phone
	Fax                      string      `json:"fax"`                      //fax
	UserType                 string      `json:"userType"`                 //userType
	AccountHash              string      `json:"accountHash"`              //accountHash
	AccountName              string      `json:"accountName"`              //accountName
}

type Roles struct {
	UserAccount Role `json:"userAccount"`
	SubAccounts Role `json:"subAccounts"`
}
type Role struct {
	Report        string `json:"report"`
	Account       string `json:"account"`
	Content       string `json:"content"`
	Configuration string `json:"configuration"`
}

type Show struct {
	deleteEncodedAssetPopup bool
	embedCode               bool
	profileTable            bool
	forceEncodePopup        bool
	metaForm                bool
	transmux                bool
	transcodes              bool
}

type Filter struct {
	explicitPolicies  bool
	inheritedPolicies bool
	defaultPolicies   bool
}

type Preferences struct {
	Home                 string  `json:"home"`
	Sessions             Session `json:"session"`
	LastReadNotification string  `json:"lastReadNotification"`
	ShowJson             bool
	ExpandedHost         map[string][]string `json:"expandedHost"`
	Barometer_url        string              `json:"barometer_url"`
	DefaultPurgeType     string              `json:"defaultPurgeType"`
	UserTasks            TaskList            `json:"userTasks"`
	SeenSitesPage        bool                `json:"seenSitesPage"`
	SecureChatToken      ChatToken           `json:"secureChatToken"`
}

type ChatToken struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TaskList struct {
	TaskList        []*Task `json:"taskList"`
	PercentComplete float32 `json:"percentComplete"`
	Complete        bool    `json:"complete"`
	Closed          bool    `json:"closed"`
}

type Task struct {
	Name              string `json:"name"`
	Order             int    `json:"order"`
	Category          string `json:"category"`
	MinPermission     string `json:"minPermission"`
	TranslationString string `json:"translationString"`
	Complete          bool   `json:"complete"`
	SeenComplete      bool   `json:"seenComplete"`
}

type Sessions struct {
	ServiceControls                Session `json:"ServiceControls"`
	SubAccountControls             Session `json:"SubAccountControls"`
	UserControls                   Session `json:"UserControls"`
	OriginController               Session `json:"OriginController"`
	HOSTController                 Session `json:"HOSTController"`
	PolicyInheritanceHelper        Session `json:"PolicyInheritanceHelper"`
	AccountSummaryController       Session `json:"AccountSummaryController"`
	HostSummaryController          Session `json:"HostSummaryController"`
	HCSObjectsController           Session `json:"HCSObjectsController"`
	MediaManagementControls        Session `json:"MediaManagementControls"`
	MediaTPlaylistsControls        Session `json:"MediaTPlaylistsControls"`
	MediaPlaylistsAddControls      Session `json:"MediaPlaylistsAddControls"`
	MediaTagsControls              Session `json:"MediaTagsControls"`
	MediaProfileControls           Session `json:"MediaProfileControls"`
	MediaManagementEncodedControls Session `json:"MediaManagementEncodedControls"`
	MediaFtpControls               Session `json:"MediaFtpControls"`
	MediaPlaylistsEditController   Session `json:"MediaPlaylistsEditController"`
	HCSTenantController            Session `json:"HCSTenantController"`
	HCSContainerController         Session `json:"HCSContainerController"`
	SitesController                Session `json:"SitesController"`
	CertificatesController         Session `json:"CertificatesController"`
	OriginsController              Session `json:"OriginsController"`
}

type Session struct {
	ControllerName string   `json:"controllerName"`
	ItemsPerPage   int      `json:"itemsPerPage"`
	SortingOrder   string   `json:"sortingOrder"`
	Reverse        bool     `json:"reverse"`
	CurrentPage    int      `json:"currentPage"`
	ShowSuspended  bool     `json:"showSuspended"`
	ShowDeleted    bool     `json:"showDeleted"`
	Sort           string   `json:"sort"`
	Filters        Filter   `json:"filters"`
	OA             []Filter `json:"OA"`
	ExpandaItemId  int      `json:"expandaItemI"`
	Show           Show     `json:"show"`
}

type UserList struct {
	List []*User
}

type PasswordReset struct {
	PasswordReset string
}

func (api *hwapi) AboutMe() (*User, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    "/api/v1/users/me",
		},
	)
	if e != nil {
		return nil, e
	}
	return api.CurrentUser, json.Unmarshal(r.body, &api.CurrentUser)
}

//Detech wether username exists under currentAccount
//Note, this function had deprecated
func (api *hwapi) HasUser(username string) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: GET,
			Url:    "/api/v1/users/" + username,
		},
	)
	if e != nil {
		return true, nil
	}
	return false, e
}

//Update currentUser
func (api *hwapi) UpdateMe(user *User) (*User, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			Url:    "/api/v1/users/me",
			Body:   &user,
		},
	)
	if e != nil {
		return nil, e
	}
	return api.CurrentUser, json.Unmarshal(r.body, &api.CurrentUser)
}

//Update user
func (api *hwapi) UpdateUser(accountHash string, uid int, user *User) (*User, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d", accountHash, uid),
			Body:   &user,
		},
	)
	if e != nil {
		return nil, e
	}
	return api.CurrentUser, json.Unmarshal(r.body, &api.CurrentUser)
}

//Get user info by userID under account
func (api *hwapi) AboutUser(accountHash string, uid int) (*User, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d", accountHash, uid),
		},
	)
	if e != nil {
		return nil, e
	}
	return api.CurrentUser, json.Unmarshal(r.body, &api.CurrentUser)
}

//Delete user by userID under account
func (api *hwapi) DeleteUser(accountHash string, uid int) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users/%d", accountHash, uid),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

//Get users for account
func (api *hwapi) GetUsers(accountHash string) (*UserList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	ul := &UserList{}
	return ul, json.Unmarshal(r.body, ul)
}

//Greate new user for account
func (api *hwapi) CreateUser(accountHash string, user *User) (*User, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/users", accountHash),
			Body:   &user,
		},
	)
	if e != nil {
		return nil, e
	}
	u := &User{}
	return u, json.Unmarshal(r.body, u)
}
