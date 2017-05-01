package subscription

type Subscription struct {
	Clients    map[*Client]bool
	ItemSearch ItemSearch
	Output     chan string
}
