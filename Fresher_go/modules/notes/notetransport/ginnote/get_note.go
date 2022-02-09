package ginnote

import (
	"Fresher_go/common"
	"Fresher_go/component"
	"Fresher_go/modules/notes/notebiz"
	"Fresher_go/modules/notes/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetNote(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewGetNoteBiz(store)

		data, err := biz.GetNote(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(data))
	}
}
