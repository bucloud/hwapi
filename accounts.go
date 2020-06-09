package hwapi

import (
	"encoding/json"
	"fmt"
)

//Contact information for the billing contact
type BillingDetails struct {
	BillingAccountNumber string `json:"billingAccountNumber"` //The account number for this account in the parent account's billing system
	Contact
}

//Contact information for a contact on a CDN account
type Contact struct {
	Id        int    `json:"id"`        //Unique id for this contact
	FirstName string `json:"firstName"` //First name
	LastName  string `json:"lastName"`  //Last name
	Email     string `json:"email"`     //Email
	Phone     string `json:"phone"`     //Phone number
	Fax       string `json:"fax"`       //Fax number
}

//Activity
type Activity struct {
	Activity map[string]DateContent `json:"activity"` //Date activity
}

//Date Content for Activity
type DateContent struct {
	Purge       ActivityAction `json:"PURGE"`        //Purge
	ConfigPush  ActivityAction `json:"CONFIG_PUSH"`  //Config_push
	Login       ActivityAction `json:"LOGIN"`        //Login
	AccountEdit ActivityAction `json:"ACCOUNT_EDIT"` //Account_edit
	Barometer   ActivityAction `json:"BAROMETER"`    //Barometer
	Total       ActivityAction `json:"TOTAL"`        //Totals
}

//Activity action
type ActivityAction struct {
	Count    int `json:"count"`    //Purge count total
	Accounts int `json:"accounts"` //Accounts
}

//Representative list
type RepresentativeList struct {
	List []*Representatives `json:"list"` //Status
}

//Representatives
type Representatives struct {
	SalesEngineer SalesRepresentative `json:"SALESENGINEER"` //Sales engineer assigned to account
	Sales         SalesRepresentative `json:"SALES"`         //Primary sales person assigned to account
	Sales2        SalesRepresentative `json:"SALES2"`        //Secondary sales person assigned to account
}

//Sales representative
type SalesRepresentative struct {
	Username  string `json:"username"`  //Secondary sales person assigned to account
	Status    string `json:"status"`    //Secondary sales person assigned to account
	FirstName string `json:"firstName"` //Secondary sales person assigned to account
	LastName  string `json:"lastName"`  //Secondary sales person assigned to account
}

//Subaccount list
type SubaccountList struct {
	List []*SimpleAccount `json:"list"` //list
}

type SimpleAccount struct {
	ID            string `json:"id"`
	AccountName   string `json:"accountName"`
	AccountHash   string `json:"accountHash"`
	AccountStatus string `json:"accountStatus"`
}

//A CDN account
type Account struct {
	Id                        int         `json:"id"`                        //The id of the account
	AccountHash               string      `json:"accountHash"`               //The hash code of the account
	AccountName               string      `json:"accountName"`               //The name of the account
	SupportEmailAddress       string      `json:"supportEmailAddress"`       //The email address of the account's support contact
	BillingAccountId          string      `json:"billingAccountId"`          //The read-only unique identifier in the billing system for the account
	BillingAccountNumber      string      `json:"billingAccountNumber"`      //The phone number of the account's billing contact
	AccountStatus             string      `json:"accountStatus"`             //The account's active status
	Parent                    int         `json:"parent"`                    //The parent account id of the account
	SubAccountCreationEnabled bool        `json:"subAccountCreationEnabled"` //Determines whether the account is allowed to create subaccounts
	MaximumDirectSubAccounts  int         `json:"maximumDirectSubAccounts"`  //The maximum number of subaccounts allowed on this account
	MaximumHosts              int         `json:"maximumHosts"`              //The maximum number of hosts allowed on this account
	MaxHcsTenants             int         `json:"maxHcsTenants"`             //The maximum number of HCS tenants allowed on this account
	BillingContact            Contact     `json:"billingContact"`            //The accounts billing contact information
	PrimaryContact            Contact     `json:"primaryContact"`            //The account's primary contact information
	TechnicalContact          Contact     `json:"technicalContact"`          //The account's technical contact information
	SubAccounts               []*Account  `json:"subAccounts"`               //If part of a recursive subaccount fetch, this will contain the subaccounts of the account
	Services                  []*Services `json:"services"`                  //Services enabled on this account
}

//GET verb implementation for a single Account The account fetched is the one set in the Account Context Thus, the AccountContextMiddleware is required
//Path /api/v1/accounts/{account_hash}
func (api *HWApi) GetAccount(accountHash string) (*Account, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Account{}
	return al, json.Unmarshal(r.body, al)
}

//Handles DELETE for accounts
//Path /api/v1/accounts/{account_hash}
func (api *HWApi) DeleteAccount(accountHash string) (bool, error) {
	_, e := api.Request(
		&Request{
			Method: DELETE,
			Url:    fmt.Sprintf("/api/v1/accounts/%s", accountHash),
		},
	)
	if e != nil {
		return false, e
	}
	return true, nil
}

//Handles PUT for accounts, updating an existing account NOTE: only send configuration when absolutely necessary, and only for the accounts which are necessary to update
//Path /api/v1/accounts/{account_hash}
func (api *HWApi) UpdateAccount(accountHash string, accountInfo Account) (*Account, error) {
	r, e := api.Request(
		&Request{
			Method: PUT,
			Url:    fmt.Sprintf("/api/v1/accounts/%s", accountHash),
			Body:   accountInfo,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Account{}
	return al, json.Unmarshal(r.body, al)
}

//Get account activity
//Path /api/v1/accounts/{account_hash}/activity
func (api *HWApi) GetAccountActivity(accountHash string) (*Activity, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/activity", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Activity{}
	return al, json.Unmarshal(r.body, al)
}

//Get sales reps
//Path /api/v1/accounts/{account_hash}/representatives
func (api *HWApi) GetSales(accountHash string) (*RepresentativeList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/representatives", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &RepresentativeList{}
	return al, json.Unmarshal(r.body, al)
}

//GET verb implementation for a single Account's sub-accounts The account fetched is the one set in the Account Context Thus, the AccountContextMiddleware is required
//Path /api/v1/accounts/{account_hash}/subaccounts
func (api *HWApi) GetSubaccounts(accountHash string, recursive string) (*Account, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v1/accounts/%s/subaccounts", accountHash),
			Query: map[string]string{
				"recursive": recursive,
			},
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Account{}
	return al, json.Unmarshal(r.body, al)
}

//Handles POST for accounts, creating a new account
//Path /api/v1/accounts/{parent_account_hash}
func (api *HWApi) CreateAccount(parentAccountHash string, accountInfo *Account) (*Account, error) {
	r, e := api.Request(
		&Request{
			Method: POST,
			Url:    fmt.Sprintf("/api/v1/accounts/%s", parentAccountHash),
			Body:   accountInfo,
		},
	)
	if e != nil {
		return nil, e
	}
	al := &Account{}
	return al, json.Unmarshal(r.body, al)
}

//GET a list of a single Account's sub-accounts The account fetched is the one set in the Account Context Thus, the AccountContextMiddleware is required This endpoint only returns minimal information regarding subaccounts
//Path /api/v3/accounts/{account_hash}/subaccounts
func (api *HWApi) GetSubaccounts2(accountHash string) (*SubaccountList, error) {
	r, e := api.Request(
		&Request{
			Method: GET,
			Url:    fmt.Sprintf("/api/v3/accounts/%s/subaccounts", accountHash),
		},
	)
	if e != nil {
		return nil, e
	}
	al := &SubaccountList{}
	return al, json.Unmarshal(r.body, al)
}
