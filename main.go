package main

import (
	"log"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/account"
	"github.com/stripe/stripe-go/v72/refund"
)

func main() {
	stripe.Key = "sk_test_..."

	allAccountsRefundCount := 0

	// get the list of all connected accounts
	// see API reference at: https://stripe.com/docs/api/accounts/list
	params := &stripe.AccountListParams{}
	i := account.List(params)
	for i.Next() {
		a := i.Account()
		// for each connected account, get the count of refunds
		refundParams := &stripe.RefundListParams{
			// see filter options at: https://stripe.com/docs/api/refunds/list
			//CreatedRange: &stripe.RangeQueryParams{GreaterThan: <UNIX timestamp for whatever date you want to filter by>},
		}
		// important: set the account ID for which we're getting the list of refunds
		refundParams.SetStripeAccount(a.ID)
		j := refund.List(refundParams)
		thisAccountRefundCount := 0
		for j.Next() {
			if j.Refund().Status == stripe.RefundStatusSucceeded {
				thisAccountRefundCount++
			}
		}
		log.Printf("Account ID %s has %d refunds", a.ID, thisAccountRefundCount)
		allAccountsRefundCount += thisAccountRefundCount
	}
	log.Printf("==========")
	log.Printf("TOTAL succeeded refunds across all connected accounts: %d", allAccountsRefundCount)
}
