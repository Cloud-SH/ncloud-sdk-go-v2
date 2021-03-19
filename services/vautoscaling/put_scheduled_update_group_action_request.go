/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type PutScheduledUpdateGroupActionRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 오토스케일링그룹번호
AutoScalingGroupNo *string `json:"autoScalingGroupNo"`

	// 스케쥴액션번호
ScheduledActionNo *string `json:"scheduledActionNo,omitempty"`

	// 스케쥴액션이름
ScheduledActionName *string `json:"scheduledActionName,omitempty"`

	// 최소용량
MinSize *int32 `json:"minSize"`

	// 최대용량
MaxSize *int32 `json:"maxSize"`

	// 기대용량
DesiredCapacity *int32 `json:"desiredCapacity,omitempty"`

	// 스케쥴시작일시
StartTime *string `json:"startTime,omitempty"`

	// 스케쥴종료일시
EndTime *string `json:"endTime,omitempty"`

	// 반복설정
Recurrence *string `json:"recurrence,omitempty"`

	// 타임존
TimeZone *string `json:"timeZone,omitempty"`
}
