package subscription_handler

import (
	"encoding/json"
	"time"
)

const PaymentDateLayout = "2006-01-02"

type PaymentDate time.Time

func (p *PaymentDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strInput[1 : len(strInput)-1] // Remove quotes
	newTime, err := time.Parse(PaymentDateLayout, strInput)
	if err != nil {
		return err
	}

	*p = PaymentDate(newTime)
	return nil
}

func (p *PaymentDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*p).Format(PaymentDateLayout))
}
