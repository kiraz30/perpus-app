package cmd

import (
	"log"
	"net/http"
	"perpus-app/helpers"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		log.Println("Authorization header is missing")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	if d.UserRepository == nil {
		log.Println("UserRepository is not initialized")
		helpers.SendResponseHTTP(ctx, http.StatusInternalServerError, "Internal Server Error", nil)
		ctx.Abort()
		return
	}

	_, err := d.UserRepository.GetUserSessionToken(ctx.Request.Context(), token)
	if err != nil {
		log.Println("Error GET user session token:", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized failed get token", nil)
		ctx.Abort()
		return
	}

	// if dataUserToken.Token == "" {
	// 	log.Println("User session token not found")
	// 	helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized token not found", nil)
	// 	ctx.Abort()
	// 	return
	// }

	claim, err := helpers.ValidateToken(ctx, token)
	if err != nil {
		log.Println("Error validating token:", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized validate token", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("Token has expired", claim.ExpiresAt)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized token expired", nil)
		return
	}
	ctx.Set("token", claim)
	ctx.Next()
	return

}
