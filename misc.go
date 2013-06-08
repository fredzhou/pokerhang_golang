package pokerhand

import (
    "strings"
    "strconv"
)

func sortByValueDesc(hand []string) []string {
    var tmp string
    var result []string = make([]string, len(hand))
    
    copy(result, hand)
    for i := range result {
        for j := i+1; j < len(result); j++{
	    	if valueOfCard(result[i]) < valueOfCard(result[j]) {
	    	    tmp = result[i]
	    	    result[i] = result[j]
	    	    result[j] = tmp
	    	}
    	}
    }
    return result
}

func valueOfCard(hand string) int {
    trimmedCard := strings.TrimRight(hand, "CDHS")
    
    var cardValue int
    switch trimmedCard {
        case "J": { 
            cardValue = 11
            break
        }
        case "Q": {
            cardValue = 12
            break
        }
        case "K": {
            cardValue = 13
            break
        }
        case "A": {
            cardValue = 14
            break
        }
        default: {cardValue,_ = strconv.Atoi(trimmedCard)}
    }
    
    return cardValue
}


