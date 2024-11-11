package puppy

import "mymodules/mymodules/dog"

func Bark() string {
	return ("Bho!!")
}

func Barks() string {
	return ("Bho Bho !")
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
