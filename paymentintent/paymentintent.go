package paymentintent
import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/paymentintent"
    "log"
)

func Create(customerId string, amount int64) (*stripe.PaymentIntent, *stripe.Error) {
    log.Println("create payment intent")

    intent, err := paymentintent.New(&stripe.PaymentIntentParams{
        Customer: &customerId,
        Amount: stripe.Int64(amount),
        Currency: stripe.String(string(stripe.CurrencyUSD)),
        PaymentMethodTypes: []*string {
            stripe.String("card"),
        },
    })

    if err != nil {
        return nil, err.(*stripe.Error)
    }

    log.Println(err)
    log.Println(intent)
    return intent, nil
}

func ListByCustomer(customerId string) ([]*stripe.PaymentIntent) {
    log.Println("list PI")

    list := paymentintent.List(&stripe.PaymentIntentListParams{
        Customer: &customerId,
    })

    var result []*stripe.PaymentIntent

    for list.Next() {
        var pi = list.PaymentIntent()
        log.Printf("%s %d\n", pi.ID, pi.Amount)
        result = append(result, pi)
    }

    return result
}
