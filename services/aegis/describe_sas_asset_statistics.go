package aegis

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

// DescribeSasAssetStatistics invokes the aegis.DescribeSasAssetStatistics API synchronously
// api document: https://help.aliyun.com/api/aegis/describesasassetstatistics.html
func (client *Client) DescribeSasAssetStatistics(request *DescribeSasAssetStatisticsRequest) (response *DescribeSasAssetStatisticsResponse, err error) {
	response = CreateDescribeSasAssetStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSasAssetStatisticsWithChan invokes the aegis.DescribeSasAssetStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesasassetstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSasAssetStatisticsWithChan(request *DescribeSasAssetStatisticsRequest) (<-chan *DescribeSasAssetStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeSasAssetStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSasAssetStatistics(request)
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

// DescribeSasAssetStatisticsWithCallback invokes the aegis.DescribeSasAssetStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesasassetstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSasAssetStatisticsWithCallback(request *DescribeSasAssetStatisticsRequest, callback func(response *DescribeSasAssetStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSasAssetStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSasAssetStatistics(request)
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

// DescribeSasAssetStatisticsRequest is the request struct for api DescribeSasAssetStatistics
type DescribeSasAssetStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp         string           `position:"Query" name:"SourceIp"`
	StatisticsColumn string           `position:"Query" name:"StatisticsColumn"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage      requests.Integer `position:"Query" name:"CurrentPage"`
	Uuids            string           `position:"Query" name:"Uuids"`
}

// DescribeSasAssetStatisticsResponse is the response struct for api DescribeSasAssetStatistics
type DescribeSasAssetStatisticsResponse struct {
	*responses.BaseResponse
	RequestId   string  `json:"RequestId" xml:"RequestId"`
	PageSize    int     `json:"PageSize" xml:"PageSize"`
	CurrentPage int     `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int     `json:"TotalCount" xml:"TotalCount"`
	AssetList   []Asset `json:"AssetList" xml:"AssetList"`
}

// CreateDescribeSasAssetStatisticsRequest creates a request to invoke DescribeSasAssetStatistics API
func CreateDescribeSasAssetStatisticsRequest() (request *DescribeSasAssetStatisticsRequest) {
	request = &DescribeSasAssetStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeSasAssetStatistics", "vipaegis", "openAPI")
	return
}

// CreateDescribeSasAssetStatisticsResponse creates a response to parse from DescribeSasAssetStatistics response
func CreateDescribeSasAssetStatisticsResponse() (response *DescribeSasAssetStatisticsResponse) {
	response = &DescribeSasAssetStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
