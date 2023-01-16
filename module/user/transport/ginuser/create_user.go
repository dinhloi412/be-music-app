package ginuser

import (
	"be-music/common"

	userbiz "be-music/module/user/biz"
	"be-music/module/user/model"
	storageuser "be-music/module/user/storage"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(appctx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetDBConnection()
		var data model.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		fmt.Println("this is data", data)
		store := storageuser.NewSQLStore(db)

		biz := userbiz.NewCreateUserBiz(store)
		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
