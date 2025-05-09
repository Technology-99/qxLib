package qxDatabase

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type UniqueIdAdminModel struct {
	CreatedBy string `gorm:"index:idx_created_by;column:created_by;comment:创建者;type: varchar(255)" json:"created_by"`
	UpdatedBy string `gorm:"index:idx_updated_by;column:updated_by;comment:更新者;type: varchar(255)" json:"updated_by"`
	DeletedBy string `gorm:"index:idx_deleted_by;column:deleted_by;comment:删除者;type: varchar(255)" json:"-"`
}

type BaseAdminModel struct {
	CreatedBy uint `gorm:"column:created_by;comment:创建者;type: int" json:"created_by"`
	UpdatedBy uint `gorm:"column:updated_by;comment:更新者;type: int" json:"updated_by"`
	DeletedBy uint `gorm:"column:deleted_by;comment:删除者;type: int" json:"-"`
}

type BaseTenantModel struct {
	CreatedTenantBy string `gorm:"index:idx_created_tenant_by;column:created_tenant_by;comment:创建数据的租户;type: varchar(255)" json:"created_tenant_by"`
	UpdatedTenantBy string `gorm:"index:idx_updated_tenant_by;column:updated_tenant_by;comment:更新数据的租户;type: varchar(255)" json:"updated_tenant_by"`
	DeletedTenantBy string `gorm:"index:idx_deleted_tenant_by;column:deleted_tenant_by;comment:删除数据的租户;type: varchar(255)" json:"-"`
}
