package pokerhand


type straightFlush_t struct {
    Steps []int
    Category int
}
func newStraightFlush() *straightFlush_t {
    straightFlush := new(straightFlush_t)    
    straightFlush.Category = CATEGORY_STRAIGHT_FLUSH	
    return straightFlush
}

func  (straightFlush *straightFlush_t)Format(hand []string) []string{
    return sortByValueDesc(hand)
}
