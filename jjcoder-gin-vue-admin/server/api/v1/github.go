package v1

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/devops-tools/github/resource"
	"strconv"
)

func NewRepo(c *gin.Context) {
	name := c.Query("name")
	private, _ := strconv.ParseBool(c.DefaultQuery("private", "true"))
	description := c.Query("description")

	repo := resource.NewRepo(&name, &private, &description)
	c.JSON(http.StatusOK, repo)
}

func ListRepos(c *gin.Context) {
	repos := resource.ListRepos()
	c.JSON(http.StatusOK, repos)
}

func verifySignature(c *gin.Context) (bool, error) {
	payloadBody, err := c.GetRawData()
	if err != nil {
		return false, err
	}
	hSignature := c.GetHeader("X-Hub-Signature")
	fmt.Printf("hSignature===%s", hSignature)
	signature := hmacSha1(payloadBody)
	return hSignature == signature, nil

}

func hmacSha1(payloadBody []byte) string {
	h := hmac.New(sha1.New, []byte("123456"))
	h.Write(payloadBody)
	fmt.Println("sha1=" + hex.EncodeToString(h.Sum(nil)))
	return "sha1=" + hex.EncodeToString(h.Sum(nil))
}

func CustomWebHooks(c *gin.Context) {
	if matched, _ := verifySignature(c); !matched {
		err := "Signatures didn't match!"
		c.String(http.StatusForbidden, err)
		fmt.Println(err)
		return
	}
	fmt.Println("Signatures is match! go!")
	c.String(http.StatusOK, "OK")
}
