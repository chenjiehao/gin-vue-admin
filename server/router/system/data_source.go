package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataSourceRouter struct{}

func (d *DataSourceRouter) InitDataSourceRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	dataSourceRouter := Router.Group("dataSource").Use(middleware.OperationRecord())
	{
		dataSourceRouter.POST("createDataSource", dataSourceApi.CreateDataSource)
		dataSourceRouter.POST("updateDataSource", dataSourceApi.UpdateDataSource)
		dataSourceRouter.POST("deleteDataSource", dataSourceApi.DeleteDataSource)
		dataSourceRouter.POST("getDataSourceList", dataSourceApi.GetDataSourceList)
		dataSourceRouter.POST("getDataSourceById", dataSourceApi.GetDataSourceById)
		dataSourceRouter.POST("testConnection", dataSourceApi.TestConnection)
		dataSourceRouter.POST("batchDelete", dataSourceApi.BatchDelete)
		dataSourceRouter.POST("batchUpdateStatus", dataSourceApi.BatchUpdateStatus)
		dataSourceRouter.GET("getTables", dataSourceApi.GetTables)
		dataSourceRouter.GET("getDatabases", dataSourceApi.GetDatabases)
	}
	return dataSourceRouter
}
