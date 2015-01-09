package resources

import (
	"github.com/feedlabs/feedify"
	"github.com/abbot/go-http-auth"
	"github.com/astaxie/beego/context"
)

var (
	a *auth.DigestAuth
)

func Auth(ctx *context.Context) *Org {
	if GetAuthType()== "basic" {
		return AuthBasic(ctx)
	} else if GetAuthType() == "digest" {
		return AuthDigest(ctx)
	}
	return nil
}

func GetApiSecret() string {
	return feedify.GetConfigKey("api::secret")
}

func GetAuthType() string {
	return feedify.GetConfigKey("auth::type")
}

func GetAuthRealm() string {
	return feedify.GetConfigKey("auth::realm")
}

func SecretBasic(user, realm string) string {
	if user == "john" {
		token := "hello"
		return Crypt(token)
	}
	if user == "aabbccddeeffgghhiijjkk" {
		token := "x-oauth-basic"
		return Crypt(token)
	}
	if user == "chris" {
		token := "352735982359372598327958723957329"
		return Crypt(token)
	}
	return ""
}

func Crypt(password string) string {
	md5 := string(auth.MD5Crypt([]byte(password), []byte(""), []byte("$$")))
	return md5
}

func SecretDigest(user, realm string) string {
	if user == "john" {
		// password is "hello" and realm "localhost"
		return "121280a68cd55fc949b5b980d47a5718"
	}
	return ""
}

func AuthBasic(ctx *context.Context) *Org {
	authenticator := auth.NewBasicAuthenticator(GetAuthRealm(), SecretBasic)

	if username := authenticator.CheckAuth(ctx.Request); username == "" {
		authenticator.RequireAuth(ctx.ResponseWriter, ctx.Request)
	}

	return &Org{"0", "", "", 0, 0, 0}
}

func AuthDigest(ctx *context.Context) *Org {
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

	return &Org{"0", "", "", 0, 0, 0}
}

func GenerateChannelID() {
	// should generate proper ID
}
