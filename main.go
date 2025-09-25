package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"mini-crm/store"
)

var dataStore store.Storer

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
			fmt.Print("Nom: ")
			nom, _ := reader.ReadString('\n')

			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')

			contact := store.Contact{
				Nom:   strings.TrimSpace(nom),
				Email: strings.TrimSpace(email),
			}
			if err := dataStore.Add(contact); err != nil {
				fmt.Println("Erreur lors de l'ajout du contact:", err)
			} else {
				fmt.Println("Contact ajouté !")
			}

		case 2: // Lister
			contacts, err := dataStore.GetAll()
			if err != nil {
				fmt.Println("Erreur lors de la récupération des contacts:", err)
				continue
			}
			fmt.Println("--- Liste des contacts ---")
			if len(contacts) == 0 {
				fmt.Println("Aucun contact trouvé.")
				continue
			}
			for _, c := range contacts {
				fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
			}

		case 3: // Chercher
			fmt.Print("Entrez l'ID du contact : ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			contact, err := dataStore.Get(uint(id))
			if err != nil {
				fmt.Println("Contact non trouvé.")
			} else {
				fmt.Printf("ID: %d, Nom: %s, Email: %s\n", contact.ID, contact.Nom, contact.Email)
			}

		case 4: // Mettre à jour
			fmt.Print("Entrez l'ID du contact à mettre à jour : ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			contact, err := dataStore.Get(uint(id))
			if err != nil {
				fmt.Println("Contact non trouvé.")
				continue
			}

			fmt.Print("Nouveau nom (laisser vide pour garder actuel) : ")
			nom, _ := reader.ReadString('\n')
			nom = strings.TrimSpace(nom)

			fmt.Print("Nouvel email (laisser vide pour garder actuel) : ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			if nom != "" {
				contact.Nom = nom
			}

			if email != "" {
				contact.Email = email
			}

			if err := dataStore.Update(contact); err != nil {
				fmt.Println("Erreur lors de la mise à jour du contact:", err)
			} else {
				fmt.Println("Contact mis à jour !")
			}

		case 5: // Supprimer
			fmt.Print("Entrez l'ID du contact à supprimer : ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			if err := dataStore.Delete(uint(id)); err != nil {
				fmt.Println("Erreur lors de la suppression du contact:", err)
			} else {
				fmt.Println("Contact supprimé !")
			}

		case 6: // Quitter
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Option invalide")
		}
	}
}

func main() {
	var err error
	dataStore, err = store.NewGORMStore("contacts.db")
	if err != nil {
		log.Fatal("Erreur lors de la création du GORMStore:", err)
	}

	// Si aucun argument n'est passé, lancer le mode interactif
	if len(os.Args) < 2 {
		startInteractiveMode()
		return
	}

	// Sinon, utiliser la logique des flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
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
		if *addNom == "" || *addEmail == "" {
			fmt.Println("Les flags -nom et -email sont requis pour ajouter un contact.")
			addCmd.Usage()
			return
		}
		contact := store.Contact{Nom: *addNom, Email: *addEmail}
		if err := dataStore.Add(contact); err != nil {
			fmt.Println("Erreur lors de l'ajout du contact:", err)
		} else {
			fmt.Println("Contact ajouté !")
		}
	case "list":
		listCmd.Parse(os.Args[2:])
		contacts, err := dataStore.GetAll()
		if err != nil {
			fmt.Println("Erreur lors de la récupération des contacts:", err)
			return
		}
		fmt.Println("--- Liste des contacts ---")
		if len(contacts) == 0 {
			fmt.Println("Aucun contact trouvé.")
			return
		}
		for _, c := range contacts {
			fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
		}
	case "search":
		searchCmd.Parse(os.Args[2:])
		if *searchId == 0 {
			fmt.Println("Le flag -id est requis pour chercher un contact.")
			searchCmd.Usage()
			return
		}
		contact, err := dataStore.Get(uint(*searchId))
		if err != nil {
			fmt.Println("Contact non trouvé.")
		} else {
			fmt.Printf("ID: %d, Nom: %s, Email: %s\n", contact.ID, contact.Nom, contact.Email)
		}
	case "update":
		updateCmd.Parse(os.Args[2:])
		if *updateId == 0 {
			fmt.Println("Le flag -id est requis pour mettre à jour un contact.")
			updateCmd.Usage()
			return
		}
		contact, err := dataStore.Get(uint(*updateId))
		if err != nil {
			fmt.Println("Contact non trouvé.")
			return
		}
		if *updateNom != "" {
			contact.Nom = *updateNom
		}
		if *updateEmail != "" {
			contact.Email = *updateEmail
		}
		if err := dataStore.Update(contact); err != nil {
			fmt.Println("Erreur lors de la mise à jour du contact:", err)
		} else {
			fmt.Println("Contact mis à jour !")
		}
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteId == 0 {
			fmt.Println("Le flag -id est requis pour supprimer un contact.")
			deleteCmd.Usage()
			return
		}
		if err := dataStore.Delete(uint(*deleteId)); err != nil {
			fmt.Println("Erreur lors de la suppression du contact:", err)
		} else {
			fmt.Println("Contact supprimé !")
		}
	default:
		fmt.Println("Commande inconnue. Commandes disponibles: add, list, search, update, delete")
	}
}
