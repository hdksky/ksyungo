package slb

//Modify with your Access Key Id and Access Key Secret
const (
	TestAccessKeyId     = "AKLTUxyeuc11TQ2gRi2yJN7FiA"
	TestAccessKeySecret = "OHp/RfCLH+/c5rrH/+k0g9Mih3289ZonVyMITpordadELY6CzsECZcIc+X/oslbJCQ=="
	TestInstanceId      = "MY_INSTANCE_ID"
	TestRegion          = "cn-beijing-6"
	TestIAmRich         = false
	TestQuick           = false
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret, TestRegion)
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret, TestRegion)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
