package cloudapi

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

// CreateLogConfig invokes the cloudapi.CreateLogConfig API synchronously
// api document: https://help.aliyun.com/api/cloudapi/createlogconfig.html
func (client *Client) CreateLogConfig(request *CreateLogConfigRequest) (response *CreateLogConfigResponse, err error) {
	response = CreateCreateLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLogConfigWithChan invokes the cloudapi.CreateLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createlogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLogConfigWithChan(request *CreateLogConfigRequest) (<-chan *CreateLogConfigResponse, <-chan error) {
	responseChan := make(chan *CreateLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLogConfig(request)
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

// CreateLogConfigWithCallback invokes the cloudapi.CreateLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createlogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLogConfigWithCallback(request *CreateLogConfigRequest, callback func(response *CreateLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLogConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateLogConfig(request)
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

// CreateLogConfigRequest is the request struct for api CreateLogConfig
type CreateLogConfigRequest struct {
	*requests.RpcRequest
	SlsLogStore   string `position:"Query" name:"SlsLogStore"`
	SlsProject    string `position:"Query" name:"SlsProject"`
	LogType       string `position:"Query" name:"LogType"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// CreateLogConfigResponse is the response struct for api CreateLogConfig
type CreateLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLogConfigRequest creates a request to invoke CreateLogConfig API
func CreateCreateLogConfigRequest() (request *CreateLogConfigRequest) {
	request = &CreateLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateLogConfig", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLogConfigResponse creates a response to parse from CreateLogConfig response
func CreateCreateLogConfigResponse() (response *CreateLogConfigResponse) {
	response = &CreateLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
