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

type RestClient struct {
	client *resty.Client
}

func InitRestClient(opts RestOptions) *RestClient {
	client := resty.New().SetRetryCount(opts.Retry).SetHeaders(opts.Headers).SetBaseURL(opts.Host)

	return &RestClient{client: client}
}

func (r *RestClient) InterceptorRequest(in func(c *resty.Client, req *resty.Request) error) {
	r.client.OnBeforeRequest(in)
}

func (r *RestClient) InterceptorResponse(in func(c *resty.Client, res *resty.Response) error) {
	r.client.OnAfterResponse(in)

}

func (r *RestClient) GET(url string, Query map[string]string) (*resty.Response, error) {
	res, err := r.client.R().SetQueryParams(Query).Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RestClient) POST(url string, data interface{}) (*resty.Response, error) {

	res, err := r.client.R().SetBody(data).Post(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RestClient) PUT(url string, data interface{}) (*resty.Response, error) {

	res, err := r.client.R().SetBody(data).Put(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RestClient) DELETE(url string, Query map[string]string) (*resty.Response, error) {

	res, err := r.client.R().SetQueryParams(Query).Delete(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
