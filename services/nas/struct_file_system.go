package nas

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

// FileSystem is a nested struct in nas response
type FileSystem struct {
	FileSystemId string                            `json:"FileSystemId" xml:"FileSystemId"`
	Description  string                            `json:"Description" xml:"Description"`
	CreateTime   string                            `json:"CreateTime" xml:"CreateTime"`
	RegionId     string                            `json:"RegionId" xml:"RegionId"`
	ProtocolType string                            `json:"ProtocolType" xml:"ProtocolType"`
	StorageType  string                            `json:"StorageType" xml:"StorageType"`
	MeteredSize  int64                             `json:"MeteredSize" xml:"MeteredSize"`
	ZoneId       string                            `json:"ZoneId" xml:"ZoneId"`
	Bandwidth    int64                             `json:"Bandwidth" xml:"Bandwidth"`
	Capacity     int64                             `json:"Capacity" xml:"Capacity"`
	Status       string                            `json:"Status" xml:"Status"`
	Ldap         Ldap                              `json:"Ldap" xml:"Ldap"`
	MountTargets MountTargetsInDescribeFileSystems `json:"MountTargets" xml:"MountTargets"`
	Packages     Packages                          `json:"Packages" xml:"Packages"`
}
