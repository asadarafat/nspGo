package main

import (
	nspgousecase "local.com/nspgo/nspGo-useCase"
)

func main() {

	// nspgousecase.ThalesLookupWithIntentCreateIntents()
	// nspgousecase.ThalesLookupWithIntentGetIntents()
	// nspgousecase.ThalesLookupWithResourceManagerObtain(1, 1000000)
	// nspgousecase.ThalesLookupWithResourceManagerRelease()

	nspgousecase.ThalesLookupWithGoCache(1, 1000000)

}
