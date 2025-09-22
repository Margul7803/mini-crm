package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\n--- Mini-CRM ---")
        fmt.Println("1. Ajouter un contact")
        fmt.Println("2. Lister tous les contacts")
        fmt.Println("3. Chercher un contact par ID")
        fmt.Println("4. Mettre à jour un contact")
        fmt.Println("5. Quitter")
        fmt.Print("Choisissez une option : ")

        choiceStr, _ := reader.ReadString('\n')
        choiceStr = strings.TrimSpace(choiceStr)
        choice, err := strconv.Atoi(choiceStr)
        if err != nil {
            fmt.Println("Veuillez entrer un nombre valide")
            continue
        }

        switch choice {
        case 1:
            ajouterContact(reader)
        case 2:
            listerContacts()
        case 3:
            chercherContact(reader)
        case 4:
            mettreAJourContact(reader)
        case 5:
            fmt.Println("Au revoir !")
            return
        default:
            fmt.Println("Option invalide")
        }
    }
}
