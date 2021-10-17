package marketplace

import (
	"github.com/gocraft/web"

	"github.com/d3417/tochka-source/modules/util"
)

func (c *Context) ViewAdminListReferralPayments(w web.ResponseWriter, r *web.Request) {
	c.ReferralPayments = FindReferralPayments()
	util.RenderTemplate(w, "referral/admin/payments", c)
}
