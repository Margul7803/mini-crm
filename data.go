package main

import (
    "encoding/json"
    "os"
)

func saveContacts() error {
    data, err := json.MarshalIndent(contacts, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile("contacts.json", data, 0644)
}

func loadContacts() error {
    data, err := os.ReadFile("contacts.json")
    if err != nil {
        if os.IsNotExist(err) {
            return nil // Le fichier n'existe pas encore, c'est normal
        }
        return err
    }
    return json.Unmarshal(data, &contacts)
}
