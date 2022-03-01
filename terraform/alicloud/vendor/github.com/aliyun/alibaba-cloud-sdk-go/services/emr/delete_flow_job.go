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

// DeleteFlowJob invokes the emr.DeleteFlowJob API synchronously
func (client *Client) DeleteFlowJob(request *DeleteFlowJobRequest) (response *DeleteFlowJobResponse, err error) {
	response = CreateDeleteFlowJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFlowJobWithChan invokes the emr.DeleteFlowJob API asynchronously
func (client *Client) DeleteFlowJobWithChan(request *DeleteFlowJobRequest) (<-chan *DeleteFlowJobResponse, <-chan error) {
	responseChan := make(chan *DeleteFlowJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFlowJob(request)
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

// DeleteFlowJobWithCallback invokes the emr.DeleteFlowJob API asynchronously
func (client *Client) DeleteFlowJobWithCallback(request *DeleteFlowJobRequest, callback func(response *DeleteFlowJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFlowJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteFlowJob(request)
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

// DeleteFlowJobRequest is the request struct for api DeleteFlowJob
type DeleteFlowJobRequest struct {
	*requests.RpcRequest
	Id        string `position:"Query" name:"Id"`
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DeleteFlowJobResponse is the response struct for api DeleteFlowJob
type DeleteFlowJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteFlowJobRequest creates a request to invoke DeleteFlowJob API
func CreateDeleteFlowJobRequest() (request *DeleteFlowJobRequest) {
	request = &DeleteFlowJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowJob", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFlowJobResponse creates a response to parse from DeleteFlowJob response
func CreateDeleteFlowJobResponse() (response *DeleteFlowJobResponse) {
	response = &DeleteFlowJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
