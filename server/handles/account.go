package handles

import (
	"strconv"

	"github.com/alist-org/alist/v3/internal/db"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/operations"
	"github.com/alist-org/alist/v3/server/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ListStorages(c *gin.Context) {
	var req common.PageReq
	if err := c.ShouldBind(&req); err != nil {
		common.ErrorResp(c, err, 400)
		return
	}
	log.Debugf("%+v", req)
	storages, total, err := db.GetStorages(req.PageIndex, req.PageSize)
	if err != nil {
		common.ErrorResp(c, err, 500)
		return
	}
	common.SuccessResp(c, common.PageResp{
		Content: storages,
		Total:   total,
	})
}

func CreateStorage(c *gin.Context) {
	var req model.Storage
	if err := c.ShouldBind(&req); err != nil {
		common.ErrorResp(c, err, 400)
		return
	}
	if err := operations.CreateStorage(c, req); err != nil {
		common.ErrorResp(c, err, 500, true)
	} else {
		common.SuccessResp(c)
	}
}

func UpdateStorage(c *gin.Context) {
	var req model.Storage
	if err := c.ShouldBind(&req); err != nil {
		common.ErrorResp(c, err, 400)
		return
	}
	if err := operations.UpdateStorage(c, req); err != nil {
		common.ErrorResp(c, err, 500, true)
	} else {
		common.SuccessResp(c)
	}
}

func DeleteStorage(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.ErrorResp(c, err, 400)
		return
	}
	if err := operations.DeleteStorageById(c, uint(id)); err != nil {
		common.ErrorResp(c, err, 500, true)
		return
	}
	common.SuccessResp(c)
}
