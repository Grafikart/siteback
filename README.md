# siteback

[![Licence: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

L'objectif de ce projet est d'avoir un outil simple qui sauvegarder un projet (seulement la base de données pour le moment) en se basant sur les configurations du projet.

## Utilisation

Ne gère pour le moment qu'un projet symfony avec un backup S3 sur s3.fr-par.scw.cloud.

```bash
siteback <bucket>
```

## Fonctionnement

- Lit le fichier .env.local
- Génère un dump MySQL (se connecte via les infos de DATABASE_URL)
- Gzip le dump
- Upload sur S3 en se basant sur S3_KEY et S3_SECRET

## A améliorer

- Le code (c'est un premier jet)
- Gérer les erreurs dans les goroutines
- Ecrire des tests
- Gérer d'autres méthodes de sauvegarde (FTP, SSH...)
- Gérer autre chose que symfony
- Sauvegarder aussi des dossiers

## 📝 License

This project is [MIT](https://choosealicense.com/licenses/mit/) licensed.
