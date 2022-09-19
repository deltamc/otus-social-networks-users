package middlewares

import (
	"github.com/deltamc/otus-social-networks-chat/models/users"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)
type handlerAuth func(w http.ResponseWriter, r *http.Request, user users.User)
