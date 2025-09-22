package main

type Contact struct {
    ID    int
    Nom   string
    Email string
}

var contacts = make(map[int]Contact)
