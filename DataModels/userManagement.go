package DataModels

type UserDetails struct {
	ORG       string  `json:"org"                              gorm:"column:org; type:varchar(32)"`
	ID        string  `json:"id"                               gorm:"column:id; type:varchar(64);"`
	EmailID   string  `json:"email_id" validate:"required"     gorm:"column:email_id; type:varchar(64);" `
	EmpID     string  `json:"emp_id" validate:"required"       gorm:"column:emp_id; type:varchar(16);" `
	FirstName string  `json:"first_name" validate:"required"   gorm:"column:first_name; type:varchar(64);"`
	LastName  string  `json:"last_name" validate:"required"    gorm:"column:last_name; type:varchar(64);"`
	Password  string  `json:"password; omitempty" validate:"required"  gorm:"column:password; type:varchar(64);"`
	Role      string  `json:"role; omitempty" validate:"required"      gorm:"column:role; type:varchar(32);" `
	Status    bool    `json:"status" validate:"required"               gorm:"column:status"`
	UDF1      string  `json:"udf_1, omitempty"                 gorm:"column:udf_1;"`
	UDF2      int64   `json:"udf_2, omitempty"                 gorm:"column:udf_2;"`
	UDF3      float64 `json:"udf_3, omitempty"                 gorm:"column:udf_3;"`
	UDF4      bool    `json:"udf_4, omitempty"                 gorm:"column:udf_4;"`
}
