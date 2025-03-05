package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

type Email struct {
	to      string
	subject string
	message string
}

func sendEmail(to, subject, body string) error {
	// CONFIGURATION FOR MAIL
	m := gomail.NewMessage()
	m.SetHeader("From", "avishshiroya2034@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer("smtp.gmail.com", 587, "avishshiroya2034@gmail.com", "nkvd jggd doyf wfcq")

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	fmt.Println("Email Send Successfully to ", to)
	return nil
}

func worker(id int, jobs <-chan Email, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("email will send", id)
	for email := range jobs {
		res := sendEmail(email.to, email.subject, email.message)
		if res == nil {
			fmt.Println("Response after send", res)
		}
	}
}

func main() {
	app := fiber.New()

	app.Get("/sendmail", func(c *fiber.Ctx) error {
		numWorkers := 3
		emailQueue := make(chan Email, 100)
		var wg sync.WaitGroup
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go worker(i, emailQueue, &wg)
		}

		emails := []Email{
			{"avish@yopmail.com", "Welcome!", "Hello User 1"},
			{"avish@yopmail.com", "Update", "Your account has been updated"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Discount", "Get 20% off today!"},
			{"avish@yopmail.com", "Reminder", "Your subscription is expiring soon"},
		}
		for _, email := range emails {
			emailQueue <- email
		}
		close(emailQueue)
		wg.Wait()
		fmt.Println("All Work Done")
		return c.JSON(fiber.Map{
			"message": "Email Seded",
		})
	})

	app.Listen(":8080")
}
