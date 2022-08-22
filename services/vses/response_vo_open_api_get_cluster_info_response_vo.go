/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type ResponseVoOpenApiGetClusterInfoResponseVo struct {
	Code      int32                            `json:"code,omitempty"`
	Message   string                           `json:"message,omitempty"`
	RequestId string                           `json:"requestId,omitempty"`
	Result    *OpenApiGetClusterInfoResponseVo `json:"result,omitempty"`
}
