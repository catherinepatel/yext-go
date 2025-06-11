package yext

import (
	"fmt"
)

const createExistingSubAccountPath = "existingsubaccountaddrequest"
const createExistingLocationPath = "existinglocationaddrequests"
const listLocationServicesPath = "services"
const cancelServicesPath = "cancelservices"

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

type CancelServicesOnLocationRequest struct {
	LocationId        string   `json:"locationId"`
	LocationAccountId string   `json:"locationAccountId"`
	Skus                      []string `json:"skus"`
}

type CancelServicesOnLocationResponse struct {
}

type AddRequest struct {
	Status string `json:"status"`
	Skus []string `json:"skus"`
	ExistingLocationId string `json:"existingLocationId"`
}

type ListLocationServicesResponse struct {
	AddRequests []*AddRequest `json:"addRequests"`
}

type Service struct {
	Status string `json:"status"`
	Sku string `json:"sku"`
	ExistingLocationId string `json:"existingLocationId"`
}
	

type ListLocationsWithServiceResponse struct {
	Services []*Service `json:"services"`
	Count int `json:"count"`
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

func (a *ServicesService) ListLocationsWithService(sku string, offset int) (*ListLocationsWithServiceResponse, *Response, error) {
	var v *ListLocationsWithServiceResponse
	r, err := a.client.DoRequest("GET", fmt.Sprintf("%s?sku=%s&limit=1000&offset=%d", listLocationServicesPath, sku, offset), &v)
	if err != nil {
		return v, r, err
	}
	return v, r, nil
}

func (a *ServicesService) ListAllLocationsWithService(sku string) ([]*Service, *Response, error) {
	var services = []*Service{}
	listLocationsWithServicesResp, resp, err := a.ListLocationsWithService(sku)
	if err != nil {
		return services, resp, err
	}
	services = append(services, listLocationsWithServicesResp.Services)
	for len(services) != listLocationsWithServicesResp.Count {
		listLocationsWithServicesResp, resp, err = a.ListLocationsWithService(sku, len(services))
		if err != nil {
			return services, resp, err
		}
		services = append(services, listLocationsWithServicesResp.Services)
	}
	return services, nil, nil
		
}

func (a *ServicesService) ListLocationServices(locationId string) (*ListLocationServicesResponse, *Response, error) {
	var v *ListLocationServicesResponse
	r, err := a.client.DoRequest("GET", fmt.Sprintf("%s?locationId=%s", listLocationServicesPath, locationId), &v)
	if err != nil {
		return v, r, err
	}

	return v, r, nil
}

func (a *ServicesService) CancelServicesOnLocation(cancelServicesOnLocationRequest *CancelServicesOnLocationRequest) (*CancelServicesOnLocationResponse, *Response, error) {
	var v *CancelServicesOnLocationResponse
	r, err := a.client.DoRequestJSON("POST", cancelServicesPath, cancelServicesOnLocationRequest, &v)
	if err != nil {
		return v, r, err
	}

	return v, r, nil
}
