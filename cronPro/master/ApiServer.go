package master

import (
	"goexamples/cronPro/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func handleJobSave(c *gin.Context) {

	var (
		inputs common.Job
	)

	err := c.BindWith(&inputs, binding.Form)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			common.BuildResponse(-1, err.Error(), nil))
		return
	}

	// 保存到etcd中
	oldJob, err := G_jobMgr.SaveJob(&inputs)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			common.BuildResponse(-1, err.Error(), nil))
		return
	}

	c.JSON(http.StatusInternalServerError,
		common.BuildResponse(0, "success", oldJob))
}

// 删除任务接口
// DELETE /job/delete  query : name=job1
func handleJobDelete(c *gin.Context) {

	name := c.DefaultQuery("name", "")

	oldJob, err := G_jobMgr.DeleteJob(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			common.BuildResponse(-1, err.Error(), nil))
		return
	}

	c.JSON(http.StatusInternalServerError,
		common.BuildResponse(0, "success", oldJob))
}

// 列举所有crontab任务
func handleJobList(c *gin.Context) {

	jobList, err := G_jobMgr.ListJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			common.BuildResponse(-1, err.Error(), nil))
		return
	}

	c.JSON(http.StatusInternalServerError,
		common.BuildResponse(0, "success", jobList))
}

// 强制杀死某个任务
// POST /job/kill  query: name=job1
func handleJobKill(c *gin.Context) {

	name := c.DefaultQuery("name", "")

	err := G_jobMgr.KillJob(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			common.BuildResponse(-1, err.Error(), nil))
		return
	}

	c.JSON(http.StatusInternalServerError,
		common.BuildResponse(0, "success", nil))

}

// 查询任务日志
func handleJobLog(c *gin.Context) {
	var (
		inputs common.LogList
	)

	err := c.BindQuery(&inputs)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			common.BuildResponse(-1, err.Error(), nil))
		return
	}

}

func InitApiServer() (err error) {

	router := gin.Default()

	router.POST("/job/save", handleJobSave)
	router.POST("/job/delete", handleJobDelete)
	router.POST("/job/list", handleJobList)
	router.POST("/job/kill", handleJobSave)
	router.POST("/job/log", handleJobSave)
	router.POST("/worker/list", handleJobSave)

	go router.Run(":" + strconv.Itoa(G_config.ApiPort))

	return
}
