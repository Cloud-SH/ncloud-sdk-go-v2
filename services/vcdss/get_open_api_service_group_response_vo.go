/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type GetOpenApiServiceGroupResponseVo struct {
	ComputeList []OpenApiGetComputeInstanceListResponseDetailVo `json:"computeList,omitempty"`
	ContractNo int32 `json:"contractNo,omitempty"`
	CreatedDate *DateTimeVo `json:"createdDate,omitempty"`
	ProductCode string `json:"productCode,omitempty"`
	RegionNo int32 `json:"regionNo,omitempty"`
	ServiceGroupInstanceNo int32 `json:"serviceGroupInstanceNo,omitempty"`
	ServiceGroupInstanceTypeCode string `json:"serviceGroupInstanceTypeCode,omitempty"`
	ServiceGroupName string `json:"serviceGroupName,omitempty"`
	StatusCode string `json:"statusCode,omitempty"`
	VpcNo int32 `json:"vpcNo,omitempty"`
}
