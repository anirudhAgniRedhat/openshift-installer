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

package eci

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteImageCache invokes the eci.DeleteImageCache API synchronously
// api document: https://help.aliyun.com/api/eci/deleteimagecache.html
func (client *Client) DeleteImageCache(request *DeleteImageCacheRequest) (response *DeleteImageCacheResponse, err error) {
	response = CreateDeleteImageCacheResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteImageCacheWithChan invokes the eci.DeleteImageCache API asynchronously
// api document: https://help.aliyun.com/api/eci/deleteimagecache.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteImageCacheWithChan(request *DeleteImageCacheRequest) (<-chan *DeleteImageCacheResponse, <-chan error) {
	responseChan := make(chan *DeleteImageCacheResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteImageCache(request)
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

// DeleteImageCacheWithCallback invokes the eci.DeleteImageCache API asynchronously
// api document: https://help.aliyun.com/api/eci/deleteimagecache.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteImageCacheWithCallback(request *DeleteImageCacheRequest, callback func(response *DeleteImageCacheResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteImageCacheResponse
		var err error
		defer close(result)
		response, err = client.DeleteImageCache(request)
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

// DeleteImageCacheRequest is the request struct for api DeleteImageCache
type DeleteImageCacheRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	RegionId             string           `position:"Query" name:"RegionId"`
	ImageCacheId         string           `position:"Query" name:"ImageCacheId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	VkClientVersion      string           `position:"Query" name:"VkClientVersion"`
}

// DeleteImageCacheResponse is the response struct for api DeleteImageCache
type DeleteImageCacheResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteImageCacheRequest creates a request to invoke DeleteImageCache API
func CreateDeleteImageCacheRequest() (request *DeleteImageCacheRequest) {
	request = &DeleteImageCacheRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eci", "2018-08-08", "DeleteImageCache", "eci", "openAPI")
	return
}

// CreateDeleteImageCacheResponse creates a response to parse from DeleteImageCache response
func CreateDeleteImageCacheResponse() (response *DeleteImageCacheResponse) {
	response = &DeleteImageCacheResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
