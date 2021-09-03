package customer
import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
    "log"
)

func Create(email string, token string) (*stripe.Customer, *stripe.Error) {
    log.Println("create customer " + email + ", " + token)

    c, err := customer.New(&stripe.CustomerParams{
        Email: &email,
        Source: &stripe.SourceParams{
            Token: &token,
        },
    })

    if err != nil {
        return nil, err.(*stripe.Error)
    }
    return c, nil
}

