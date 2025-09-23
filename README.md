# EDT-AMU-Minfo

Ce projet d'initiative étudiante est un mini site permettant de filtrer les cours affichés sur l'emploi du temps en fonction des différents groupes sélectionnés par l'utilisateur.

N'hésitez pas à contribuer en ouvrant une issue, et/ou en me contactant directement.

## Build

1. [Installer go](https://go.dev/dl/)
2. Cloner le dépôt
```bash
git clone https://github.com/Vrock691/EDT-AMU-Minfo.git
```
3. Installer les dépendances
```
go mod tidy
```

4. Exécuter en mode développement
```
go run ./src
```

5. Build et exécution
```
cd src
go build -o out
./out 
```