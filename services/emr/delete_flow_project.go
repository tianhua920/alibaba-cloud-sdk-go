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

// DeleteFlowProject invokes the emr.DeleteFlowProject API synchronously
// api document: https://help.aliyun.com/api/emr/deleteflowproject.html
func (client *Client) DeleteFlowProject(request *DeleteFlowProjectRequest) (response *DeleteFlowProjectResponse, err error) {
	response = CreateDeleteFlowProjectResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFlowProjectWithChan invokes the emr.DeleteFlowProject API asynchronously
// api document: https://help.aliyun.com/api/emr/deleteflowproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFlowProjectWithChan(request *DeleteFlowProjectRequest) (<-chan *DeleteFlowProjectResponse, <-chan error) {
	responseChan := make(chan *DeleteFlowProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFlowProject(request)
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

// DeleteFlowProjectWithCallback invokes the emr.DeleteFlowProject API asynchronously
// api document: https://help.aliyun.com/api/emr/deleteflowproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFlowProjectWithCallback(request *DeleteFlowProjectRequest, callback func(response *DeleteFlowProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFlowProjectResponse
		var err error
		defer close(result)
		response, err = client.DeleteFlowProject(request)
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

// DeleteFlowProjectRequest is the request struct for api DeleteFlowProject
type DeleteFlowProjectRequest struct {
	*requests.RpcRequest
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DeleteFlowProjectResponse is the response struct for api DeleteFlowProject
type DeleteFlowProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteFlowProjectRequest creates a request to invoke DeleteFlowProject API
func CreateDeleteFlowProjectRequest() (request *DeleteFlowProjectRequest) {
	request = &DeleteFlowProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowProject", "emr", "openAPI")
	return
}

// CreateDeleteFlowProjectResponse creates a response to parse from DeleteFlowProject response
func CreateDeleteFlowProjectResponse() (response *DeleteFlowProjectResponse) {
	response = &DeleteFlowProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
