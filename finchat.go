package main
import (
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/stripe/stripe-go"
    _ "finchat/dotenv"
    "finchat/customer"
    "finchat/paymentintent"
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

    app.Listen("0.0.0.0:" + os.Getenv("PORT"))
}

// Create new customer
func createCustomer(c *fiber.Ctx) error {
    type Body struct {
        Email string `json:"email"`
        Token string `json:"stripeCreditCardToken"`
    }

    var body = new(Body)

    if err := c.BodyParser(body); err != nil {
        return err
    }

    if body.Email == "" || body.Token == "" {
        return fiber.NewError(400, "Validation error")
    }

    cus, err := customer.Create(body.Email, body.Token)
    if err != nil {
        return fiber.NewError(err.HTTPStatusCode, err.Error())
    }

    return c.JSON(struct{ StripeCustomerID string `json:"stripeCustomerID"` }{ cus.ID })
}

// new payment charge
func createPayment(c *fiber.Ctx) error {
    type Body struct {
        CustomerId string `json:"stripeCustomerID"`
        Amount int64 `json:"amount"`
    }

    var body = new(Body)
    if err := c.BodyParser(body); err != nil {
        return err
    }

    if body.CustomerId == "" || body.Amount == 0 {
        return fiber.NewError(400, "Validation error")
    }

    intent, err := paymentintent.Create(body.CustomerId, body.Amount)
    if err != nil {
        return fiber.NewError(err.HTTPStatusCode, err.Error())
    }

    return c.JSON(struct {  PaymentIntentID string `json:"paymentIntentID"` }{ intent.ID })
}

// retrieves all payments for given customer
func paymentsByCustomer(c *fiber.Ctx) error {

    var customerId = c.Params("customerId")

    piList := paymentintent.ListByCustomer(customerId)

    type Payment struct {
        Id string `json:"id"`
        Amount int64 `json:"amount"`
    }

    var payments []Payment

    for _, pi := range piList {
        payments = append(payments, Payment {
            pi.ID,
            pi.Amount,
        })
    }

    type Response struct {
        Payments []Payment `json:"payments"`
    }
    return c.JSON(Response{ payments })
}
