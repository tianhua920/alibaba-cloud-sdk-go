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

// MetastoreCreateTable invokes the emr.MetastoreCreateTable API synchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatetable.html
func (client *Client) MetastoreCreateTable(request *MetastoreCreateTableRequest) (response *MetastoreCreateTableResponse, err error) {
	response = CreateMetastoreCreateTableResponse()
	err = client.DoAction(request, response)
	return
}

// MetastoreCreateTableWithChan invokes the emr.MetastoreCreateTable API asynchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatetable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreCreateTableWithChan(request *MetastoreCreateTableRequest) (<-chan *MetastoreCreateTableResponse, <-chan error) {
	responseChan := make(chan *MetastoreCreateTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetastoreCreateTable(request)
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

// MetastoreCreateTableWithCallback invokes the emr.MetastoreCreateTable API asynchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatetable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreCreateTableWithCallback(request *MetastoreCreateTableRequest, callback func(response *MetastoreCreateTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetastoreCreateTableResponse
		var err error
		defer close(result)
		response, err = client.MetastoreCreateTable(request)
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

// MetastoreCreateTableRequest is the request struct for api MetastoreCreateTable
type MetastoreCreateTableRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer                 `position:"Query" name:"ResourceOwnerId"`
	FieldDelimiter  string                           `position:"Query" name:"FieldDelimiter"`
	Column          *[]MetastoreCreateTableColumn    `position:"Query" name:"Column"  type:"Repeated"`
	CreateWith      string                           `position:"Query" name:"CreateWith"`
	Partition       *[]MetastoreCreateTablePartition `position:"Query" name:"Partition"  type:"Repeated"`
	DbName          string                           `position:"Query" name:"DbName"`
	CreateSql       string                           `position:"Query" name:"CreateSql"`
	Comment         string                           `position:"Query" name:"Comment"`
	LocationUri     string                           `position:"Query" name:"LocationUri"`
	TableName       string                           `position:"Query" name:"TableName"`
	DatabaseId      string                           `position:"Query" name:"DatabaseId"`
}

// MetastoreCreateTableColumn is a repeated param struct in MetastoreCreateTableRequest
type MetastoreCreateTableColumn struct {
	Name    string `name:"Name"`
	Comment string `name:"Comment"`
	Type    string `name:"Type"`
}

// MetastoreCreateTablePartition is a repeated param struct in MetastoreCreateTableRequest
type MetastoreCreateTablePartition struct {
	Name    string `name:"Name"`
	Comment string `name:"Comment"`
	Type    string `name:"Type"`
}

// MetastoreCreateTableResponse is the response struct for api MetastoreCreateTable
type MetastoreCreateTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMetastoreCreateTableRequest creates a request to invoke MetastoreCreateTable API
func CreateMetastoreCreateTableRequest() (request *MetastoreCreateTableRequest) {
	request = &MetastoreCreateTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "MetastoreCreateTable", "emr", "openAPI")
	return
}

// CreateMetastoreCreateTableResponse creates a response to parse from MetastoreCreateTable response
func CreateMetastoreCreateTableResponse() (response *MetastoreCreateTableResponse) {
	response = &MetastoreCreateTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
