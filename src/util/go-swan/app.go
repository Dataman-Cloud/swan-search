package swan

import (
	"net/url"
)

// CreateApplication creates a new application in Swan
// application:		the structure holding the application configuration
func (r *swanClient) CreateApplication(version *Version) (*Application, error) {
	result := new(Application)
	if err := r.apiPost(swanAPIApps, &version, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Applications retrieves an array of all the applications in swan
func (r *swanClient) Applications(v url.Values) ([]*Application, error) {
	applications := new([]*Application)
	err := r.apiGet(swanAPIApps+"?"+v.Encode(), nil, applications)
	if err != nil {
		return nil, err
	}

	return *applications, nil
}

// DeleteApplication deletes an application in Swan
func (r *swanClient) DeleteApplication(appID string) error {
	if err := r.apiDelete(swanAPIApps+"/"+appID, nil, nil); err != nil {
		return err
	}

	return nil
}

// GetApplication retrieves an application from Swan
func (r *swanClient) GetApplication(appID string) (*Application, error) {
	result := new(Application)
	if err := r.apiGet(swanAPIApps+"/"+appID, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
