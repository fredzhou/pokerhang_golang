package pokerhand

import (
    "strings"
)

type threeOfKind_t struct {
    Steps []int
    Category int
}
func newThreeOfKind() *threeOfKind_t {
    threeOfKind := new(threeOfKind_t)
    threeOfKind.Steps = []int{3, 1, 1}
    threeOfKind.Category = CATEGORY_THREE_OF_KIND
    return threeOfKind
}

func indexOfThree(hand []string) int {
    for i := range hand {
        if strings.Contains(hand[i], strings.TrimRight(hand[2], "CSDH")) {return i}
    }
    
    return 0

}

func moveThreeToHead(hand []string, index int) []string {
    stringThree := strings.Join(hand[index:index+2], "|")

    stringOthers := strings.Replace(strings.Join(hand, "|"), stringThree, "", -1)

    stringOthers = strings.Replace(stringOthers, "||", "|", -1)

    stringHand := stringThree + "|" + stringOthers

    return strings.Split(stringHand, "|")

}

func  (threeOfKind *threeOfKind_t)Format(hand []string) []string{
    sortedHand := sortByValueDesc(hand)
    index := indexOfThree(sortedHand)

    result := moveThreeToHead(sortedHand, index)

    return result
}
