# API REST Vidéo

Ce projet est un test technique

## Prérequis

- Go 1.21 ou supérieur
- Une instance de base de données PostgreSQL en cours d'exécution

## Installation

1. Clonez ce dépôt dans votre espace de travail Go :

```bash
git clone https://github.com/Sebastienpanda/videos.git
cd videos
```

2. Installez les dépendances du projet :

```bash
go mod download
```

3. Copiez le fichier .env.test en .env et mettez à jour les variables d'environnement selon votre configuration :

```bash
cp .env.test .env
```

## Configuration de la base de données

Ce projet utilise PostgreSQL comme base de données. Assurez-vous d'avoir une instance de PostgreSQL en cours d'exécution et mettez à jour le fichier .env avec les informations de connexion à la base de données.

### Exécution du projet

Pour effectuer les migrations, utilisez la commande suivante

```bash
go run migrate/migrate.go
```

### Pour lancer le projet

Pour lancer le projet, utilisez la commande suivante

```bash
CompileDaemon -command="./videos"
```

## Endpoints

- `GET /videos` : Récupère la liste de toutes les vidéos
- `GET /videos/:id` : Récupère une vidéo spécifique par ID
- `POST /videos` : Crée une nouvelle vidéo
- `PUT /videos/:id` : Met à jour une vidéo spécifique par ID
- `DELETE /videos/:id` : Supprime une vidéo spécifique par ID
