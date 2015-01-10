package helper

import (
	auth "github.com/abbot/go-http-auth"
	"github.com/astaxie/beego/context"
	"github.com/feedlabs/api/resource"
)

var (
	a *auth.DigestAuth
)

func Auth(ctx *context.Context) *resource.Admin {
	if GetAuthType() == "basic" {
		return AuthBasic(ctx)
	} else if GetAuthType() == "digest" {
		return AuthDigest(ctx)
	}
	return nil
}

func SecretBasic(user, realm string) string {
	if user == "john" {
		token := "hello"
		return GetCrypt(token)
	}
	return ""
}

func SecretDigest(user, realm string) string {
	if user == "john" {
		token := "hello"
		return GetMd5(user + ":" + realm + ":" + token)
	}
	return ""
}

func AuthBasic(ctx *context.Context) *resource.Admin {
	authenticator := auth.NewBasicAuthenticator(GetAuthRealm(), SecretBasic)

	if username := authenticator.CheckAuth(ctx.Request); username == "" {
		authenticator.RequireAuth(ctx.ResponseWriter, ctx.Request)
	}

	return &resource.Admin{"0", &resource.Org{"0", "", "", 0, 0, 0}, "", 0}
}

func AuthDigest(ctx *context.Context) *resource.Admin {
	if a == nil {
		a = auth.NewDigestAuthenticator(GetAuthRealm(), SecretDigest)
	}

	if username, authinfo := a.CheckAuth(ctx.Request); username == "" {
		a.RequireAuth(ctx.ResponseWriter, ctx.Request)
	} else {
		ar := &auth.AuthenticatedRequest{Request: *ctx.Request, Username: username}
		if authinfo != nil {
			ctx.ResponseWriter.Header().Set("Authentication-Info", *authinfo)
		}

		if ar.Username == "" {
			return nil
		}
	}

	return &resource.Admin{"0", &resource.Org{"0", "", "", 0, 0, 0}, "", 0}
}
