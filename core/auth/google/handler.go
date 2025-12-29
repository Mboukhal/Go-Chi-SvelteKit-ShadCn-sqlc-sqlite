package googleoauth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Mboukhal/FactoryBase/core/auth"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)



func init() {
	// errenv := godotenv.Load()
	// if errenv != nil {
	// 	fmt.Println("Error loading .env file")
	// }
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectUrl :=  os.Getenv("APP_DOMAIN") + os.Getenv("GOOGLE_REDIRECT_URL")

	// println("Google OAuth Config:", clientId, redirectUrl)
	// println("Google OAuth Config Secret:", clientSecret)
	

	if clientId == "" || clientSecret == "" || redirectUrl == "" {
		panic("Missing GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, or GOOGLE_REDIRECT_URL environment variables")
	}


	googleOauthConfig = &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}


// RegisterRoutes sets up the OAuth routes on the given router.

func handleOAuthLogin(w http.ResponseWriter, r *http.Request) {
	// Implementation for OAuth login
	// url := googleOauthConfig.AuthCodeURL(oauthStateString,
	// 	oauth2.AccessTypeOffline,
	// 	oauth2.SetAuthURLParam("prompt", "consent"))

	url := googleOauthConfig.AuthCodeURL(oauthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleOAuthCallback(w http.ResponseWriter, r *http.Request) {
	// Implementation for OAuth callback
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	jsonData := []byte(content)

	var userInfo auth.UserInfo

	err = json.Unmarshal(jsonData, &userInfo)
	if err != nil {
		log.Println("Error unmarshaling user info:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// TODO: check if user exists in DB `profiles`, if not create user, else redirect to /sign-up?error=access-denied
	userInfo.Role, err = auth.GetUserFromDB(userInfo.Email) // default role assignment
	if err != nil {
		http.Redirect(w, r, "/sign-in?error=" + err.Error(), http.StatusTemporaryRedirect)
		return
	}

	tokenString, err := auth.CreateJWT(userInfo)
	if err != nil {
		// log.Println("Error creating JWT:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// println("Generated JWT:", os.Getenv("APP_TOKEN_ISSUER"), os.Getenv("APP_TOKEN_ISSUER"))


	fmt.Fprintf(w, `
		<script>
		localStorage.setItem("%s", "%s");
		window.location.href = "/";
		</script>
		`, os.Getenv("APP_TOKEN_ISSUER"), tokenString)
}


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		log.Println("Auth Header:", authHeader)
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// http.Error(w, "Invalid token format", http.StatusUnauthorized)
			// redirect to /sign-in
			http.Redirect(w, r, "/sign-in", http.StatusTemporaryRedirect)
			return
		}

		// Validate token
		jwtSecret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		ctx := context.WithValue(r.Context(), "user", token.Claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: implement revocation if needed

	// Clear the token from local storage via JavaScript
	fmt.Fprintf(w, `
		<script>
		localStorage.removeItem("%s");
		window.location.href = "/";
		</script>
		`, os.Getenv("APP_TOKEN_ISSUER"))
}
