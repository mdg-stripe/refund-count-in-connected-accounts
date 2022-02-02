
import stripe
stripe.api_key = 'sk_test_...'


def main():
    allAccountsRefundCount = 0

    # get the list of all connected accounts
    # see API reference at: https://stripe.com/docs/api/accounts/list
    accounts = stripe.Account.list()
    for a in accounts:
        # for each connected account, get the count of refunds
        refunds = stripe.Refund.list(stripe_account=a.id)
        thisAccountRefundCount = 0
        for r in refunds:
            if r.status == 'succeeded':
                # only count successful refunds
                thisAccountRefundCount += 1
        print('Account ID %s has %d refunds' % (a.id, thisAccountRefundCount))
        allAccountsRefundCount += thisAccountRefundCount
    print('TOTAL succeded refunds across all accounts: %d' % (allAccountsRefundCount))

if __name__ == '__main__':
    main()