package pokerhand

import "strings"

type fourOfKind_t struct {
    Steps []int
    Category int
}
func newFourOfKind() *fourOfKind_t {
    fourOfKind := new(fourOfKind_t)    
    fourOfKind.Category = CATEGORY_FOUR_OF_KIND	
    return fourOfKind
}

func  (fourOfKind *fourOfKind_t)Format(hand []string) []string{
    sortedHand := sortByValueDesc(hand)
    var fourPart, singlePart, wholeHand string
    
    if sortedHand[0] != sortedHand[1] {
        fourPart = strings.Join(sortedHand[1:4], "|")
        singlePart = sortedHand[0]
    } else {
        fourPart = strings.Join(sortedHand[0:3], "|")
        singlePart = sortedHand[4]
    }
    
    wholeHand = fourPart + "|" + singlePart
    
    return strings.Split(wholeHand, "|")
}
