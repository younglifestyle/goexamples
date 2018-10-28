package articles

import (
	"goexamples/gin-demo/dbops/db"
	"goexamples/gin-demo/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var dbs db.DBPool = db.Con()

// Create an article
func Create1(c *gin.Context) {
	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		return
	}

	err = dbs.DbFiles.Table(models.TableArticle).Create(article).Error
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}

// Edit an article
func Edit1(c *gin.Context) {
	article := models.Article{}

	err := dbs.DbFiles.Table(models.TableArticle).First(&article,
		"id = ?", c.Param("_id")).Error
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   "Edit article",
		"article": article,
	})
}

// List all articles
func List1(c *gin.Context) {
	articles := []models.Article{}

	err := dbs.DbFiles.Table(models.TableArticle).
		Order("updated_on desc").Find(&articles).Error
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title":    "Articles",
		"articles": articles,
	})
}

// Update an article
func Update1(c *gin.Context) {

	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		return
	}

	err = dbs.DbFiles.Table(models.TableArticle).Where("id=?", c.Param("_id")).
		Update(map[string]interface{}{
			"title":      article.Title,
			"body":       article.Body,
			"updated_on": time.Now().UnixNano() / int64(time.Millisecond),
		}).Error
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}

// Delete an article
func Delete1(c *gin.Context) {

	err := dbs.DbFiles.Table(models.TableArticle).Delete(nil,
		"id=?", c.Param("_id")).Error
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}
