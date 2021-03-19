/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type DeleteScheduledActionRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 오토스케일링그룹번호
AutoScalingGroupNo *string `json:"autoScalingGroupNo"`

	// 스케쥴액션번호
ScheduledActionNo *string `json:"scheduledActionNo"`
}
