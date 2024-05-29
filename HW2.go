package main

import (
	"fmt"
)

type Entity struct {
    Name string
}

type Animal struct {
    Entity
}

type Cage struct {
    Number int
    Animal *Animal
}

type Zookeeper struct {
    Entity
    Caught []*Animal
}

func (z *Zookeeper) Catch(animal *Animal) {
    fmt.Printf("%s зловив %s\n", z.Name, animal.Name)
    z.Caught = append(z.Caught, animal)
}

func (z *Zookeeper) ReturnToCages(cages []*Cage) {
    for _, animal := range z.Caught {
        emptyCage := getEmptyCage(cages)
        if emptyCage != nil {
            emptyCage.Animal = animal
            fmt.Printf("%s повернув %s до клітки %d\n", z.Name, animal.Name, emptyCage.Number)
        }
    }
    
    z.Caught = []*Animal{}
}

func getEmptyCage(cages []*Cage) *Cage {
    for _, cage := range cages {
        if cage.Animal == nil {
            return cage
        }
    }
    return nil
}

func main() {

    animals := []*Animal{
        {Entity{Name: "Жираф"}},
        {Entity{Name: "Тигр"}},
        {Entity{Name: "Вовк"}},
        {Entity{Name: "Слон"}},
        {Entity{Name: "Зебра"}},
    }

    cages := []*Cage{
        {Number: 1},
        {Number: 2},
        {Number: 3},
        {Number: 4},
        {Number: 5},
    }

    zookeeper := Zookeeper{Entity: Entity{Name: "Роман"}}

    for _, animal := range animals {
        zookeeper.Catch(animal)
    }

    zookeeper.ReturnToCages(cages)

    for _, cage := range cages {
        if cage.Animal != nil {
            fmt.Printf("У клітці %d знаходиться %s\n", cage.Number, cage.Animal.Name)
        } else {
            fmt.Printf("Клітка %d порожня\n", cage.Number)
        }
    }
}
