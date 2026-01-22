# Groupie Tracker

Un site web simple pour explorer les artistes et leurs informations de base.

## Fonctionnalités

- Affichage des artistes avec leurs informations principales
- Page de détails pour chaque artiste
- Barre de recherche basique
- Design responsive

## Technologies

- **Backend** : Go (Golang)
- **Frontend** : HTML, CSS
- **API** : Groupie Trackers API

## Prérequis

- Go 1.16 ou version ultérieure


## Installation

1. Clonez le dépôt :
   git clone https://github.com/votre-utilisateur/groupie-tracker.git
   cd groupie-tracker

2. Lancez le serveur :
    go run main.go

3. Accédez à l'application à l'adresse : http://localhost:8080


**Structure du projet**


groupie-tracker/
├── static/
│   └── style.css
├── templates/
│   ├── index.html
│   ├── artist.html
│   └── error.html
├── internal/
│   └── app/
│       ├── app.go
│       ├── data.go
│       ├── handlers.go
│       └── types.go
├── main.go
└── README.md
