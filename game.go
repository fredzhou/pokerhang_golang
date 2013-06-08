package pokerhand

import (
    "strings"
    "reflect"
)

type game_t struct {
    Winner string
}

type category_intf interface {
    Format(hand []string) []string    
}

const (
    CATEGORY_HIGH_CARD = iota
    CATEGORY_PAIR
    CATEGORY_TWO_PAIRS
    CATEGORY_THREE_OF_KIND
    CATEGORY_STRAIGHT
    CATEGORY_FLUSH
    CATEGORY_FULL_HOUSE
    CATEGORY_FOUR_OF_KIND
    CATEGORY_STRAIGHT_FLUSH
)

var Game *game_t = new(game_t)

func isCategoryThreeOfKind(hand []string) bool {
    for i := range hand {
        if strings.Count(strings.Join(hand, ""), strings.TrimRight(hand[i], "CDHS")) == 3 {
            return true
        }
    }
    return false
}

func isCategoryTwoPairs(hand []string) bool {
    pairCardValue := 0
    secondPairCardValue := 0
    
    sortedHand := sortByValueDesc(hand)
    
    for i := range sortedHand {
        if pairCardValue == valueOfCard(sortedHand[i]) {
            for j := i+1; j < len(sortedHand); j++ {
                if secondPairCardValue == valueOfCard(sortedHand[j]) {
                    return true
                }
                secondPairCardValue = valueOfCard(sortedHand[j])
            }
        }
        pairCardValue = valueOfCard(sortedHand[i])
    }
    
    return false
}

func isCategoryPair(hand []string) bool {
    pairCardValue := 0

    for i := range hand {
        if pairCardValue == valueOfCard(hand[i]) {
            for j := i+1; j < len(hand); j++ {
                if valueOfCard(hand[j]) == pairCardValue {return false}
            }
            return true
        }
        pairCardValue = valueOfCard(hand[i])
    }
    
    return false
}

func isCategoryStraight(hand []string) bool {
    sortedHand := sortByValueDesc(hand)
    for i := 0; i < len(sortedHand) - 1; i ++ {
        if valueOfCard(sortedHand[i]) - valueOfCard(sortedHand[i+1]) != 1 {return false}
    }
	return true
}

func isCategoryFlush(hand []string) bool {
    for i := 0; i < len(hand) -1; i++ {
        if hand[i][len(hand[i])-1] != hand[i+1][len(hand[i+1])-1] {return false}
    }
    
    return true
}

func isCategoryFullHouse(hand []string) bool {
   
    if isCategoryThreeOfKind(hand) && isCategoryPair(hand) {
        return true
    } 
    return false
}

func isCategoryFourOfKind(hand []string) bool {
    sortedHand := sortByValueDesc(hand)
    if strings.Count(strings.Join(sortedHand, ""), strings.TrimRight(sortedHand[2], "CSDH"))  == 4 {
       return true
    }
    
    return false 
}

func isCategoryStraightFlush(hand []string) bool {
        if isCategoryStraight(hand) && isCategoryFlush(hand) {
            return true
        } 
        
        return false
}

func categorize(hand []string) category_intf {  
    var category category_intf
    switch {
        case isCategoryStraightFlush(hand): {category = newStraightFlush(); break}
        case isCategoryFourOfKind(hand): {category = newFourOfKind(); break}
        case isCategoryFullHouse(hand): {category = newFullHouse(); break}
        case isCategoryFlush(hand): {category = newFlush(); break}
        case isCategoryStraight(hand):  {category = newStraight(); break}
        case isCategoryThreeOfKind(hand): {category = newThreeOfKind(); break}
        case isCategoryTwoPairs(hand): {category = newTwoPairs(); break}
        case isCategoryPair(hand): {category = newPair(); break}
        default: {category = newHighCard()}
    }
    

    return category
}


func getCategoryRank (black []string, white []string) (blackCategory int, whiteCategory int){

    
    blackCategory = reflect.ValueOf(categorize(black)).Elem().FieldByName("Category").Interface().(int)
    whiteCategory = reflect.ValueOf(categorize(white)).Elem().FieldByName("Category").Interface().(int)
    
    return
}

func CompareSameCategory(black []string, white []string) { 
    formatedBlack := categorize(black).Format(black)
    formatedWhite := categorize(white).Format(white)
    
    steps := reflect.ValueOf(categorize(black)).Elem().FieldByName("Steps").Interface().([]int)

    subCompare(formatedBlack, formatedWhite, steps, 0)
    
}

func subCompare(black []string, white []string, Steps []int, index int) {
    if valueOfCard(black[index]) == valueOfCard(white[index]) && index < len(Steps) - 1  {
        index += Steps[index]
	    subCompare(black, white, Steps, index)
	} else {
	    compareByCard(valueOfCard(black[index]), valueOfCard(white[index]))
	    return
	}
}

func compareByCard(blackCardValue int , whiteCardValue int) {
    if blackCardValue > whiteCardValue {
        Game.Winner = "B"
    } else if blackCardValue < whiteCardValue {
        Game.Winner = "W"
    } else {
        Game.Winner = "Tie"
    }
}
func (game *game_t) Compare(black []string, white []string) {
    blackCategory, whiteCategory := getCategoryRank(black, white)
    
    if blackCategory  != whiteCategory {
        if blackCategory > whiteCategory {
            game.Winner = "B"
        } else {
            game.Winner = "W"
        }
    } else {
        CompareSameCategory(black, white)
    }

    return
}
