package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteMonitorGroupNotifyPolicy invokes the cms.DeleteMonitorGroupNotifyPolicy API synchronously
func (client *Client) DeleteMonitorGroupNotifyPolicy(request *DeleteMonitorGroupNotifyPolicyRequest) (response *DeleteMonitorGroupNotifyPolicyResponse, err error) {
	response = CreateDeleteMonitorGroupNotifyPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMonitorGroupNotifyPolicyWithChan invokes the cms.DeleteMonitorGroupNotifyPolicy API asynchronously
func (client *Client) DeleteMonitorGroupNotifyPolicyWithChan(request *DeleteMonitorGroupNotifyPolicyRequest) (<-chan *DeleteMonitorGroupNotifyPolicyResponse, <-chan error) {
	responseChan := make(chan *DeleteMonitorGroupNotifyPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMonitorGroupNotifyPolicy(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteMonitorGroupNotifyPolicyWithCallback invokes the cms.DeleteMonitorGroupNotifyPolicy API asynchronously
func (client *Client) DeleteMonitorGroupNotifyPolicyWithCallback(request *DeleteMonitorGroupNotifyPolicyRequest, callback func(response *DeleteMonitorGroupNotifyPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMonitorGroupNotifyPolicyResponse
		var err error
		defer close(result)
		response, err = client.DeleteMonitorGroupNotifyPolicy(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteMonitorGroupNotifyPolicyRequest is the request struct for api DeleteMonitorGroupNotifyPolicy
type DeleteMonitorGroupNotifyPolicyRequest struct {
	*requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	GroupId    string `position:"Query" name:"GroupId"`
}

// DeleteMonitorGroupNotifyPolicyResponse is the response struct for api DeleteMonitorGroupNotifyPolicy
type DeleteMonitorGroupNotifyPolicyResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    int    `json:"Result" xml:"Result"`
}

// CreateDeleteMonitorGroupNotifyPolicyRequest creates a request to invoke DeleteMonitorGroupNotifyPolicy API
func CreateDeleteMonitorGroupNotifyPolicyRequest() (request *DeleteMonitorGroupNotifyPolicyRequest) {
	request = &DeleteMonitorGroupNotifyPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteMonitorGroupNotifyPolicy", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMonitorGroupNotifyPolicyResponse creates a response to parse from DeleteMonitorGroupNotifyPolicy response
func CreateDeleteMonitorGroupNotifyPolicyResponse() (response *DeleteMonitorGroupNotifyPolicyResponse) {
	response = &DeleteMonitorGroupNotifyPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
