import service from '@/utils/request'

// @Tags OfflineSync
// @Summary 分页获取同步任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取同步任务列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /offlineSync/getOfflineSyncList [get]
export const getOfflineSyncList = (params) => {
  return service({
    url: '/offlineSync/getOfflineSyncList',
    method: 'get',
    params
  })
}

// @Tags OfflineSync
// @Summary 创建同步任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateOfflineSyncRequest true "创建同步任务"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"创建成功"}"
// @Router /offlineSync/createOfflineSync [post]
export const createOfflineSync = (data) => {
  return service({
    url: '/offlineSync/createOfflineSync',
    method: 'post',
    data
  })
}

// @Tags OfflineSync
// @Summary 更新同步任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateOfflineSyncRequest true "更新同步任务"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"更新成功"}"
// @Router /offlineSync/updateOfflineSync [put]
export const updateOfflineSync = (data) => {
  return service({
    url: '/offlineSync/updateOfflineSync',
    method: 'put',
    data
  })
}

// @Tags OfflineSync
// @Summary 删除同步任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteOfflineSyncRequest true "删除同步任务"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"删除成功"}"
// @Router /offlineSync/deleteOfflineSync [delete]
export const deleteOfflineSync = (data) => {
  return service({
    url: '/offlineSync/deleteOfflineSync',
    method: 'delete',
    data
  })
}

// @Tags OfflineSync
// @Summary 手动触发同步执行
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TriggerSyncRequest true "手动触发同步执行"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"触发成功"}"
// @Router /offlineSync/triggerSync [post]
export const triggerSync = (data) => {
  return service({
    url: '/offlineSync/triggerSync',
    method: 'post',
    data
  })
}

// @Tags OfflineSync
// @Summary 获取同步执行历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetSyncHistoryRequest true "获取同步执行历史"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /offlineSync/getSyncHistory [get]
export const getSyncHistory = (params) => {
  return service({
    url: '/offlineSync/getSyncHistory',
    method: 'get',
    params
  })
}
