package yext

const createExistingSubAccountPath = "existingsubaccountaddrequest"
const createExistingLocationPath = "existinglocationaddrequests"
const listLocationServicesPath = "services"

type ServicesService struct {
	client *Client
}

type ExistingSubAccountAddRequest struct {
	SubAccountId string        `json:"subAccountId"`
	SkuAdditions []SkuAddition `json:"skuAdditions"`
	AgreementId  string        `json:"agreementId"`
}

type ExistingLocationAddRequest struct {
	ExistingLocationId        string   `json:"existingLocationId"`
	ExistingLocationAccountId string   `json:"existingLocationAccountId"`
	Skus                      []string `json:"skus"`
	AgreementId               *string   `json:"agreementId,omitempty"`
	ForceReview               bool   `json:"forceReview"`
}

type SkuAddition struct {
	Sku      string `json:"sku"`
	Quantity string `json:"quantity"`
}

type ExistingSubAccountAddResponse struct {
	Id            int           `json:"id"`
	SubAccountId  string        `json:"subAccountId"`
	SkuAdditions  []SkuAddition `json:"skuAdditions"`
	AgreementId   string        `json:"agreementId"`
	Status        string        `json:"status"`
	DateSubmitted string        `json:"dateSubmitted"`
	StatusDetail  string        `json:"statusDetail"`
}

type ExistingLocationAddResponse struct {
	Id                        int      `json:"id"`
	LocationMode              string   `json:"locationMode"`
	ExistingLocationId        string   `json:"existingLocationId"`
	NewLocationId             string   `json:"newLocationId"`
	NewLocationAccountId      string   `json:"newLocationAccountId"`
	NewLocationAccountName    string   `json:"newLocationAccountName"`
	NewAccountParentAccountId string   `json:"newAccountParentAccountId"`
	NewLocationData           string   `json:"newLocationData"`
	NewEntityData             string   `json:"newEntityData"`
	Skus                      []string `json:"skus"`
	AgreementId               int      `json:"agreementId"`
	Status                    string   `json:"status"`
	DateSubmitted             string   `json:"dateSubmitted"`
	DateCompleted             string   `json:"dateCompleted"`
	StatusDetail              string   `json:"statusDetail"`
	AddRequestId string `json:"addRequestId"`
}

type Service {
	Sku string `json:"sku"`
}

type ListLocationServicesResponse struct {
	Services []*Service `json:"services"`
}

func (a *ServicesService) CreateAddRequestExistingSubAccount(existingSubAccountAddRequest *ExistingSubAccountAddRequest) (*ExistingSubAccountAddResponse, *Response, error) {
	var v *ExistingSubAccountAddResponse
	r, err := a.client.DoRequest("POST", createExistingSubAccountPath, &v)
	if err != nil {
		return v, r, err
	}

	return v, r, nil
}

func (a *ServicesService) CreateAddRequestExistingLocation(existingLocationAddRequest *ExistingLocationAddRequest) (*ExistingLocationAddResponse, *Response, error) {
	var v *ExistingLocationAddResponse
	r, err := a.client.DoRequestJSON("POST", createExistingLocationPath, existingLocationAddRequest, &v)
	if err != nil {
		return v, r, err
	}

	return v, r, nil
}

func (a *ServicesService) ListLocationServices(locationId string) (*ListLocationServicesResponse, *Response, error) {
	var v *ListLocationServicesResponse
	r, err := a.client.DoRequest("GET", fmt.Sprintf("%s?locationId=%s", listLocationServicesPath, locationId), &v)
	if err != nil {
		return v, r, err
	}

	return v, r, nil
}
