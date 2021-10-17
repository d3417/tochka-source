package marketplace

import (
	"github.com/gocraft/web"

	"github.com/d3417/tochka-source/modules/util"
)

func (c *Context) AdminReviews(w web.ResponseWriter, r *web.Request) {
	// c.Reviews = GetAllReviews()
	util.RenderTemplate(w, "reviews/admin/reviews", c)
}
