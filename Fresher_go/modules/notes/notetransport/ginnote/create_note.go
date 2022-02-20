package ginnote

import (
	"Fresher_go/common"
	"Fresher_go/component"
	"Fresher_go/modules/notes/notebiz"
	"Fresher_go/modules/notes/notemodel"
	"Fresher_go/modules/notes/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNote(AppCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data notemodel.NoteCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSONP(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := notestorage.NewSQLStore(AppCtx.GetMainDBConnection())
		biz := notebiz.NewCreateNoteBiz(store)
		if err := biz.Create_Note(c.Request.Context(), &data); err != nil {
			c.JSONP(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSONP(http.StatusOK, common.SimpleSuccesResponse(data))
	}
}
