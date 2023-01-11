/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v2)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses2

type ResponseVoOpenApiGetClusterInfoResponseVo struct {
	Code      int32                            `json:"code,omitempty"`
	Message   string                           `json:"message,omitempty"`
	RequestId string                           `json:"requestId,omitempty"`
	Result    *OpenApiGetClusterInfoResponseVo `json:"result,omitempty"`
}
