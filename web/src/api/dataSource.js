import service from '@/utils/request'

// @Tags DataSource
// @Summary 创建数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDataSource true "数据源信息"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"创建成功"}"
// @Router /dataSource/createDataSource [post]
export const createDataSource = (data) => {
  return service({
    url: '/dataSource/createDataSource',
    method: 'post',
    data
  })
}

// @Tags DataSource
// @Summary 更新数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDataSource true "数据源信息"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"更新成功"}"
// @Router /dataSource/updateDataSource [post]
export const updateDataSource = (data) => {
  return service({
    url: '/dataSource/updateDataSource',
    method: 'post',
    data
  })
}

// @Tags DataSource
// @Summary 删除数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ID"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"删除成功"}"
// @Router /dataSource/deleteDataSource [post]
export const deleteDataSource = (data) => {
  return service({
    url: '/dataSource/deleteDataSource',
    method: 'post',
    data
  })
}

// @Tags DataSource
// @Summary 获取数据源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /dataSource/getDataSourceList [post]
export const getDataSourceList = (data) => {
  return service({
    url: '/dataSource/getDataSourceList',
    method: 'post',
    data
  })
}

// @Tags DataSource
// @Summary 获取数据源详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ID"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /dataSource/getDataSourceById [post]
export const getDataSourceById = (data) => {
  return service({
    url: '/dataSource/getDataSourceById',
    method: 'post',
    data
  })
}

// @Tags DataSource
// @Summary 测试数据源连接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDataSource true "数据源信息"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"连接成功"}"
// @Router /dataSource/testConnection [post]
export const testConnection = (data) => {
  return service({
    url: '/dataSource/testConnection',
    method: 'post',
    data
  })
}

// @Tags DataSource
// @Summary 批量删除数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"删除成功"}"
// @Router /dataSource/batchDelete [post]
export const batchDelete = (data) => {
  return service({
    url: '/dataSource/batchDelete',
    method: 'post',
    data
  })
}

// @Tags DataSource
// @Summary 批量更新数据源状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BatchUpdateStatusReq true "ID列表和状态"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"更新成功"}"
// @Router /dataSource/batchUpdateStatus [post]
export const batchUpdateStatus = (data) => {
  return service({
    url: '/dataSource/batchUpdateStatus',
    method: 'post',
    data
  })
}

// @Tags DataSource
// @Summary 获取数据源表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param dataSourceId query string true "数据源ID"
// @Param database query string false "数据库名（不传则使用数据源配置的库）"
// @Success 200 {string} string "{"code":0,"data":{"tables":[]},"msg":"获取成功"}"
// @Router /dataSource/getTables [get]
export const getDataSourceTables = (dataSourceId, database) => {
  return service({
    url: '/dataSource/getTables',
    method: 'get',
    params: { dataSourceId, database }
  })
}
