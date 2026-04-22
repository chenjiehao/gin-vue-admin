package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DataSourceApi
// @Tags      DataSource
// @Summary   数据源管理
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @success   200   {object}  response.Response{data=string}  "创建基础api"

type DataSourceApi struct{}

var DataSourceApiApp = new(DataSourceApi)

// CreateDataSource
// @Tags      DataSource
// @Summary   创建数据源
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDataSource           true  "数据源信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /dataSource/createDataSource [post]
func (d *DataSourceApi) CreateDataSource(c *gin.Context) {
	var dataSource system.SysDataSource
	err := c.ShouldBindJSON(&dataSource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dataSourceService.CreateDataSource(&dataSource)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateDataSource
// @Tags      DataSource
// @Summary   更新数据源
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDataSource           true  "数据源信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /dataSource/updateDataSource [post]
func (d *DataSourceApi) UpdateDataSource(c *gin.Context) {
	var dataSource system.SysDataSource
	err := c.ShouldBindJSON(&dataSource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dataSourceService.UpdateDataSource(&dataSource)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteDataSource
// @Tags      DataSource
// @Summary   删除数据源
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                  true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /dataSource/deleteDataSource [post]
func (d *DataSourceApi) DeleteDataSource(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dataSourceService.DeleteDataSource(req.Uint())
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetDataSourceList
// @Tags      DataSource
// @Summary   获取数据源列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                 true  "分页参数"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /dataSource/getDataSourceList [post]
func (d *DataSourceApi) GetDataSourceList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dataSourceService.GetDataSourceList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetDataSourceById
// @Tags      DataSource
// @Summary   获取数据源详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                  true  "ID"
// @Success   200   {object}  response.Response{data=system.SysDataSource,msg=string}  "获取成功"
// @Router    /dataSource/getDataSourceById [post]
func (d *DataSourceApi) GetDataSourceById(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dataSource, err := dataSourceService.GetDataSourceById(req.Uint())
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(dataSource, "获取成功", c)
}

// TestConnection
// @Tags      DataSource
// @Summary   测试数据源连接
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce    application/json
// @Param     data  body      system.SysDataSource           true  "数据源信息"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "测试结果"
// @Router    /dataSource/testConnection [post]
func (d *DataSourceApi) TestConnection(c *gin.Context) {
	var dataSource system.SysDataSource
	err := c.ShouldBindJSON(&dataSource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	success, message := dataSourceService.TestConnection(&dataSource)
	if success {
		response.OkWithDetailed(map[string]interface{}{"success": true, "message": message}, "连接成功", c)
	} else {
		response.FailWithDetailed(map[string]interface{}{"success": false, "message": message}, message, c)
	}
}
