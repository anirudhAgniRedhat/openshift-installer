package emr

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

// RunClusterServiceAction invokes the emr.RunClusterServiceAction API synchronously
func (client *Client) RunClusterServiceAction(request *RunClusterServiceActionRequest) (response *RunClusterServiceActionResponse, err error) {
	response = CreateRunClusterServiceActionResponse()
	err = client.DoAction(request, response)
	return
}

// RunClusterServiceActionWithChan invokes the emr.RunClusterServiceAction API asynchronously
func (client *Client) RunClusterServiceActionWithChan(request *RunClusterServiceActionRequest) (<-chan *RunClusterServiceActionResponse, <-chan error) {
	responseChan := make(chan *RunClusterServiceActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunClusterServiceAction(request)
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

// RunClusterServiceActionWithCallback invokes the emr.RunClusterServiceAction API asynchronously
func (client *Client) RunClusterServiceActionWithCallback(request *RunClusterServiceActionRequest, callback func(response *RunClusterServiceActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunClusterServiceActionResponse
		var err error
		defer close(result)
		response, err = client.RunClusterServiceAction(request)
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

// RunClusterServiceActionRequest is the request struct for api RunClusterServiceAction
type RunClusterServiceActionRequest struct {
	*requests.RpcRequest
	HostGroupIdList             *[]string        `position:"Query" name:"HostGroupIdList"  type:"Repeated"`
	ResourceOwnerId             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ServiceActionName           string           `position:"Query" name:"ServiceActionName"`
	IsRolling                   requests.Boolean `position:"Query" name:"IsRolling"`
	TotlerateFailCount          requests.Integer `position:"Query" name:"TotlerateFailCount"`
	ServiceName                 string           `position:"Query" name:"ServiceName"`
	ExecuteStrategy             string           `position:"Query" name:"ExecuteStrategy"`
	OnlyRestartStaleConfigNodes requests.Boolean `position:"Query" name:"OnlyRestartStaleConfigNodes"`
	NodeCountPerBatch           requests.Integer `position:"Query" name:"NodeCountPerBatch"`
	ClusterId                   string           `position:"Query" name:"ClusterId"`
	CustomCommand               string           `position:"Query" name:"CustomCommand"`
	ComponentNameList           string           `position:"Query" name:"ComponentNameList"`
	Comment                     string           `position:"Query" name:"Comment"`
	CustomParams                string           `position:"Query" name:"CustomParams"`
	Interval                    requests.Integer `position:"Query" name:"Interval"`
	HostIdList                  string           `position:"Query" name:"HostIdList"`
	TurnOnMaintenanceMode       requests.Boolean `position:"Query" name:"TurnOnMaintenanceMode"`
}

// RunClusterServiceActionResponse is the response struct for api RunClusterServiceAction
type RunClusterServiceActionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRunClusterServiceActionRequest creates a request to invoke RunClusterServiceAction API
func CreateRunClusterServiceActionRequest() (request *RunClusterServiceActionRequest) {
	request = &RunClusterServiceActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "RunClusterServiceAction", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRunClusterServiceActionResponse creates a response to parse from RunClusterServiceAction response
func CreateRunClusterServiceActionResponse() (response *RunClusterServiceActionResponse) {
	response = &RunClusterServiceActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
