package search

import{
	"log"
	"sync"
}


var matchers = make(map[string]Matcher)


