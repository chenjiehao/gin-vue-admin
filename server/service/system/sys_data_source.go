package system

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
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
//@author: claude
//@function: TestConnection
//@description: 测试数据源连接
//@param: dataSource *system.SysDataSource
//@return: success bool, message string
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

// BatchDelete 批量删除数据源
//@author: claude
//@function: BatchDelete
//@description: 批量删除数据源
//@param: ids []uint
//@return: successCount int, failCount int, err error
func (dataSourceService *DataSourceService) BatchDelete(ids []uint) (successCount int, failCount int, err error) {
	if len(ids) == 0 {
		return 0, 0, nil
	}

	tx := global.GVA_DB.Delete(&system.SysDataSource{}, "id IN ?", ids)
	if tx.Error != nil {
		return 0, len(ids), tx.Error
	}

	successCount = int(tx.RowsAffected)
	failCount = len(ids) - successCount
	return successCount, failCount, nil
}

// BatchUpdateStatus 批量更新数据源状态
//@author: claude
//@function: BatchUpdateStatus
//@description: 批量更新数据源状态
//@param: ids []uint, status uint
//@return: successCount int, failCount int, err error
func (dataSourceService *DataSourceService) BatchUpdateStatus(ids []uint, status uint) (successCount int, failCount int, err error) {
	if len(ids) == 0 {
		return 0, 0, nil
	}

	tx := global.GVA_DB.Model(&system.SysDataSource{}).Where("id IN ?", ids).Update("status", status)
	if tx.Error != nil {
		return 0, len(ids), tx.Error
	}

	successCount = int(tx.RowsAffected)
	failCount = len(ids) - successCount
	return successCount, failCount, nil
}

// GetDatabases 获取数据源实例下的所有数据库
//@author: claude
//@function: GetDatabases
//@description: 获取数据源实例下的所有数据库
//@param: dataSourceId uint
//@return: databases []string, err error
func (dataSourceService *DataSourceService) GetDatabases(dataSourceId uint) ([]string, error) {
	// 1. 根据 dataSourceId 查询数据源配置
	var dataSource system.SysDataSource
	if err := global.GVA_DB.First(&dataSource, dataSourceId).Error; err != nil {
		return nil, errors.New("数据源不存在")
	}

	// 2. 根据 type 创建对应的 GORM Dialector（不指定 database）
	var dsn string
	var driver gorm.Dialector

	switch dataSource.Type {
	case "MySQL":
		// MySQL 连接时不指定数据库，查询所有数据库
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
			dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
		driver = mysql.Open(dsn)
	case "PostgreSQL":
		// PostgreSQL 连接时不指定数据库，查询所有数据库
		dsn = fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=disable",
			dataSource.Host, dataSource.Username, dataSource.Password, dataSource.Port)
		driver = postgres.Open(dsn)
	case "SQLServer":
		// SQLServer 需要指定数据库来查询，但可以连接 master 库
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=master",
			dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
		driver = sqlserver.Open(dsn)
	case "达梦", "人大金仓", "Oracle":
		return nil, errors.New("暂不支持 " + dataSource.Type + " 类型的数据库查询")
	default:
		return nil, errors.New("未知的数据源类型: " + dataSource.Type)
	}

	// 3. 建立连接
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return nil, errors.New("连接失败: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("获取数据库对象失败: " + err.Error())
	}
	defer sqlDB.Close()

	// 4. 根据数据库类型执行查询数据库列表的 SQL
	var databases []string
	switch dataSource.Type {
	case "MySQL":
		// 使用 information_schema 查询数据库列表，这种方式权限要求更低
		rows, err := sqlDB.Query("SELECT SCHEMA_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME NOT IN ('information_schema', 'performance_schema', 'mysql', 'sys', 'ndb', 'tmp')")
		if err != nil {
			return nil, errors.New("查询数据库列表失败: " + err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var dbName string
			if err := rows.Scan(&dbName); err != nil {
				return nil, err
			}
			databases = append(databases, dbName)
		}
	case "PostgreSQL":
		rows, err := sqlDB.Query("SELECT datname FROM pg_database WHERE datistemplate = false")
		if err != nil {
			return nil, errors.New("查询数据库列表失败: " + err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var dbName string
			if err := rows.Scan(&dbName); err != nil {
				return nil, err
			}
			// 过滤掉系统数据库
			if dbName != "postgres" && dbName != "template0" && dbName != "template1" {
				databases = append(databases, dbName)
			}
		}
	case "SQLServer":
		rows, err := sqlDB.Query("SELECT name FROM sys.databases WHERE state = 0")
		if err != nil {
			return nil, errors.New("查询数据库列表失败: " + err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var dbName string
			if err := rows.Scan(&dbName); err != nil {
				return nil, err
			}
			// 过滤掉系统数据库
			if dbName != "master" && dbName != "tempdb" && dbName != "model" && dbName != "msdb" {
				databases = append(databases, dbName)
			}
		}
	}

	return databases, nil
}

// GetTables 获取数据源下的表列表
//@author: claude
//@function: GetTables
//@description: 获取数据源下的表列表
//@param: dataSourceId uint, database string
//@return: tables []string, err error
func (dataSourceService *DataSourceService) GetTables(dataSourceId uint, database string) ([]string, error) {
	// 1. 根据 dataSourceId 查询数据源配置
	var dataSource system.SysDataSource
	if err := global.GVA_DB.First(&dataSource, dataSourceId).Error; err != nil {
		return nil, errors.New("数据源不存在")
	}

	// 2. 如果未指定 database，使用数据源配置的 database
	if database == "" {
		database = dataSource.Database
	}

	// 3. 根据 type 创建对应的 GORM Dialector
	var dsn string
	var driver gorm.Dialector

	switch dataSource.Type {
	case "MySQL":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port, database)
		driver = mysql.Open(dsn)
	case "PostgreSQL":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			dataSource.Host, dataSource.Username, dataSource.Password, database, dataSource.Port)
		driver = postgres.Open(dsn)
	case "SQLServer":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port, database)
		driver = sqlserver.Open(dsn)
	case "达梦", "人大金仓", "Oracle":
		return nil, errors.New("暂不支持 " + dataSource.Type + " 类型的表查询")
	default:
		return nil, errors.New("未知的数据源类型: " + dataSource.Type)
	}

	// 3. 建立连接
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return nil, errors.New("连接失败: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("获取数据库对象失败: " + err.Error())
	}
	defer sqlDB.Close()

	// 4. 根据数据库类型执行查询表的 SQL
	var tables []string
	switch dataSource.Type {
	case "MySQL":
		err = sqlDB.QueryRow("SHOW TABLES").Scan(&tables)
		// SHOW TABLES 返回的结果需要用不同方式解析
		rows, err := sqlDB.Query("SHOW TABLES")
		if err != nil {
			return nil, errors.New("查询表失败: " + err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				return nil, err
			}
			tables = append(tables, tableName)
		}
	case "PostgreSQL":
		rows, err := sqlDB.Query("SELECT tablename FROM pg_tables WHERE schemaname = 'public'")
		if err != nil {
			return nil, errors.New("查询表失败: " + err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				return nil, err
			}
			tables = append(tables, tableName)
		}
	case "SQLServer":
		rows, err := sqlDB.Query("SELECT name FROM sys.tables")
		if err != nil {
			return nil, errors.New("查询表失败: " + err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				return nil, err
			}
			tables = append(tables, tableName)
		}
	}

	return tables, nil
}
