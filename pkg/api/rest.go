package api

import (
	"github.com/go-resty/resty/v2"
)

type RestOptions struct {
	Host    string
	Headers map[string]string
	Body    interface{}
	Retry   int
}

type restClient struct {
	client *resty.Client
}

func InitRestClient(opts RestOptions) *restClient {
	client := resty.New().SetRetryCount(opts.Retry).SetHeaders(opts.Headers).SetBaseURL(opts.Host)

	return &restClient{client: client}
}

func (r *restClient) InterceptorRequest(in func(c *resty.Client, req *resty.Request) error) {
	r.client.OnBeforeRequest(in)
}

func (r *restClient) InterceptorResponse(in func(c *resty.Client, res *resty.Response) error) {
	r.client.OnAfterResponse(in)

}

func (r *restClient) GET(url string, Query map[string]string) (*resty.Response, error) {
	res, err := r.client.R().SetQueryParams(Query).Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *restClient) POST(url string, data interface{}) (*resty.Response, error) {

	res, err := r.client.R().SetBody(data).Post(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *restClient) PUT(url string, data interface{}) (*resty.Response, error) {

	res, err := r.client.R().SetBody(data).Put(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *restClient) DELETE(url string, Query map[string]string) (*resty.Response, error) {

	res, err := r.client.R().SetQueryParams(Query).Delete(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
