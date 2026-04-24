package system

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	offlineReq "github.com/flipped-aurora/gin-vue-admin/server/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// OfflineSyncApi 离线同步任务API
type OfflineSyncApi struct{}

// CreateOfflineSync 创建离线同步任务
// @Summary   创建离线同步任务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CreateOfflineSyncRequest  true  "创建离线同步任务请求"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建成功"
// @Router    /offlineSync/createOfflineSync [post]
func (o *OfflineSyncApi) CreateOfflineSync(c *gin.Context) {
	var req offlineReq.CreateOfflineSyncRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	offlineSync, err := service.ServiceGroupApp.OfflineSyncService.CreateOfflineSync(req)
	if err != nil {
		global.GVA_LOG.Error("创建离线同步任务失败!", zap.Error(err))
		response.FailWithDetailed(map[string]interface{}{}, "创建失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"record": offlineSync}, "创建成功", c)
}

// UpdateOfflineSync 更新离线同步任务
// @Summary   更新离线同步任务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.UpdateOfflineSyncRequest  true  "更新离线同步任务请求"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /offlineSync/updateOfflineSync [put]
func (o *OfflineSyncApi) UpdateOfflineSync(c *gin.Context) {
	var req offlineReq.UpdateOfflineSyncRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.ServiceGroupApp.OfflineSyncService.UpdateOfflineSync(req); err != nil {
		global.GVA_LOG.Error("更新离线同步任务失败!", zap.Error(err))
		response.FailWithDetailed(map[string]interface{}{}, "更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteOfflineSync 删除离线同步任务
// @Summary   删除离线同步任务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id   path      int  true  "任务ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /offlineSync/deleteOfflineSync [delete]
func (o *OfflineSyncApi) DeleteOfflineSync(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.FailWithMessage("缺少任务ID", c)
		return
	}

	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		response.FailWithMessage("无效的任务ID", c)
		return
	}

	if err := service.ServiceGroupApp.OfflineSyncService.DeleteOfflineSync(id); err != nil {
		global.GVA_LOG.Error("删除离线同步任务失败!", zap.Error(err))
		response.FailWithDetailed(map[string]interface{}{}, "删除失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetOfflineSyncById 获取离线同步任务详情
// @Summary   获取离线同步任务详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id   path      int  true  "任务ID"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "获取成功"
// @Router    /offlineSync/getOfflineSyncById [get]
func (o *OfflineSyncApi) GetOfflineSyncById(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.FailWithMessage("缺少任务ID", c)
		return
	}

	var req struct {
		ID uint `json:"id"`
	}
	if _, err := fmt.Sscanf(idStr, "%d", &req.ID); err != nil {
		response.FailWithMessage("无效的任务ID", c)
		return
	}

	offlineSync, err := service.ServiceGroupApp.OfflineSyncService.GetOfflineSyncById(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取离线同步任务失败!", zap.Error(err))
		response.FailWithDetailed(map[string]interface{}{}, "获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"record": offlineSync}, "获取成功", c)
}

// GetOfflineSyncList 获取离线同步任务列表
// @Summary   获取离线同步任务列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo  true  "分页查询请求"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "获取成功"
// @Router    /offlineSync/getOfflineSyncList [post]
func (o *OfflineSyncApi) GetOfflineSyncList(c *gin.Context) {
	var pageInfo request.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := service.ServiceGroupApp.OfflineSyncService.GetOfflineSyncList(pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		global.GVA_LOG.Error("获取离线同步任务列表失败!", zap.Error(err))
		response.FailWithDetailed(map[string]interface{}{}, "获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"list":  list,
		"total": total,
	}, "获取成功", c)
}

// RegenerateScript 重新生成脚本
// @Summary   重新生成 SeaTunnel 脚本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id   path      int  true  "任务ID"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "生成成功"
// @Router    /offlineSync/regenerateScript [post]
func (o *OfflineSyncApi) RegenerateScript(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	script, err := service.ServiceGroupApp.OfflineSyncService.RegenerateScript(req.ID)
	if err != nil {
		global.GVA_LOG.Error("重新生成脚本失败!", zap.Error(err))
		response.FailWithDetailed(map[string]interface{}{}, "生成失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"script": script}, "生成成功", c)
}

// GetScript 获取脚本内容
// @Summary   获取 SeaTunnel 脚本内容
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id   path      int  true  "任务ID"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "获取成功"
// @Router    /offlineSync/getScript [get]
func (o *OfflineSyncApi) GetScript(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.FailWithMessage("缺少任务ID", c)
		return
	}

	var req struct {
		ID uint `json:"id"`
	}
	if _, err := fmt.Sscanf(idStr, "%d", &req.ID); err != nil {
		response.FailWithMessage("无效的任务ID", c)
		return
	}

	script, err := service.ServiceGroupApp.OfflineSyncService.GetScriptById(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取脚本失败!", zap.Error(err))
		response.FailWithDetailed(map[string]interface{}{}, "获取失败: "+err.Error(), c)
		return
	}

	// 解析返回的 JSON 字符串
	var scriptObj interface{}
	if err := json.Unmarshal([]byte(script), &scriptObj); err != nil {
		response.FailWithDetailed(map[string]interface{}{}, "解析脚本失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"script": scriptObj}, "获取成功", c)
}

// GenerateScript 生成 SeaTunnel 脚本（不保存，仅预览）
// @Summary   生成 SeaTunnel 脚本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      offlineReq.GenerateSeatunnelScriptRequest  true  "生成脚本请求"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "生成成功"
// @Router    /offlineSync/generateScript [post]
func (o *OfflineSyncApi) GenerateScript(c *gin.Context) {
	var req offlineReq.GenerateSeatunnelScriptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	script, err := service.ServiceGroupApp.SeatunnelScriptService.GenerateScript(req)
	if err != nil {
		global.GVA_LOG.Error("生成 SeaTunnel 脚本失败!", zap.Error(err))
		response.FailWithDetailed(map[string]interface{}{}, "生成失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"script": script}, "生成成功", c)
}