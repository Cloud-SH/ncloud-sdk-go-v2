/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-10-18T06:18:35Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type CdnPlusPurgeHistory struct {

	// CDN인스턴스번호
CdnInstanceNo *string `json:"cdnInstanceNo,omitempty"`

	// 퍼지ID
PurgeId *string `json:"purgeId,omitempty"`

	// 전체퍼지여부
IsWholePurge *bool `json:"isWholePurge,omitempty"`

	// 전체도메인퍼지여부
IsWholeDomain *bool `json:"isWholeDomain,omitempty"`

	// CDN+서비스도메인리스트
CdnPlusServiceDomainList []*CdnPlusServiceDomain `json:"cdnPlusServiceDomainList,omitempty"`

	// 대상디렉토리명
TargetDirectoryName *string `json:"targetDirectoryName,omitempty"`

	// 타겟파일리스트
TargetFileList []*string `json:"targetFileList,omitempty"`

	// 요청날짜
RequestDate *string `json:"requestDate,omitempty"`

	// 퍼지상태
PurgeStatusName *string `json:"purgeStatusName,omitempty"`
}
