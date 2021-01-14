package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const backupDir = "/tmp/sttp"

var baseDir = "."

func init() {
	if err := os.MkdirAll(backupDir, os.ModePerm); err != nil {
		panic(err)
	}
}

// Run the server
func Run(port int, baseDir string) {
	if strings.HasPrefix(baseDir, "/") {
		baseDir = baseDir
	} else if baseDir != "" {
		wd, _ := os.Getwd()
		baseDir = filepath.Join(wd, baseDir)
	}
	router := gin.Default()
	router.GET("/*real", GetHandler)
	router.POST("/*real", PostHandler)
	router.DELETE("/*real", DeleteHandler)
	router.Run(fmt.Sprintf(":%d", port))
}

// GetHandler serve download request
func GetHandler(c *gin.Context) {
	// TODO support browser
	dst := absPath(c.Param("real"))
	info, err := os.Stat(dst)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if info.IsDir() {
		// TODO
		c.Status(http.StatusNotImplemented)
		return
	}
	c.File(dst)
}

// PostHandler serve upload request
func PostHandler(c *gin.Context) {
	//if c.ContentType() == "application/octet-stream" {
	if c.ContentType() == "multipart/form-data" {
		// TODO
		c.Status(http.StatusNotImplemented)
		return
	}
	dst := absPath(c.Param("real"))
	err := writeFile(c.Request.Body, dst)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

// DeleteHandler serve download request
func DeleteHandler(c *gin.Context) {
	dst := absPath(c.Param("real"))
	_, err := os.Stat(dst)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = os.RemoveAll(dst)
	if err != nil {
		c.Status(http.StatusNotImplemented)
		return
	}
	c.Status(http.StatusOK)
}

func absPath(real string) string {
	return filepath.Join(baseDir, strings.TrimPrefix(real, "/"))
}

func writeFile(body io.ReadCloser, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}
	if _, err := os.Stat(dst); err == nil {
		err = os.Rename(dst, filepath.Join(backupDir, filepath.Base(dst)))
		return err
	} else if !os.IsNotExist(err) {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, body)
	return err
}
