package main

import (
    "fmt"
)

func ajouterContact(id int, nom string, email string) {
    if _, ok := contacts[id]; ok {
        fmt.Println("Un contact avec cet ID existe déjà.")
        return
    }

    contacts[id] = Contact{
        ID:    id,
        Nom:   nom,
        Email: email,
    }
    if err := saveContacts(); err != nil {
        fmt.Println("Erreur lors de la sauvegarde des contacts:", err)
    }
    fmt.Println("Contact ajouté !")
}

func listerContacts() {
    fmt.Println("--- Liste des contacts ---")
    if len(contacts) == 0 {
        fmt.Println("Aucun contact trouvé.")
        return
    }
    for _, c := range contacts {
        fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
    }
}

func chercherContact(id int) {
    if c, ok := contacts[id]; ok {
        fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
    } else {
        fmt.Println("Contact non trouvé.")
    }
}

func mettreAJourContact(id int, nom string, email string) {
    c, ok := contacts[id]
    if !ok {
        fmt.Println("Contact non trouvé.")
        return
    }

    if nom != "" {
        c.Nom = nom
    }

    if email != "" {
        c.Email = email
    }

    contacts[id] = c
    if err := saveContacts(); err != nil {
        fmt.Println("Erreur lors de la sauvegarde des contacts:", err)
    }
    fmt.Println("Contact mis à jour !")
}

func supprimerContact(id int) {
    if _, ok := contacts[id]; !ok {
        fmt.Println("Contact non trouvé.")
        return
    }

    delete(contacts, id)
    if err := saveContacts(); err != nil {
        fmt.Println("Erreur lors de la sauvegarde des contacts:", err)
    }
    fmt.Println("Contact supprimé !")
}