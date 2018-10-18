/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-09-19T05:44:45Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type BackupFile struct {

	// 호스트이름
HostName *string `json:"hostName,omitempty"`

	// 파일이름
FileName *string `json:"fileName,omitempty"`

	// 데이터베이스이름
DatabaseName *string `json:"databaseName,omitempty"`

	// 시작LSN
FirstLsn *int64 `json:"firstLsn,omitempty"`

	// 마지막LSN
LastLsn *int64 `json:"LastLsn,omitempty"`

	// 백업유형
BackupType *CommonCode `json:"backupType,omitempty"`

	// 백업시작시간
BackupStartTime *string `json:"backupStartTime,omitempty"`

	// 백업종료시간
BackupEndTime *string `json:"backupEndTime,omitempty"`
}
