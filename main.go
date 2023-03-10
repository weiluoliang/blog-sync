package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	r := gin.Default()
	r.GET("/sync", func(c *gin.Context) {
		// 去同步我的文件
		go func() {

			// 删除目录
			rm := exec.Command("rm", "-rf", "/usr/share/nginx/weiluoliang.github.io")
			output, err2 := rm.CombinedOutput()
			fmt.Printf("delete result :\n%s\n", string(output))
			if err2 != nil {
				log.Printf("sync failed with %s\n", err2)
			}

			// check 指定分支
			//  git clone -b gh-page https://github.com/weiluoliang/weiluoliang.github.io.git
			cmd := exec.Command("git", "clone", "-b", "gh-page", "https://github.com/weiluoliang/weiluoliang.github.io.git")
			cmd.Dir = "/usr/share/nginx"
			out, err := cmd.CombinedOutput()
			fmt.Printf("sync result :\n%s\n", string(out))
			if err != nil {
				log.Printf("sync failed with %s\n", err)
			}
		}()

		c.JSON(http.StatusOK, gin.H{
			"message": "repo sync OK",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
