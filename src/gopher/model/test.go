package model

type gopher struct {
        name string
        age  int
        isAdult bool
}

type horse struct {
        name string
        weight int
}

func (g gopher) Jump() string {
        if g.age < 65 {
                return g.name + " can jump HIGH"
        }
        return g.name + " with age " + string(g.age) + " can still jump"
}

func (h horse) Jump() string {
        if h.weight > 2500 {
                return "I'm too heavy, can't jump..."
        }
        return "I will jump, neigh!!"
}

type jumper interface {
        Jump() string
}

func GetLists() []jumper {
        gopher1 := gopher{name: "Phil", age: 30}
        gopher2 := gopher{name: "John", age: 90}
        horse1 := horse{name: "Barbaro", weight: 2000}
        list := []jumper{&gopher1, &gopher2, &horse1}
        return list
}

