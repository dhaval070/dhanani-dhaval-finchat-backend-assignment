package main
import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/stripe/stripe-go"
    _ "finchat/dotenv"
    "finchat/customer"
)

func init() {
    stripe.Key = os.Getenv("STRIPE_SECRET")
}

func main() {
    app := fiber.New()

    app.Use(recover.New())

    app.Get("/", func (c *fiber.Ctx) error {
        return c.Send([]byte("hello"))
    })

    app.Post("/api/customer", createCustomer)
    app.Post("/api/payments", createPayment)
    app.Get("/api/payments/:customerId", paymentsByCustomer)

    app.Listen(":5000")
}

// Create new customer
func createCustomer(c *fiber.Ctx) error {
    type Body struct {
        Email string `json:"email"`
        StripeCreditCardToken string `json: "stripeCreditCardToken"`
    }

    var body = new(Body)

    if err := c.BodyParser(body); err != nil {
        return err
    }

    if body.Email == "" || body.StripeCreditCardToken == "" {
        return fiber.NewError(402, "Validation error")
    }

    cus := customer.Create(body.Email, body.StripeCreditCardToken)
    log.Println(cus.ID)
    return c.JSON(struct{ StripeCustomerID string `json:"stripeCustomerID"` }{ cus.ID })
}

// new payment charge
func createPayment(c *fiber.Ctx) error {

    return c.Send([]byte("payment"))
}

// retrieves all payments for given customer
func paymentsByCustomer(c *fiber.Ctx) error {

    return c.Send([]byte("payment"))
}
