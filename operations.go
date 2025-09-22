package main

import (
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

func ajouterContact(reader *bufio.Reader) {
    fmt.Print("ID: ")
    idStr, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idStr))

    if _, ok := contacts[id]; ok {
        fmt.Println("Un contact avec cet ID existe déjà.")
        return
    }

    fmt.Print("Nom: ")
    nom, _ := reader.ReadString('\n')
    fmt.Print("Email: ")
    email, _ := reader.ReadString('\n')

    contacts[id] = Contact{
        ID:    id,
        Nom:   strings.TrimSpace(nom),
        Email: strings.TrimSpace(email),
    }
    fmt.Println("Contact ajouté !")
}

func listerContacts() {
    fmt.Println("\n--- Liste des contacts ---")
    if len(contacts) == 0 {
        fmt.Println("Aucun contact trouvé.")
        return
    }
    for _, c := range contacts {
        fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
    }
}

func chercherContact(reader *bufio.Reader) {
    fmt.Print("Entrez l'ID du contact : ")
    idStr, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idStr))

    if c, ok := contacts[id]; ok {
        fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
    } else {
        fmt.Println("Contact non trouvé.")
    }
}

func mettreAJourContact(reader *bufio.Reader) {
    fmt.Print("Entrez l'ID du contact à mettre à jour : ")
    idStr, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idStr))

    if c, ok := contacts[id]; ok {
        fmt.Print("Nouveau nom (laisser vide pour garder actuel) : ")
        nom, _ := reader.ReadString('\n')
        nom = strings.TrimSpace(nom)
        if nom != "" {
            c.Nom = nom
        }

        fmt.Print("Nouvel email (laisser vide pour garder actuel) : ")
        email, _ := reader.ReadString('\n')
        email = strings.TrimSpace(email)
        if email != "" {
            c.Email = email
        }

        contacts[id] = c
        fmt.Println("Contact mis à jour !")
    } else {
        fmt.Println("Contact non trouvé.")
    }
}
