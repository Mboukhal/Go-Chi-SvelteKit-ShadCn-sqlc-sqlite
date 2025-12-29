package magiclink

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mboukhal/FactoryBase/internal/adapter/email"
	"github.com/Mboukhal/FactoryBase/internal/settings"
)

const magicLinkExpiryMinutes = 3

func sendMagicLink(ctx context.Context, user_email string) error {
	// find if email exists in db loginpl
	queries := settings.GetQueries(ctx)
	if queries == nil {
		return errors.New("database queries not available")
	}

	token, _ := queries.GetTokenByEmail(ctx, user_email)
	if token.CreatedAt.Valid && token.CreatedAt.Time.Add(magicLinkExpiryMinutes * time.Minute).Before(time.Now()) && token.CounterRequest <= 3 {
		queries.UpdateTokenCount(ctx, user_email)
		return errors.New("Login link already sent recently, please check your email")
	} else if token.CreatedAt.Valid {
		log.Println("Token count:", token.CounterRequest)	
		if token.CounterRequest < 3 {
				// resend the same token
				queries.UpdateTokenCount(ctx, 	user_email)
				log.Println("Resending token email to:", user_email, "with token:", token.ID)
				sendTokenEmail(user_email, token.ID)
				return errors.New("Login link resent, please check your email")
			}
				remaining := time.Until(token.CreatedAt.Time.Add(magicLinkExpiryMinutes * time.Minute))
				if remaining < 0 {
					queries.DeleteTokenByEmail(ctx, user_email)
			} else {

				mins := int(remaining.Minutes())
				secs := int(remaining.Seconds()) % 60
				return errors.New("Request limit not reached, please wait for: " + fmt.Sprintf("%02dm%02ds", mins, secs) + " before requesting a new link")
		}
	}

	user_id, err := queries.CreateToken(ctx, user_email)
	if err != nil {
		return err
	}

	// start cron job to delete token after 10 minutes
	go func(e string) {
		time.Sleep(magicLinkExpiryMinutes * time.Minute)
		queries.DeleteTokenByEmail(context.Background(), e)
		log.Println("Deleted token for email:", e)
	}(user_email)

	log.Println("Sending token email to:", user_email, "with token:", user_id)

	// send email with login link containing the token
	sendTokenEmail(user_email,  user_id)

	
	return nil
	
}

func sendTokenEmail(to string, token string) error {
	// implement email sending logic here
	auth_link := os.Getenv("APP_DOMAIN") + "/auth?token=" + token

	log.Println("Sending token email to:", to, "with token:", auth_link)
	err := email.SendEmailSys(to, "Your Login Link", "Click the following link to log in: \n" + auth_link)

	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Println("Email sent successfully to:", to)
	
	return nil
}

func checkEmailAuthorization(ctx context.Context, email string) error {
	
	// log.Println("Checking email authorization for:", email)
	// check if email has profile in db
		// Get queries from context
	queries := settings.GetQueries(ctx)
	// log.Println("Queries in login link handler:", queries)
	if queries == nil {
		return http.ErrServerClosed
	}

	
	count, err := queries.CheckUserEmailExists(ctx, email)
	// log.Println("Email authorization check count:", count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("email not authorized")
	}
	return nil
}