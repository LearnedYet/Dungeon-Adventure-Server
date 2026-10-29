package auth

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func (a *Authenticator) HandleGuestAuth(w http.ResponseWriter, r *http.Request) {
	guestId := int(a.guestCounter.Add(1))
	email := "guest" + strconv.Itoa(guestId) + "@jrpg.com"
	nickname := "游客 " + strconv.Itoa(guestId)
	credentials := UserCredentials{Email: email, Nickname: nickname}
	status := a.handler.HandleGuestUserAuthenticated(credentials)
	if !status.IsAuthenticated {
		log.Info("游客认证：认证被拒绝！")
		http.Redirect(w, r, a.config.RedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	url := a.config.RedirectUrl +
		"/?token=" + string(status.Token) +
		"&isNewPlayer=" + strconv.FormatBool(status.IsNewPlayer) +
		"&isGuest=true"
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
