package kec

//Modify with your Access Key Id and Access Key Secret

const (
	TestAccessKeyId     = "MY_ID"
	TestAccessKeySecret = "MY_Secret"
	TestInstanceId      = "MY_TEST_INSTANCE_ID"
	TestSecurityGroupId = "MY_TEST_SECURITY_GROUP_ID"
	TestImageId         = "MY_TEST_IMAGE_ID"
	TestAccountId       = "MY_TEST_ACCOUNT_ID" //Get from https://account.console.aliyun.com

	TestIAmRich = false
	TestQuick   = false
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret, "cn-beijing-6")
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret, "cn-beijing-6")
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
