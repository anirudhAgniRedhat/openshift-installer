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

// CreateResourceQueue invokes the emr.CreateResourceQueue API synchronously
func (client *Client) CreateResourceQueue(request *CreateResourceQueueRequest) (response *CreateResourceQueueResponse, err error) {
	response = CreateCreateResourceQueueResponse()
	err = client.DoAction(request, response)
	return
}

// CreateResourceQueueWithChan invokes the emr.CreateResourceQueue API asynchronously
func (client *Client) CreateResourceQueueWithChan(request *CreateResourceQueueRequest) (<-chan *CreateResourceQueueResponse, <-chan error) {
	responseChan := make(chan *CreateResourceQueueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateResourceQueue(request)
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

// CreateResourceQueueWithCallback invokes the emr.CreateResourceQueue API asynchronously
func (client *Client) CreateResourceQueueWithCallback(request *CreateResourceQueueRequest, callback func(response *CreateResourceQueueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateResourceQueueResponse
		var err error
		defer close(result)
		response, err = client.CreateResourceQueue(request)
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

// CreateResourceQueueRequest is the request struct for api CreateResourceQueue
type CreateResourceQueueRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer             `position:"Query" name:"ResourceOwnerId"`
	QualifiedName   string                       `position:"Query" name:"QualifiedName"`
	ResourcePoolId  requests.Integer             `position:"Query" name:"ResourcePoolId"`
	ClusterId       string                       `position:"Query" name:"ClusterId"`
	Leaf            requests.Boolean             `position:"Query" name:"Leaf"`
	ParentQueueId   requests.Integer             `position:"Query" name:"ParentQueueId"`
	Name            string                       `position:"Query" name:"Name"`
	Config          *[]CreateResourceQueueConfig `position:"Query" name:"Config"  type:"Repeated"`
}

// CreateResourceQueueConfig is a repeated param struct in CreateResourceQueueRequest
type CreateResourceQueueConfig struct {
	ConfigKey   string `name:"ConfigKey"`
	Note        string `name:"Note"`
	ConfigValue string `name:"ConfigValue"`
	Category    string `name:"Category"`
}

// CreateResourceQueueResponse is the response struct for api CreateResourceQueue
type CreateResourceQueueResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateResourceQueueRequest creates a request to invoke CreateResourceQueue API
func CreateCreateResourceQueueRequest() (request *CreateResourceQueueRequest) {
	request = &CreateResourceQueueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateResourceQueue", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateResourceQueueResponse creates a response to parse from CreateResourceQueue response
func CreateCreateResourceQueueResponse() (response *CreateResourceQueueResponse) {
	response = &CreateResourceQueueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
