package pokerhand

type straight_t struct {
    Steps []int
    Category int
}
func newStraight() *straight_t {
    straight := new(straight_t)    
    straight.Steps = []int{1, 1, 1, 1, 1}
    straight.Category = CATEGORY_STRAIGHT
    
    return straight    
}

func  (straight *straight_t)Format(hand []string) []string{
    sortedHand := sortByValueDesc(hand)
    return sortedHand
}
