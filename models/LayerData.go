package models

type LayerData struct {
	Id        int
	Name      string
	Upvotes   int
	Downvotes int
}

func (m LayerData) TotalVotes() int {
	return m.Upvotes + m.Downvotes
}

// VotePercentagePositive Calculate positive percentage
func (m LayerData) VotePercentagePositive() float64 {
	return (float64(m.Upvotes) * float64(100)) / float64(m.TotalVotes())
}

// VotePercentageNegative Calculate negative percentage
func (m LayerData) VotePercentageNegative() float64 {
	return (float64(m.Downvotes) * float64(100)) / float64(m.TotalVotes())
}

// VotePercentageDifference Calculate the difference between positive and negative percentages
func (m LayerData) VotePercentageDifference() float64 {
	return m.VotePercentagePositive() - m.VotePercentageNegative()
}
