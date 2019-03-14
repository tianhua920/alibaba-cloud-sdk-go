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

// ListSupportedServiceName invokes the emr.ListSupportedServiceName API synchronously
// api document: https://help.aliyun.com/api/emr/listsupportedservicename.html
func (client *Client) ListSupportedServiceName(request *ListSupportedServiceNameRequest) (response *ListSupportedServiceNameResponse, err error) {
	response = CreateListSupportedServiceNameResponse()
	err = client.DoAction(request, response)
	return
}

// ListSupportedServiceNameWithChan invokes the emr.ListSupportedServiceName API asynchronously
// api document: https://help.aliyun.com/api/emr/listsupportedservicename.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSupportedServiceNameWithChan(request *ListSupportedServiceNameRequest) (<-chan *ListSupportedServiceNameResponse, <-chan error) {
	responseChan := make(chan *ListSupportedServiceNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSupportedServiceName(request)
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

// ListSupportedServiceNameWithCallback invokes the emr.ListSupportedServiceName API asynchronously
// api document: https://help.aliyun.com/api/emr/listsupportedservicename.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSupportedServiceNameWithCallback(request *ListSupportedServiceNameRequest, callback func(response *ListSupportedServiceNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSupportedServiceNameResponse
		var err error
		defer close(result)
		response, err = client.ListSupportedServiceName(request)
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

// ListSupportedServiceNameRequest is the request struct for api ListSupportedServiceName
type ListSupportedServiceNameRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
}

// ListSupportedServiceNameResponse is the response struct for api ListSupportedServiceName
type ListSupportedServiceNameResponse struct {
	*responses.BaseResponse
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	SupportedServiceNameList SupportedServiceNameList `json:"SupportedServiceNameList" xml:"SupportedServiceNameList"`
}

// CreateListSupportedServiceNameRequest creates a request to invoke ListSupportedServiceName API
func CreateListSupportedServiceNameRequest() (request *ListSupportedServiceNameRequest) {
	request = &ListSupportedServiceNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListSupportedServiceName", "emr", "openAPI")
	return
}

// CreateListSupportedServiceNameResponse creates a response to parse from ListSupportedServiceName response
func CreateListSupportedServiceNameResponse() (response *ListSupportedServiceNameResponse) {
	response = &ListSupportedServiceNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
