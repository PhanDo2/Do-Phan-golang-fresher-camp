package ginnote

import (
	"Fresher_go/common"
	"Fresher_go/component"
	"Fresher_go/modules/notes/notebiz"
	"Fresher_go/modules/notes/notemodel"
	"Fresher_go/modules/notes/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteNote(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSONP(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data notemodel.NoteUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSONP(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewDeleteNoteBiz(store)

		if err := biz.DeleteNote(c.Request.Context(), id); err != nil {
			c.JSONP(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse("delete finish "))
	}
}
