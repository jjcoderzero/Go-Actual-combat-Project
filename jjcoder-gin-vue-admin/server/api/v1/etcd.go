package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/devops-tools/etcd/resource"
)

func PutEtcd(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	putResponse := resource.PutCtx(key, value)
	c.JSON(http.StatusOK, putResponse)
}

func GetEtcd(c *gin.Context) {
	vars := make(map[string]string)
	key := c.Query("key")
	getResponse := resource.GetCtx(key)
	for _, ev := range getResponse.Kvs {
		vars[string(ev.Key)] = string(ev.Value)
	}
	c.JSON(http.StatusOK, vars)
}

func DeleteEtcd(c *gin.Context) {
	key := c.Query("key")
	deleteResponse := resource.DeleteCtx(key)
	c.JSON(http.StatusOK, deleteResponse)
}

func WatchEtcd(c *gin.Context) {
	key := c.Query("key")
	resource.WatchKey(key)
	c.String(http.StatusOK, "watching")
}
