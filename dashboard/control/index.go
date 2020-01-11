package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {

}


func (ctl *Index) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "success")
}
