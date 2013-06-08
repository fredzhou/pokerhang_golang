package pokerhand

import "strings"

type fullHouse_t struct {
    Steps []int
    Category int
}

func newFullHouse() *fullHouse_t {
    fullHouse := new(fullHouse_t)    
    fullHouse.Steps = []int{3, 2}
    fullHouse.Category = CATEGORY_FULL_HOUSE
    
    return fullHouse 
}



func  (fullHouse *fullHouse_t)Format(hand []string) []string{
    sortedHand := sortByValueDesc(hand)
    var wholeHand string
    if sortedHand[0] != sortedHand[2] {
        threePart := strings.Join(sortedHand[2:4], "|")
        pairPart := strings.Join(sortedHand[0:1], "|")
        
        wholeHand = threePart + "|" + pairPart
    }
    return strings.Split(wholeHand, "|")
}
