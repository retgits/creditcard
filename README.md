# creditcard

A Go module to perform basic credit card validation.

## Installation

```bash
go get github.com/retgits/creditcard
```

## Usage

Create a `creditcard.Card` struct and use the `Validate()` method to perform validation

```go
package main

import (
    "fmt"
    "github.com/retgits/creditcard"
)

func main() {
    card := creditcard.Card{
        Type:        "Something",
        Number:      "5019717010103742",
        ExpiryMonth: 11,
        ExpiryYear:  2019,
        CVV:         "1234",
    }
    validation := card.Validate()
    fmt.Printf("%+v\n", validation)
    fmt.Printf("%+v\n", validation.Card)
    // This prints
    // &{Card:0xc000092040 ValidCardNumber:false ValidExpiryMonth:true ValidExpiryYear:true ValidCVV:false IsExpired:false Errors:[given card type doesn't match determined card type]}
    // &{Type:Something Number:5019717010103742 ExpiryMonth:11 ExpiryYear:2019 CVV:1234}
}
```

The sample here shows that the card's supplied type "_Something_" doesn't match what the type actually should be.

## Supported Credit Card Types

This module supports a variety of credit cards:

- American Express
- Aura
- Bankcard
- Cabal
- China UnionPay
- Dankort
- Diners Club Carte Blanche
- Diners Club Enroute
- Diners Club International
- Discover
- Elo
- Hipercard
- InstaPayment
- InterPayment
- JCB
- Maestro
- Mastercard
- Visa
- Visa Electron

## Test Card Numbers

A list of test credit cards is available from [PayPal](http://www.paypalobjects.com/en_US/vhelp/paypalmanager_help/credit_card_numbers.htm).

## LICENSE

This module is provided under the [MIT](./LICENSE) license

## Acknowledgements

This lib is built as a Go equivalent to the Node.js [credit-card](https://github.com/cjihrig/credit-card/blob/master/README.md) library and the amazing work by [HubCash](https://github.com/hubcash/cards)
