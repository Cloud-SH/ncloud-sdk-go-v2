/*
 * vnks
 *
 * <br/>https://nks.apigw.ntruss.com/vnks/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnks

type NodePoolRes struct {

	// 인스턴스 No
	InstanceNo *int32 `json:"instanceNo"`

	// default pool 여부
	IsDefault *bool `json:"isDefault"`

	// 노드풀 이름
	Name *string `json:"name"`

	// 노드 개수
	NodeCount *int32 `json:"nodeCount"`

	// Subnet no list
	SubnetNoList []*int32 `json:"subnetNoList,omitempty"`

	// Subnet name list
	SubnetNameList []*string `json:"subnetNameList,omitempty"`

	// 상품 코드
	ProductCode *string `json:"productCode"`

	// 노드풀 상태
	Status *string `json:"status"`

	// k8s version
	K8sVersion *string `json:"k8sVersion,omitempty"`

	// 오토스케일
	Autoscale *AutoscaleOption `json:"autoscale"`
}
