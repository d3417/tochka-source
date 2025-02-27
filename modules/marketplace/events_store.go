package marketplace

import (
	"fmt"
	"strings"

	"github.com/d3417/tochka-source/modules/apis"
)

func EventNewTrustedStoreRequest(store Store) {
	var (
		marketUrl = MARKETPLACE_SETTINGS.SiteURL
		text      = fmt.Sprintf("[%s](%s/store/%s) has requested for a trusted store status",
			store.Storename, marketUrl, store.Storename,
		)
	)
	go apis.PostMattermostEvent(MARKETPLACE_SETTINGS.MattermostIncomingHookTrustedVendors, text)
}

func EventNewTrustedStoreVerificationThreadPost(user User, store Store, message Message) {
	var (
		marketUrl = MARKETPLACE_SETTINGS.SiteURL
		text      = fmt.Sprintf("[@%s](%s/user/%s) has posted in store verification thread [@%s](%s/staff/stores/%s):\n> %s",
			user.Username, marketUrl, user.Username,
			store.Storename, marketUrl, store.Storename,
			strings.Replace(message.Text, "\n", "\n > ", -1), //------------|
		)
	)
	go apis.PostMattermostEvent(MARKETPLACE_SETTINGS.MattermostIncomingHookTrustedVendors, text)
}
