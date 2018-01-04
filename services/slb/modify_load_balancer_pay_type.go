package slb

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

func (client *Client) ModifyLoadBalancerPayType(request *ModifyLoadBalancerPayTypeRequest) (response *ModifyLoadBalancerPayTypeResponse, err error) {
	response = CreateModifyLoadBalancerPayTypeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyLoadBalancerPayTypeWithChan(request *ModifyLoadBalancerPayTypeRequest) (<-chan *ModifyLoadBalancerPayTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyLoadBalancerPayTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLoadBalancerPayType(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) ModifyLoadBalancerPayTypeWithCallback(request *ModifyLoadBalancerPayTypeRequest, callback func(response *ModifyLoadBalancerPayTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLoadBalancerPayTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyLoadBalancerPayType(request)
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

type ModifyLoadBalancerPayTypeRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Duration             requests.Integer `position:"Query" name:"Duration"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	PayType              string           `position:"Query" name:"PayType"`
	PricingCycle         string           `position:"Query" name:"PricingCycle"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

type ModifyLoadBalancerPayTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   int    `json:"OrderId" xml:"OrderId"`
}

func CreateModifyLoadBalancerPayTypeRequest() (request *ModifyLoadBalancerPayTypeRequest) {
	request = &ModifyLoadBalancerPayTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerPayType", "", "")
	return
}

func CreateModifyLoadBalancerPayTypeResponse() (response *ModifyLoadBalancerPayTypeResponse) {
	response = &ModifyLoadBalancerPayTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}