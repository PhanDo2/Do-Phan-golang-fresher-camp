package ginnote

import (
	"Fresher_go/common"
	"Fresher_go/component"
	"Fresher_go/modules/notes/notebiz"
	"Fresher_go/modules/notes/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListNote(AppCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSONP(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := notestorage.NewSQLStore(AppCtx.GetMainDBConnection())

		biz := notebiz.NewListNoteBiz(store)
		result, err := biz.ListNote(c.Request.Context(), &paging)
		if err != nil {
			c.JSONP(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSONP(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
