package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

const badGinVersion = "v1.8.2"

const leakGitHubToken = "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"

const leakSSHKey = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB03zpHzR9bWQHUx2NACvhr1sbMdT+sEz8+1jELnKmIWgAAAJCnW3xep1t8
XgAAAAtzc2gtZWQyNTUxOQAAACB03zpHzR9bWQHUx2NACvhr1sbMdT+sEz8+1jELnKmIWg
AAAEAcfx3f66PksfHpaVQLpo3VqGjLfMmSzHikNKdPoUz3M3TfOkfNH1tZAdTHY0AK+GvW
xsx1P6wTPz7WMQucqYhaAAAACWxlYWsgdGVzdAECAwQ=
-----END OPENSSH PRIVATE KEY-----
`

func main() {
	gin.SetMode(gin.DebugMode)
	if gin.Version != badGinVersion {
		log.Fatalf("Current Gin Version: %s, Want: %s - This is deliberately a vulnerable package version for security scanners to detect. Please make sure your go.mod has not been modified or Github and pushed a 'fix' to this repo.", gin.Version, badGinVersion)
	}
	log.Println("This is a EXTREMELY vulnerable application and container. Please DO NOT run in a production environment")
	log.Println("Starting...")
	r := gin.Default()
	r.GET("/download", func(c *gin.Context) {
		dir := "/download/"
		filename := c.Params.ByName("filename") // eg = "malicious.sh\";dummy=.txt"
		if filename == "" {
			c.AbortWithStatus(400)
			return
		}
		//CVE-2023-29401
		c.FileAttachment(dir+filename, filename)
	})
	r.GET("/hash", func(c *gin.Context) {
		data := c.Params.ByName("data")
		if data == "" {
			c.AbortWithStatus(400)
			return
		}
		//MD5 should not *really* be used
		c.String(200, fmt.Sprintf("%x", md5.Sum([]byte(data))))
	})
	r.GET("/leak", func(c *gin.Context) {
		s := new(strings.Builder)
		//leak sensitive information
		s.WriteString(leakGitHubToken)
		s.WriteString(leakSSHKey)
		c.String(200, s.String())
	})
	err := r.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
