package main

import "testing"


func TestAnimalAction(t *testing.T) {
    // Create a Dog instance and test AnimalAction
    dog := Dog{}
    expectedDog := "s ays: Woof!\nmoves by: Running\n"

    if result := AnimalAction(dog); result != expectedDog {
        t.Errorf("Dog action was incorrect, got: %s, want: %s", result, expectedDog)
    }

    // Create a Cat instance and test AnimalAction
    cat := Cat{}
    expectedCat := "says: Meow!\nmoves by: Walking\n"

    if result := AnimalAction(cat); result != expectedCat {
        t.Errorf("Cat action was incorrect, got: %s, want: %s", result, expectedCat)
    }
}

