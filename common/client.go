package common

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hdksky/ksyungo/util"
)

// A Client represents a client of kec services
type Client struct {
	AccessKeyId     string //Access Key Id
	AccessKeySecret string //Access Key Secret
	Region          string
	Service         string
	debug           bool
	httpClient      *http.Client
	endpoint        string
	version         string
}

// Init creates a new instance of client
func (client *Client) Init(endpoint, version, accessKeyId, accessKeySecret, region, service string) {
	client.AccessKeyId = accessKeyId
	client.AccessKeySecret = accessKeySecret
	client.Region = region
	client.Service = service
	client.debug = false
	client.httpClient = &http.Client{}
	client.endpoint = endpoint
	client.version = version
}

// SetEndpoint sets custom endpoint
func (client *Client) SetEndpoint(endpoint string) {
	client.endpoint = endpoint
}

// SetEndpoint sets custom version
func (client *Client) SetVersion(version string) {
	client.version = version
}

// SetAccessKeyId sets new AccessKeyId
func (client *Client) SetAccessKeyId(id string) {
	client.AccessKeyId = id
}

// SetAccessKeySecret sets new AccessKeySecret
func (client *Client) SetAccessKeySecret(secret string) {
	client.AccessKeySecret = secret + "&"
}

// SetDebug sets debug mode to log the request/response message
func (client *Client) SetDebug(debug bool) {
	client.debug = debug
}

func formatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}

// Invoke sends the raw HTTP request for ksyun services
func (client *Client) Invoke(action string, args interface{}, response interface{}) error {

	request := Request{}
	request.init(client.version, action, client.AccessKeyId)

	query := util.ConvertToQueryValues(request)
	util.SetQueryValues(args, &query)

	// Generate the request URL
	requestURL := client.endpoint + "?" + query.Encode()

	httpReq, err := http.NewRequest(RequestMethod, requestURL, nil)
	if err != nil {
		return GetClientError(err)
	}

	// TODO move to util and add build val flag
	httpReq.Header.Set("X-SDK-Client", `KsyunGO/`+Version)
	//httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sign4(httpReq, Credentials{AccessKeyID: client.AccessKeyId, SecretAccessKey: client.AccessKeySecret}, client.Region, client.Service)

	t0 := time.Now()
	httpResp, err := client.httpClient.Do(httpReq)
	t1 := time.Now()
	if err != nil {
		return GetClientError(err)
	}
	statusCode := httpResp.StatusCode

	if client.debug {
		log.Printf("Invoke %s %s %d (%v)", RequestMethod, requestURL, statusCode, t1.Sub(t0))
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return GetClientError(err)
	}

	if client.debug {
		data, _ := formatXML(body)
		log.Println(string(data))
	}

	if statusCode >= 400 && statusCode <= 599 {
		errorResponse := ErrorResponse{}
		err = xml.Unmarshal(body, &errorResponse)
		ecsError := &Error{
			ErrorResponse: errorResponse,
			StatusCode:    statusCode,
		}
		return ecsError
	}

	err = xml.Unmarshal(body, response)
	//log.Printf("%++v", response)
	if err != nil {
		return GetClientError(err)
	}

	return nil
}

// GenerateClientToken generates the Client Token with random string
func (client *Client) GenerateClientToken() string {
	return util.CreateRandomString()
}

func GetClientErrorFromString(str string) error {
	return &Error{
		ErrorResponse: ErrorResponse{
			Code:    "KsyunGoClientFailure",
			Message: str,
		},
		StatusCode: -1,
	}
}

func GetClientError(err error) error {
	return GetClientErrorFromString(err.Error())
}
