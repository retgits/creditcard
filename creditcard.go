// Package creditcard is a module that performs credit card validation.
package creditcard

import (
	"fmt"
	"strconv"
	"time"
)

// Card is a struct that contains a credit card type. This holds generic information about the credit card
type Card struct {
	// Type is an optional string with one of the supported card types. If the type is supplied, it will be validated that the type matches the number
	Type string
	// Number is the credit card number
	Number string
	// ExpiryMonth is the credit card expiration month
	ExpiryMonth int
	// ExpiryYear is the credit card expiration year
	ExpiryYear int
	// CVV is the credit card CVV code
	CVV string
}

// Validation is the object returned by the validate method, which contains the validation result of the card
type Validation struct {
	// A pointer to the struct that was passed in for validation
	Card *Card
	// ValidCardNumber is a boolean that indicates if a credit card number is valid for a given credit card type if given and verifies that the credit card number passes the Luhn algorithm
	ValidCardNumber bool
	// ValidExpiryMonth is a boolean that indicates if a value is a valid credit card expiry month (the range is 1 to 12)
	ValidExpiryMonth bool
	// ValidExpiryYear is a boolean that indicates if a value is a valid credit card expiry year (the range is 1900 to 2200)
	ValidExpiryYear bool
	// ValidCVV is a boolean that indicates if a CVV is valid for a given credit card type. For example, American Express requires a four digit CVV, while Visa and Mastercard require a three digit CVV
	ValidCVV bool
	// IsExpired is a boolean that indicates if a credit card's expiration date has been reached
	IsExpired bool
	// Errors is an array of validation errors that might occur during validation
	Errors []string
}

type cardType int

const (
	// Unknown card type
	Unknown cardType = iota
	// AmericanExpress  card type
	AmericanExpress
	// Aura card type
	Aura
	// Bankcard card type
	Bankcard
	// Cabal card type
	Cabal
	// ChinaUnionPay card type
	ChinaUnionPay
	// Dankort card type
	Dankort
	// DinersClubCarteBlanche card type
	DinersClubCarteBlanche
	// DinersClubEnroute card type
	DinersClubEnroute
	// DinersClubInternational card type
	DinersClubInternational
	// Discover card type
	Discover
	// Elo card type
	Elo
	// Hipercard card type
	Hipercard
	// InstaPayment card type
	InstaPayment
	// InterPayment card type
	InterPayment
	// JCB card type
	JCB
	// Maestro card type
	Maestro
	// Mastercard card type
	Mastercard
	// Visa card type
	Visa
	// VisaElectron card type
	VisaElectron
)

var cardTypeNames = [...]string{
	"Unknown Card",
	"American Express",
	"Aura",
	"Bankcard",
	"Cabal",
	"China UnionPay",
	"Dankort",
	"Diners Club Carte Blanche",
	"Diners Club Enroute",
	"Diners Club International",
	"Discover",
	"Elo",
	"Hipercard",
	"InstaPayment",
	"InterPayment",
	"JCB",
	"Maestro",
	"Mastercard",
	"Visa",
	"Visa Electron",
}

type digits [6]int

// at returns the digits from the start to the given length
func (d *digits) at(i int) int {
	return d[i-1]
}

// name returns the string representation of the card
func (cardType cardType) name() string {
	return cardTypeNames[cardType]
}

// Validate performs validation on the card. Apart from a copy of the card, it also returns
// - ValidCardNumber is a boolean that indicates if a credit card number is valid for a given credit card type if given and verifies that the credit card number passes the Luhn algorithm
// - ValidExpiryMonth is a boolean that indicates if a value is a valid credit card expiry month (the range is 1 to 12)
// - ValidExpiryYear is a boolean that indicates if a value is a valid credit card expiry year (the range is 1900 to 2200)
// - ValidCVV is a boolean that indicates if a CVV is valid for a given credit card type. For example, American Express requires a four digit CVV, while Visa and Mastercard require a three digit CVV
// - IsExpired is a boolean that indicates if a credit card's expiration date has been reached
// - Errors is an array of validation errors that might occur during validation
func (c *Card) Validate() *Validation {
	val := &Validation{
		Card:   c,
		Errors: make([]string, 0),
	}

	val.ValidExpiryMonth = c.validExpiryMonth()
	if !val.ValidExpiryMonth {
		val.Errors = append(val.Errors, fmt.Sprintf("month '%d' is not a valid month", c.ExpiryMonth))
	}

	val.ValidExpiryYear = c.validExpiryYear()
	if !val.ValidExpiryYear {
		val.Errors = append(val.Errors, fmt.Sprintf("year '%d' is not a valid year", c.ExpiryYear))
	}

	val.IsExpired = c.isExpired()
	if val.IsExpired {
		val.Errors = append(val.Errors, "creditcard is expired")
	}

	if len(c.Type) == 0 {
		cardType, err := c.determineCardType()
		if err != nil {
			val.Errors = append(val.Errors, err.Error())
		}
		c.Type = cardType.name()
	}

	val.ValidCVV = c.matchCVV()
	if val.ValidCVV {
		val.Errors = append(val.Errors, "cvv doesn't match")
	}

	validNumber, err := c.validCardNumber()
	if err != nil {
		val.Errors = append(val.Errors, err.Error())
	}
	val.ValidCardNumber = validNumber
	if val.ValidCardNumber {
		val.Errors = append(val.Errors, "card number is not valid")
	}

	return val
}

// validCardNumber checks whether the given card type matches the actual expected card type and whether the number passes the luhn check
func (c *Card) validCardNumber() (bool, error) {
	cardType, err := c.determineCardType()
	if err != nil {
		return false, err
	}

	if cardType.name() != c.Type {
		return false, fmt.Errorf("given card type doesn't match determined card type")
	}

	return c.validateLuhn(), nil
}

// validExpiryMonth validates whether the expiry month is a proper month (between 1 and 12)
func (c *Card) validExpiryMonth() bool {
	if c.ExpiryMonth < 1 || 12 < c.ExpiryMonth {
		return false
	}
	return true
}

// validExpiryYear validates whether the expiry year is a valid year (between 1900 and 2200)
func (c *Card) validExpiryYear() bool {
	if c.ExpiryYear < 1900 || c.ExpiryYear > 2200 {
		return false
	}
	return true
}

// isExpired is a boolean that indicates whether the card is expired or not
func (c *Card) isExpired() bool {
	if !c.validExpiryMonth() || !c.validExpiryYear() {
		return true
	}

	date := fmt.Sprintf("%d-%d-01", c.ExpiryYear, c.ExpiryMonth)
	parsetime, _ := time.Parse("2006-01-02", date)

	return parsetime.Before(time.Now())
}

// matchCVV checks whether the CVV length matches the expected length
func (c *Card) matchCVV() bool {
	switch c.Type {
	case "American Express":
		return len(c.CVV) == 4
	default:
		return len(c.CVV) == 3
	}
}

// determineCardType determines which card type the credit card has
func (c *Card) determineCardType() (cardType, error) {
	ccLen := len(c.Number)
	ccDigits := digits{}

	// Take the first 6 digits of the card number,
	// convert to a integer to allow easy comparison after
	for i := 0; i < 6; i++ {
		if i < ccLen {
			ccDigits[i], _ = strconv.Atoi(c.Number[:i+1])
		}
	}

	// The switch below compares the first digits, and the security code size,
	// for return a company for each bin range using the card number
	switch {
	case ccDigits.at(4) == 4011 || ccDigits.at(6) == 431274 || ccDigits.at(6) == 438935 ||
		ccDigits.at(6) == 451416 || ccDigits.at(6) == 457393 || ccDigits.at(4) == 4576 ||
		ccDigits.at(6) == 457631 || ccDigits.at(6) == 457632 || ccDigits.at(6) == 504175 ||
		ccDigits.at(6) == 627780 || ccDigits.at(6) == 636297 || ccDigits.at(6) == 636368 ||
		ccDigits.at(6) == 636369 || (ccDigits.at(6) >= 506699 && ccDigits.at(6) <= 506778) ||
		(ccDigits.at(6) >= 509000 && ccDigits.at(6) <= 509999) ||
		(ccDigits.at(6) >= 650031 && ccDigits.at(6) <= 650051) ||
		(ccDigits.at(6) >= 650035 && ccDigits.at(6) <= 650033) ||
		(ccDigits.at(6) >= 650405 && ccDigits.at(6) <= 650439) ||
		(ccDigits.at(6) >= 650485 && ccDigits.at(6) <= 650538) ||
		(ccDigits.at(6) >= 650541 && ccDigits.at(6) <= 650598) ||
		(ccDigits.at(6) >= 650700 && ccDigits.at(6) <= 650718) ||
		(ccDigits.at(6) >= 650720 && ccDigits.at(6) <= 650727) ||
		(ccDigits.at(6) >= 650901 && ccDigits.at(6) <= 650920) ||
		(ccDigits.at(6) >= 651652 && ccDigits.at(6) <= 651679) ||
		(ccDigits.at(6) >= 655000 && ccDigits.at(6) <= 655019) ||
		(ccDigits.at(6) >= 655021 && ccDigits.at(6) <= 655021):
		return Elo, nil

	case ccDigits.at(6) >= 604201 && ccDigits.at(6) <= 604219:
		return Cabal, nil

	case ccDigits.at(6) == 384100 || ccDigits.at(6) == 384140 || ccDigits.at(6) == 384160 ||
		ccDigits.at(6) == 606282 || ccDigits.at(6) == 637095 || ccDigits.at(4) == 637568 ||
		ccDigits.at(4) == 637599 || ccDigits.at(4) == 637609 || ccDigits.at(4) == 637612:
		return Hipercard, nil

	case ccDigits.at(2) == 34 || ccDigits.at(2) == 37:
		return AmericanExpress, nil

	case ccDigits.at(4) == 5610 || (ccDigits.at(6) >= 560221 && ccDigits.at(6) <= 560225):
		return Bankcard, nil

	case ccDigits.at(2) == 62:
		return ChinaUnionPay, nil

	case ccDigits.at(3) >= 300 && ccDigits.at(3) <= 305 && ccLen == 15:
		return DinersClubCarteBlanche, nil

	case ccDigits.at(4) == 2014 || ccDigits.at(4) == 2149:
		return DinersClubEnroute, nil

	case ((ccDigits.at(3) >= 300 && ccDigits.at(3) <= 305) || ccDigits.at(3) == 309 ||
		ccDigits.at(2) == 36 || ccDigits.at(2) == 38 || ccDigits.at(2) == 39) && ccLen <= 14:
		return DinersClubInternational, nil

	case ccDigits.at(4) == 6011 || (ccDigits.at(6) >= 622126 && ccDigits.at(6) <= 622925) ||
		(ccDigits.at(3) >= 644 && ccDigits.at(3) <= 649) || ccDigits.at(2) == 65:
		return Discover, nil

	case ccDigits.at(3) == 636 && ccLen >= 16 && ccLen <= 19:
		return InterPayment, nil

	case ccDigits.at(3) >= 637 && ccDigits.at(3) <= 639 && ccLen == 16:
		return InstaPayment, nil

	case ccDigits.at(4) == 5018 || ccDigits.at(4) == 5020 || ccDigits.at(4) == 5038 ||
		ccDigits.at(4) == 5612 || ccDigits.at(4) == 5893 || ccDigits.at(4) == 6304 ||
		ccDigits.at(4) == 6759 || ccDigits.at(4) == 6761 || ccDigits.at(4) == 6762 ||
		ccDigits.at(4) == 6763 || c.Number[:3] == "0604" || ccDigits.at(4) == 6390:
		return Maestro, nil

	case ccDigits.at(4) == 5019:
		return Dankort, nil

	case ccDigits.at(2) >= 51 && ccDigits.at(2) <= 55:
		return Mastercard, nil

	case ccDigits.at(2) == 35:
		return JCB, nil

	case ccDigits.at(2) == 50:
		return Aura, nil

	case ccDigits.at(4) == 4026 || ccDigits.at(6) == 417500 || ccDigits.at(4) == 4405 ||
		ccDigits.at(4) == 4508 || ccDigits.at(4) == 4844 || ccDigits.at(4) == 4913 ||
		ccDigits.at(4) == 4917:
		return VisaElectron, nil

	case ccDigits.at(1) == 4:
		return Visa, nil

	default:
		return 0, fmt.Errorf("unknown creditcard type")
	}
}

// http://en.wikipedia.org/wiki/Luhn_algorithm
// validateLuhn will check the credit card's number against the Luhn algorithm
func (c *Card) validateLuhn() bool {
	var sum int
	var alternate bool

	// Gets the Card number length
	numberLen := len(c.Number)

	// For numbers that is lower than 13 and
	// bigger than 19, must return as false
	if numberLen < 13 || numberLen > 19 {
		return false
	}

	// Parse all numbers of the card into a for loop
	for i := numberLen - 1; i > -1; i-- {
		// Takes the mod, converting the current number in integer
		mod, _ := strconv.Atoi(string(c.Number[i]))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate
		sum += mod
	}

	return sum%10 == 0
}
