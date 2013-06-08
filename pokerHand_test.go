package pokerhand

import (
    "testing"
    "strings"
)

func AssertTrue(t *testing.T, condition bool, desc string) {
        if (!condition) {
                t.Error(desc)
        }
}


func TestValueOfCard(t *testing.T) {
    hand := "8D"
    AssertTrue(t, valueOfCard(hand) == 8, "")

    hand = "7H"
    AssertTrue(t, valueOfCard(hand) == 7, "")

    hand = "6C"
    AssertTrue(t, valueOfCard(hand) == 6, "")

    hand = "5S"
    AssertTrue(t, valueOfCard(hand) == 5, "")

}

func TestHighCardFor5thCard(t *testing.T) {
    black := []string{"2C", "4D", "6H", "7H", "8D"}
    white := []string{"2C", "4D", "6H", "7H", "9D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")

    black = []string{"2C", "4D", "6H", "7H", "10D"}
    white = []string{"2C", "4D", "6H", "7H", "9D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

    black = []string{"2C", "4D", "6H", "7H", "JD"}
    white = []string{"2C", "4D", "6H", "7H", "10D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")
}

func TestHighCardFor1stCard(t *testing.T) {
    black := []string{"2C", "4D", "6H", "7H", "8D"}
    white := []string{"3C", "4D", "6H", "7H", "8D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")

    black = []string{"3C", "4D", "6H", "7H", "8D"}
    white = []string{"2C", "4D", "6H", "7H", "8D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

}

func TestHighCardforInMiddleCard(t *testing.T) {
    black := []string{"3C", "4D", "5H", "7H", "8D"}
    white := []string{"3C", "4D", "6H", "7H", "8D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")

}

func TestPair(t *testing.T) {
    black := []string{"3C", "4D", "5H", "5S", "8D"}
    white := []string{"3C", "4D", "6H", "6S", "7D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")

    black = []string{"3C", "4D", "5H", "5S", "7D"}
    white = []string{"3C", "4H", "4S", "7D", "8H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

}

func TestPairXHighCard(t *testing.T) {
    black := []string{"3C", "4D", "5H", "6S", "8D"}
    white := []string{"3C", "4D", "6H", "6S", "7D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")

    black = []string{"3C", "4D", "5H", "5S", "7D"}
    white = []string{"3C", "4D", "6H", "7S", "8D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

}

func TestTwoPairs(t *testing.T) {
    black := []string{"4C", "4D", "5H", "5S", "8D"}
    white := []string{"3C", "3D", "6H", "6S", "7D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")

    black = []string{"4C", "4D", "6H", "6S", "8D"}
    white = []string{"3C", "3D", "6H", "6S", "7D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

}

func TestTwoPairsFormat(t *testing.T) {
    twoPairs := new(twoPairs_t)
    hand := []string{"9C", "4H", "5H", "5S", "4S"}
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[0], "5"), "")
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[2], "4"), "")
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[4], "9"), "")

    hand = []string{"4C", "3H", "5H", "5S", "3S"}
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[0], "5"), "")
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[2], "3"), "")
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[4], "4"), "")

    hand = []string{"5S", "9H", "4C", "5D", "9S"}
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[0], "9"), "")
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[2], "5"), "")
    AssertTrue(t, strings.Contains(twoPairs.Format(hand)[4], "4"), "")
}

func TestThreeOfKind(t *testing.T) {
    black := []string{"4C", "4D", "4S", "5S", "8D"}
    white := []string{"3C", "3D", "3H", "6S", "7D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

    black = []string{"4C", "4D", "4S", "5S", "6D"}
    white = []string{"3C", "3D", "3H", "8S", "9D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

    black = []string{"5C", "5D", "5S", "3S", "4S"}
    white = []string{"4C", "4D", "4H", "8S", "9D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")
}

func TestStraight(t *testing.T) {
    black := []string{"4C", "5D", "6S", "7S", "8D"}
    white := []string{"3C", "4D", "5H", "6S", "7D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

    black = []string{"5D", "4C", "6S", "7S", "8D"}
    white = []string{"6S", "7D", "3C", "4D", "5H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

    black = []string{"5D", "4C", "6S", "7S", "8D"}
    white = []string{"6S", "9D", "3C", "4D", "5H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

    black = []string{"4C", "5D", "6S", "7S", "8D"}
    white = []string{"3C", "3D", "3H", "9S", "10D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")
}

func TestFlush (t *testing.T) {
    black := []string{"2C", "4C", "6C", "7C", "8C"}
    white := []string{"2D", "4D", "6D", "7D", "9D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")

    black = []string{"10H", "JH", "QH", "KH", "AH"}
    white = []string{"2D", "4D", "6D", "7D", "8D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")

    black = []string{"10C", "JD", "QS", "KS", "AD"}
    white = []string{"2D", "4D", "6D", "7D", "8D"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")
}

func TestFullHouse(t *testing.T) {
    black := []string{"2C", "4C", "6C", "7C", "8C"}
    white := []string{"4D", "4S", "3C", "3D", "3H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")
    
    black = []string{"6D", "6S", "5C", "5D", "5H"}
    white = []string{"4D", "4S", "3C", "3D", "3H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")    

    black = []string{"6D", "6S", "5C", "5D", "5H"}
    white = []string{"AD", "AS", "3C", "3D", "3H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")     
    
}


func testXCategoryFourOfKind(t *testing.T) {
    black := []string{"2C", "2D", "2H", "2S", "8C"}
    white := []string{"4D", "4S", "3C", "3D", "3H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")
    
    black = []string{"5C", "5D", "5H", "5S", "8C"}
    white = []string{"4D", "4S", "3C", "3D", "3H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")  
    
    black = []string{"8C", "5C", "5D", "5H", "5S"}
    white = []string{"4D", "4S", "3C", "3D", "3H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")      
}    

func testSameCategoryFourOfKind(t *testing.T) {

    black := []string{"8H", "2C", "2D", "2H", "2S"}
    white := []string{"3H", "9D", "9S", "9C", "9H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")
    
    black = []string{"4H", "5C", "5D", "5H", "5S"}
    white = []string{"3H", "9D", "9S", "9C", "9H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")
    
    black = []string{"AH", "5C", "5D", "5H", "5S"}
    white = []string{"QH", "9D", "9S", "9C", "9H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")    
    
    black = []string{"AH", "5C", "5D", "5H", "5S"}
    white = []string{"3H", "9D", "9S", "9C", "9H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")        
}

func TestFourOfKind(t *testing.T) {
    testXCategoryFourOfKind(t) 
    testSameCategoryFourOfKind(t)
}

func testXCategoryStraightFlush(t *testing.T) {
    black := []string{"8H", "2C", "2D", "2H", "2S"}
    white := []string{"3H", "4H", "5H", "6H", "7H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")    
    
    black = []string{"8H", "2C", "2D", "2H", "2S"}
    white = []string{"7H", "4H", "5H", "6H", "3H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "W", "")   
}

func testSameCategoryStraightFlush(t *testing.T) {
    black := []string{"4C", "5C", "6C", "7C", "8C"}
    white := []string{"7H", "3H", "4H", "5H", "6H"}
    Game.Compare(black, white)
    AssertTrue(t, Game.Winner == "B", "")   
}

func TestStraightFlush(t *testing.T) {
    testXCategoryStraightFlush(t)
    testSameCategoryStraightFlush(t)
     
}
