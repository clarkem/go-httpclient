package gohttp

/* packages encapsulate behaviour */

/* need a struct implementing interface to interact with interface */
type httpClient struct{}

func New() HttpClient {
	client := &httpClient{}
	return client
}

/* struct must implement all of these to be of this type */
type HttpClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

/* method that form public interface defined by methods in struct
   Not package function - need instance of struct
   Helps to keep public interface clean */

func (c *httpClient) Get() {}

func (c *httpClient) Post() {}

func (c *httpClient) Put() {}

func (c *httpClient) Patch() {}

func (c *httpClient) Delete() {}
