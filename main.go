package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Fonction pour le mode interactif
func startInteractiveMode() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\n--- Mini-CRM (Mode Interactif) ---")
        fmt.Println("1. Ajouter un contact")
        fmt.Println("2. Lister tous les contacts")
        fmt.Println("3. Chercher un contact par ID")
        fmt.Println("4. Mettre à jour un contact")
        fmt.Println("5. Supprimer un contact")
        fmt.Println("6. Quitter")
        fmt.Print("Choisissez une option : ")

        choiceStr, _ := reader.ReadString('\n')
        choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
        if err != nil {
            fmt.Println("Veuillez entrer un nombre valide")
            continue
        }

        switch choice {
        case 1: // Ajouter
            fmt.Print("ID: ")
            idStr, _ := reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(idStr))

            fmt.Print("Nom: ")
            nom, _ := reader.ReadString('\n')

            fmt.Print("Email: ")
            email, _ := reader.ReadString('\n')

            ajouterContact(id, strings.TrimSpace(nom), strings.TrimSpace(email))

        case 2: // Lister
            listerContacts()

        case 3: // Chercher
            fmt.Print("Entrez l'ID du contact : ")
            idStr, _ := reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(idStr))
            chercherContact(id)

        case 4: // Mettre à jour
            fmt.Print("Entrez l'ID du contact à mettre à jour : ")
            idStr, _ := reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(idStr))

            fmt.Print("Nouveau nom (laisser vide pour garder actuel) : ")
            nom, _ := reader.ReadString('\n')

            fmt.Print("Nouvel email (laisser vide pour garder actuel) : ")
            email, _ := reader.ReadString('\n')

            mettreAJourContact(id, strings.TrimSpace(nom), strings.TrimSpace(email))

        case 5: // Supprimer
            fmt.Print("Entrez l'ID du contact à supprimer : ")
            idStr, _ := reader.ReadString('\n')
            id, _ := strconv.Atoi(strings.TrimSpace(idStr))
            supprimerContact(id)

        case 6: // Quitter
            fmt.Println("Au revoir !")
            return

        default:
            fmt.Println("Option invalide")
        }
    }
}

func main() {
    if err := loadContacts(); err != nil {
        fmt.Println("Erreur lors du chargement des contacts:", err)
        os.Exit(1)
    }

    // Si aucun argument n'est passé, lancer le mode interactif
    if len(os.Args) < 2 {
        startInteractiveMode()
        return
    }

    // Sinon, utiliser la logique des flags
    addCmd := flag.NewFlagSet("add", flag.ExitOnError)
    addId := addCmd.Int("id", 0, "ID du contact")
    addNom := addCmd.String("nom", "", "Nom du contact")
    addEmail := addCmd.String("email", "", "Email du contact")

    listCmd := flag.NewFlagSet("list", flag.ExitOnError)

    searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
    searchId := searchCmd.Int("id", 0, "ID du contact à chercher")

    updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
    updateId := updateCmd.Int("id", 0, "ID du contact à mettre à jour")
    updateNom := updateCmd.String("nom", "", "Nouveau nom du contact")
    updateEmail := updateCmd.String("email", "", "Nouvel email du contact")

    deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
    deleteId := deleteCmd.Int("id", 0, "ID du contact à supprimer")


    switch os.Args[1] {
    case "add":
        addCmd.Parse(os.Args[2:])
        if *addId == 0 || *addNom == "" || *addEmail == "" {
            fmt.Println("Les flags -id, -nom et -email sont requis pour ajouter un contact.")
            addCmd.Usage()
            return
        }
        ajouterContact(*addId, *addNom, *addEmail)
    case "list":
        listCmd.Parse(os.Args[2:])
        listerContacts()
    case "search":
        searchCmd.Parse(os.Args[2:])
        if *searchId == 0 {
            fmt.Println("Le flag -id est requis pour chercher un contact.")
            searchCmd.Usage()
            return
        }
        chercherContact(*searchId)
    case "update":
        updateCmd.Parse(os.Args[2:])
        if *updateId == 0 {
            fmt.Println("Le flag -id est requis pour mettre à jour un contact.")
            updateCmd.Usage()
            return
        }
        mettreAJourContact(*updateId, *updateNom, *updateEmail)
    case "delete":
        deleteCmd.Parse(os.Args[2:])
        if *deleteId == 0 {
            fmt.Println("Le flag -id est requis pour supprimer un contact.")
            deleteCmd.Usage()
            return
        }
        supprimerContact(*deleteId)
    default:
        fmt.Println("Commande inconnue. Commandes disponibles: add, list, search, update, delete")
    }
}
