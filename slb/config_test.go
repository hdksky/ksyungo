package slb

//Modify with your Access Key Id and Access Key Secret
const (
	TestAccessKeyId     = "MY_ID"
	TestAccessKeySecret = "MY_SECRET"
	TestInstanceId      = "MY_INSTANCE_ID"
	Region              = "cn-beijing-6"
	TestIAmRich         = false
	TestQuick           = false
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret, Region)
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret, Region)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
