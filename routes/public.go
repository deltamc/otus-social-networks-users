package routes

import (
	c "github.com/deltamc/otus-social-networks-chat/controllers"
	m "github.com/deltamc/otus-social-networks-chat/middlewares"
	"net/http"
)

func Public() {
	http.HandleFunc("/sign-up", m.Cors(m.Post(c.HandleSignUp)))
	http.HandleFunc("/sign-in", m.Cors(m.Post(c.HandleSignIn)))
	http.HandleFunc("/users", m.Cors(m.Get(c.HandleUsers)))
}
