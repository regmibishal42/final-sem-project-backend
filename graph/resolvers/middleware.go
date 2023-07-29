package resolvers

import (
	"backend/graph/model"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
)

var UserCtxKey = &contextKey{"user"}
var DomainCtxKey = &contextKey{"domain"}

type Domain string
type contextKey struct {
	id string
}

func UserForContext(ctx context.Context) *model.User {
	user := ctx.Value(UserCtxKey).(*model.User)
	return user
}

func DomainForContext(ctx context.Context) Domain {
	domain := ctx.Value(DomainCtxKey).(Domain)
	return domain
}

//check if user is logged in
func CheckLoggedIn(user *model.User) error {
	if user == nil {
		return errors.New("need authentication")
	}
	return nil
}

func Middleware(next http.Handler, resolver *Resolver) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var domain = Domain("https://" + r.Host)
		ctx := context.WithValue(r.Context(), DomainCtxKey, domain)
		r = r.WithContext(ctx)

		var user *model.User
		if AuthHeader := r.Header.Get("Authorization"); AuthHeader != "" {
			token := strings.TrimSpace(strings.TrimPrefix(AuthHeader, "Bearer"))
			userID, err := model.ParseAuthToken(token)
			if err != nil {
				http.Error(w, "Invalid Token", http.StatusUnauthorized)
				log.Printf("Error In Middleware %s", err.Error())
				return
			}
			user, err = resolver.AuthDomain.GetUserByID(r.Context(), userID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
			if user != nil {
				if !user.IsVerified {
					http.Error(w, errors.New("User is not Verified").Error(), http.StatusNotAcceptable)
				}
			}
		}
		ctx = context.WithValue(r.Context(), UserCtxKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
