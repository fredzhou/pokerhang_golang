package pokerhand


type pair_t struct {
    Steps []int
    Category int
}
func newPair() *pair_t {
    pair := new(pair_t)    
    pair.Steps = []int{2, 1, 1, 1}
    pair.Category = CATEGORY_PAIR
    
    return pair    
}

func locatePair(hand []string) (index int) {
    for i := range hand {
        if valueOfCard(hand[i]) == valueOfCard(hand[i+1]) {
            index = i
            break 
        }
    }    
    
    return 
    
}

func movePairToHead(hand []string, pairIndex int) []string {
    pairCard1 := hand[pairIndex]
    pairCard2 := hand[pairIndex+1] 
    
    for i := pairIndex - 1; i >= 0; i-- {
        hand[i+2] = hand[i]
    }
    hand[0] = pairCard1
    hand[1] = pairCard2    
    
    return hand
}

func  (pair *pair_t)Format(hand []string) []string{
    sortedHand := sortByValueDesc(hand)

    pairIndex := locatePair(sortedHand)
       
    result := movePairToHead(sortedHand, pairIndex)
    
    return result
}
