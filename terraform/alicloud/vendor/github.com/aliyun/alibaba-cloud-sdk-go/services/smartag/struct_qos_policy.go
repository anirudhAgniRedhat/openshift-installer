package smartag

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

// QosPolicy is a nested struct in smartag response
type QosPolicy struct {
	IpProtocol      string                               `json:"IpProtocol" xml:"IpProtocol"`
	Name            string                               `json:"Name" xml:"Name"`
	Priority        int                                  `json:"Priority" xml:"Priority"`
	SourcePortRange string                               `json:"SourcePortRange" xml:"SourcePortRange"`
	DestPortRange   string                               `json:"DestPortRange" xml:"DestPortRange"`
	SourceCidr      string                               `json:"SourceCidr" xml:"SourceCidr"`
	StartTime       string                               `json:"StartTime" xml:"StartTime"`
	QosId           string                               `json:"QosId" xml:"QosId"`
	EndTime         string                               `json:"EndTime" xml:"EndTime"`
	DestCidr        string                               `json:"DestCidr" xml:"DestCidr"`
	Description     string                               `json:"Description" xml:"Description"`
	QosPolicyId     string                               `json:"QosPolicyId" xml:"QosPolicyId"`
	DpiSignatureIds DpiSignatureIdsInDescribeQosPolicies `json:"DpiSignatureIds" xml:"DpiSignatureIds"`
	DpiGroupIds     DpiGroupIdsInDescribeQosPolicies     `json:"DpiGroupIds" xml:"DpiGroupIds"`
}
