package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan withdraw)

type withdraw struct {
	amount   int
	response chan<- bool
}

// Deposit deposits amount into the account
func Deposit(amount int) { deposits <- amount }

// Balance check the account balance
func Balance() int { return <-balances }

// Withdraw withdraws amount from the account
func Withdraw(amount int) bool {
	response := make(chan bool)
	withdraws <- withdraw{amount, response}
	return <-response
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdraw := <-withdraws:
			if withdraw.amount > balance {
				withdraw.response <- false
			} else {
				balance -= withdraw.amount
				withdraw.response <- true
			}
		case balances <- balance:
			// do nothing
		}
	}
}

func init() {
	go teller()
}
