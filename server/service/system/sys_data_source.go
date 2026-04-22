package system

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//@author: claude
//@function: CreateDataSource
//@description: 创建数据源
//@param: dataSource *model.SysDataSource
//@return: err error

type DataSourceService struct{}

var DataSourceServiceApp = new(DataSourceService)

func (dataSourceService *DataSourceService) CreateDataSource(dataSource *system.SysDataSource) error {
	return global.GVA_DB.Create(dataSource).Error
}

//@author: claude
//@function: UpdateDataSource
//@description: 更新数据源
//@param: dataSource *model.SysDataSource
//@return: err error

func (dataSourceService *DataSourceService) UpdateDataSource(dataSource *system.SysDataSource) error {
	return global.GVA_DB.Save(dataSource).Error
}

//@author: claude
//@function: DeleteDataSource
//@description: 删除数据源
//@param: id uint
//@return: err error

func (dataSourceService *DataSourceService) DeleteDataSource(id uint) error {
	return global.GVA_DB.Delete(&system.SysDataSource{}, id).Error
}

//@author: claude
//@function: GetDataSourceList
//@description: 获取数据源列表
//@param: info request.PageInfo
//@return: list []system.SysDataSource, total int64, err error

func (dataSourceService *DataSourceService) GetDataSourceList(info request.PageInfo) (list []system.SysDataSource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysDataSource{})

	if info.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return
}

//@author: claude
//@function: GetDataSourceById
//@description: 根据ID获取数据源
//@param: id uint
//@return: dataSource *system.SysDataSource, err error

func (dataSourceService *DataSourceService) GetDataSourceById(id uint) (dataSource *system.SysDataSource, err error) {
	dataSource = &system.SysDataSource{}
	err = global.GVA_DB.First(dataSource, id).Error
	return
}

// TestConnection 测试数据源连接
// @param dataSource *system.SysDataSource
// @return success bool, message string
func (dataSourceService *DataSourceService) TestConnection(dataSource *system.SysDataSource) (bool, string) {
	var dsn string
	var driver gorm.Dialector

	switch dataSource.Type {
	case "MySQL":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port, dataSource.Database)
		driver = mysql.Open(dsn)
	case "PostgreSQL":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			dataSource.Host, dataSource.Username, dataSource.Password, dataSource.Database, dataSource.Port)
		driver = postgres.Open(dsn)
	case "SQLServer":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port, dataSource.Database)
		driver = sqlserver.Open(dsn)
	case "达梦", "人大金仓", "Oracle":
		// 暂不支持，显示友好错误
		return false, "暂不支持 " + dataSource.Type + " 类型的连接测试"
	default:
		return false, "未知的数据源类型: " + dataSource.Type
	}

	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return false, "连接失败: " + err.Error()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return false, "获取数据库对象失败: " + err.Error()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return false, "Ping 失败: " + err.Error()
	}

	sqlDB.Close()
	return true, "连接成功"
}
