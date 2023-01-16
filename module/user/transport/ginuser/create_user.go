package ginuser

import (
	userbiz "be-music/module/user/biz"
	"be-music/module/user/model"
	storageuser "be-music/module/user/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println("this is data", data)
		store := storageuser.NewSQLStore(db)

		biz := userbiz.NewCreateUserBiz(store)
		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
