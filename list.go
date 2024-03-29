package soumuradio

import (
	"context"

	"github.com/tomato3713/soumuradio/internal"
)

// getListAPI is Method for getting number
func (c *Client) getListAPI(ctx context.Context, opt ListOpts) (*internal.Lists, error) {
	spath := "/musen/list"

	req, err := c.newRequest(ctx, "GET", spath, opt, nil)
	if err != nil {
		logf("try URL: %v", req.URL)
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err = checkStatusCode(res); err != nil {
		return nil, err
	}

	var lists internal.Lists
	if err := decodeBody(res, &lists, nil); err != nil {
		return nil, err
	}

	return &lists, nil
}

func (c *Client) GetRadioLicenseList(ctx context.Context, opt ListOpts) (LicenseList, error) {
	opt.ST = License
	result, err := c.getListAPI(ctx, opt)
	if err != nil {
		return LicenseList{}, err
	}

	ret, err := convertInternalLists2LicenseList(result)
	if err != nil {
		return LicenseList{}, err
	}

	return ret, nil
}

func (c *Client) GetRadioRegistrationList(ctx context.Context, opt ListOpts) (RegistrationList, error) {
	opt.ST = Registration
	result, err := c.getListAPI(ctx, opt)
	if err != nil {
		return RegistrationList{}, err
	}

	ret, err := convertInternalLists2RegistrationList(result)
	if err != nil {
		return RegistrationList{}, err
	}

	return ret, nil
}
