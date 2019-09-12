package creditcard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCards(t *testing.T) {
	assert := assert.New(t)

	card := Card{
		Type: "Something", Number: "5019717010103742", ExpiryMonth: 11, ExpiryYear: 2019, CVV: "1234",
	}
	val := card.Validate()
	assert.Contains(val.Errors, "given card type doesn't match determined card type")

	card = Card{
		Type: "Something", Number: "5019717010103742", ExpiryMonth: 111, ExpiryYear: 2019, CVV: "1234",
	}
	val = card.Validate()
	assert.Contains(val.Errors, "month '111' is not a valid month")

	card = Card{
		Type: "Something", Number: "5019717010103742", ExpiryMonth: 11, ExpiryYear: 1899, CVV: "1234",
	}
	val = card.Validate()
	assert.Contains(val.Errors, "year '1899' is not a valid year")

	card = Card{
		Type: "Dankort", Number: "5019717010103742", ExpiryMonth: 11, ExpiryYear: 1899, CVV: "1234",
	}
	val = card.Validate()
	assert.Contains(val.Errors, "year '1899' is not a valid year")

	card = Card{
		Number: "5019717010103742", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Dankort")
	assert.Empty(val.Errors)

	card = Card{
		Number: "0000000000", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Contains(val.Errors, "unknown creditcard type")

	card = Card{
		Number: "378282246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "American Express")
	assert.Empty(val.Errors)

	card = Card{
		Number: "655021246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Elo")
	assert.Empty(val.Errors)

	card = Card{
		Number: "604201246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Cabal")
	assert.Empty(val.Errors)

	card = Card{
		Number: "384140246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Hipercard")
	assert.Empty(val.Errors)

	card = Card{
		Number: "560221246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Bankcard")
	assert.Empty(val.Errors)

	card = Card{
		Number: "620221246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "China UnionPay")
	assert.Empty(val.Errors)

	card = Card{
		Number: "300221246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Diners Club Carte Blanche")
	assert.Empty(val.Errors)

	card = Card{
		Number: "201421246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Diners Club Enroute")
	assert.Empty(val.Errors)

	card = Card{
		Number: "39022124631000", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Diners Club International")
	assert.Empty(val.Errors)

	card = Card{
		Number: "601121246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Discover")
	assert.Empty(val.Errors)

	card = Card{
		Number: "63612124631000500", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "InterPayment")
	assert.Empty(val.Errors)

	card = Card{
		Number: "6371212463100050", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "InstaPayment")
	assert.Empty(val.Errors)

	card = Card{
		Number: "501821246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Maestro")
	assert.Empty(val.Errors)

	card = Card{
		Number: "511821246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Mastercard")
	assert.Empty(val.Errors)

	card = Card{
		Number: "351821246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "JCB")
	assert.Empty(val.Errors)

	card = Card{
		Number: "508821246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Aura")
	assert.Empty(val.Errors)

	card = Card{
		Number: "402621246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Visa Electron")
	assert.Empty(val.Errors)

	card = Card{
		Number: "409921246310005", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	val = card.Validate()
	assert.Equal(val.Card.Type, "Visa")
	assert.Empty(val.Errors)

	card = Card{
		Number: "0000000000", ExpiryMonth: 11, ExpiryYear: 2020, CVV: "1234",
	}
	luhn := card.validateLuhn()
	assert.Equal(luhn, false)
}
