package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"uni-web/internal/infrastructure/security"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func (f *Form) Login(r *http.Request, w http.ResponseWriter, userId int) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	session.Values["user_id"] = userId
	_ = session.Save(r, w)
}

func (f *Form) LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/login.html"))
	session, _ := store.Get(r, "cookie-name")

	if auth, _ := session.Values["authenticated"].(bool); auth {
		http.Error(w, "Вы уже авторизированы!", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		data := make(map[string]string)
		login := strings.TrimSpace(r.PostFormValue("login"))
		password := strings.TrimSpace(r.PostFormValue("password"))
		user, err := f.userApp.GetUserByLogin(login)

		if err != nil {
			data["login"] = "Логин не найден"
			data["loginValue"] = login
		} else {
			if err = security.VerifyPassword(user.Password, password); err != nil {
				data["password"] = "Неправильный пароль"
				data["loginValue"] = login
				data["passwordValue"] = password
			}
		}

		if len(data) == 0 {
			data["successMsg"] = "Успех! Вы вошли!"
			f.Login(r, w, user.Id)
		}
		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
