# Mini-CRM CLI

Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go.

## Fonctionnalités
*   **Gestion complète des contacts (CRUD)** : Ajouter, Lister, Mettre à jour et Supprimer des contacts.
*   **Interface en ligne de commande** : Commandes et sous-commandes claires et standardisées grâce à Cobra.
*   **Configuration externe** : Le comportement de l'application (notamment le type de stockage) peut être modifié sans recompiler grâce à Viper.
*   **Persistance des données** : Support de multiples backends de stockage :
    *   **GORM/SQLite** : Une base de données SQL robuste contenue dans un simple fichier (`contacts.db`).
    *   **Fichier JSON** : Une sauvegarde simple et lisible (`contacts.json`).
    *   **En mémoire** : Un stockage éphémère pour les tests.

## Installation et utilisation

1.  **Clonez le dépôt** :
    ```bash
    git clone <URL_DU_REPO>
    cd mini-crm
    ```

2.  **Téléchargez les dépendances** :
    ```bash
    go mod tidy
    ```

3.  **Construisez l'application** :
    ```bash
    go build -o mini-crm-cli
    ```

4.  **Utilisation** :
    L'application peut être configurée via un fichier `config.yaml` à la racine du projet ou dans votre répertoire personnel (`~/.mini-crm.yaml`).

    Exemple de `config.yaml` :
    ```yaml
    storage:
      type: gorm # ou 'json', ou 'memory'
      path: contacts.db # Ignoré si type est 'memory'
    ```

    **Commandes disponibles** :
    ```bash
    ./mini-crm-cli --help
    ```

    *   **Ajouter un contact** :
        ```bash
        ./mini-crm-cli add --nom "John Doe" --email "john.doe@example.com"
        ```

    *   **Lister tous les contacts** :
        ```bash
        ./mini-crm-cli list
        ```

    *   **Mettre à jour un contact** :
        ```bash
        ./mini-crm-cli update --id 1 --nom "Jane Doe" --email "jane.doe@example.com"
        ```

    *   **Supprimer un contact** :
        ```bash
        ./mini-crm-cli delete --id 1
        ```