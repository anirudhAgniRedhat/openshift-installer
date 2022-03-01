package polardb

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

// DescribeDBClusters invokes the polardb.DescribeDBClusters API synchronously
func (client *Client) DescribeDBClusters(request *DescribeDBClustersRequest) (response *DescribeDBClustersResponse, err error) {
	response = CreateDescribeDBClustersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClustersWithChan invokes the polardb.DescribeDBClusters API asynchronously
func (client *Client) DescribeDBClustersWithChan(request *DescribeDBClustersRequest) (<-chan *DescribeDBClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusters(request)
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

// DescribeDBClustersWithCallback invokes the polardb.DescribeDBClusters API asynchronously
func (client *Client) DescribeDBClustersWithCallback(request *DescribeDBClustersRequest, callback func(response *DescribeDBClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusters(request)
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

// DescribeDBClustersRequest is the request struct for api DescribeDBClusters
type DescribeDBClustersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer         `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string                   `position:"Query" name:"DBClusterDescription"`
	DBClusterStatus      string                   `position:"Query" name:"DBClusterStatus"`
	PageNumber           requests.Integer         `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                   `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer         `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeDBClustersTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string                   `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                   `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer         `position:"Query" name:"OwnerId"`
	DBType               string                   `position:"Query" name:"DBType"`
	DBClusterIds         string                   `position:"Query" name:"DBClusterIds"`
}

// DescribeDBClustersTag is a repeated param struct in DescribeDBClustersRequest
type DescribeDBClustersTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeDBClustersResponse is the response struct for api DescribeDBClusters
type DescribeDBClustersResponse struct {
	*responses.BaseResponse
	TotalRecordCount int                       `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int                       `json:"PageRecordCount" xml:"PageRecordCount"`
	RequestId        string                    `json:"RequestId" xml:"RequestId"`
	PageNumber       int                       `json:"PageNumber" xml:"PageNumber"`
	Items            ItemsInDescribeDBClusters `json:"Items" xml:"Items"`
}

// CreateDescribeDBClustersRequest creates a request to invoke DescribeDBClusters API
func CreateDescribeDBClustersRequest() (request *DescribeDBClustersRequest) {
	request = &DescribeDBClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusters", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClustersResponse creates a response to parse from DescribeDBClusters response
func CreateDescribeDBClustersResponse() (response *DescribeDBClustersResponse) {
	response = &DescribeDBClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
