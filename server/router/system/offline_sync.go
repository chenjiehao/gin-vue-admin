package system

import (
	"github.com/gin-gonic/gin"
)

type OfflineSyncRouter struct{}

func (o *OfflineSyncRouter) InitOfflineSyncRouter(Router *gin.RouterGroup) {
	offlineSyncRouter := Router.Group("offlineSync")
	{
		offlineSyncRouter.POST("createOfflineSync", offlineSyncApi.CreateOfflineSync)
		offlineSyncRouter.PUT("updateOfflineSync", offlineSyncApi.UpdateOfflineSync)
		offlineSyncRouter.DELETE("deleteOfflineSync", offlineSyncApi.DeleteOfflineSync)
		offlineSyncRouter.GET("getOfflineSyncById", offlineSyncApi.GetOfflineSyncById)
		offlineSyncRouter.POST("getOfflineSyncList", offlineSyncApi.GetOfflineSyncList)
		offlineSyncRouter.POST("regenerateScript", offlineSyncApi.RegenerateScript)
		offlineSyncRouter.GET("getScript", offlineSyncApi.GetScript)
		offlineSyncRouter.POST("generateScript", offlineSyncApi.GenerateScript)
	}
}