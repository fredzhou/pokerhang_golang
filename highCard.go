package pokerhand

import (

)

type highCard_t struct {
    Steps []int
    Category int
}
func newHighCard() *highCard_t {
    highCard := new(highCard_t)
    
    highCard.Steps = []int{1, 1, 1, 1, 1}
    highCard.Category = CATEGORY_HIGH_CARD
    
    return highCard    
}

func  (highCard *highCard_t) Format(hand []string) []string{
    sortedHand := sortByValueDesc(hand)
    return sortedHand
}

