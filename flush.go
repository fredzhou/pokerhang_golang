package pokerhand

type flush_t struct {
    Steps []int
    Category int
}
func newFlush() *flush_t {
    flush := new(flush_t)    
    flush.Steps = []int{1, 1, 1, 1, 1}
    flush.Category = CATEGORY_FLUSH
    
    return flush    
}



func  (flush *flush_t)Format(hand []string) []string{
    sortedHand := sortByValueDesc(hand)
    return sortedHand
}

