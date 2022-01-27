/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type ExportBackupToObjectStorageRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// 백업파일이름
	FileName *string `json:"fileName"`

	// ObjectStorage버킷이름
	BucketName *string `json:"bucketName"`

	// CloudMysql인스턴스번호
	CloudMysqlInstanceNo *string `json:"cloudMysqlInstanceNo"`
}
