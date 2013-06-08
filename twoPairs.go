package pokerhand

type twoPairs_t struct {
    Steps []int
    Category int
}

func newTwoPairs() *twoPairs_t {
    twoPairs := new(twoPairs_t)
    
    twoPairs.Steps = []int{2, 2, 1}
    twoPairs.Category = CATEGORY_TWO_PAIRS
    
    return twoPairs    
}

func indexOfSingleCardInTwoPairs(hand []string) int{
    var isPair bool
    var index int
    for i := range hand {
        isPair = false
        for j := range hand {
            if i == j {continue}
            if valueOfCard(hand[i]) == valueOfCard(hand[j]) {isPair = true}
        }
        if !isPair {index = i; break}
    }
    
    return index
}

func moveSingleCardToTwoPairsRear(hand []string) []string{
    result := sortByValueDesc(hand)
    index := indexOfSingleCardInTwoPairs(result)
   
    singleCard := result[index]
    for i := index; i < len(result); i++ {
     if i+1 < len(result) {
         result[i] = result[i+1]
     }
    }
    result[4] = singleCard
    return result
}

func  (twoPairs *twoPairs_t) Format(hand []string) []string{    
    result := moveSingleCardToTwoPairsRear(hand)
    
    return result
}
