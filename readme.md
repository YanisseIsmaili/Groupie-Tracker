### Le projet GROUPIE TRACKER est réalisé par ISMAILI YANISSE et YAPI THÉAU dans le cadre de leur première année au sein de MONTPELLIER YNOV CAMPUS.

---

### Le projet consiste à récupérer les données JSON d'une API en ligne ("https://groupietrackers.herokuapp.com/api") et à les afficher dans un format stylisé sur un site internet hébergé localement.

---

Pour ce faire, les données sont d'abord récupérées depuis l'API à l'aide d'un programme en GOLANG qui utilise la librairie "encoding/json". Le programme lit le contenu de l'API, retient toutes les données JSON et les encode ensuite en données HTML selon une structure GOLANG.

Ensuite, les données sont intégrées aux pages HTML selon des balises du type : "{{.Name}}, {{.Id}}, {{.Members}}, etc".

Enfin, pour afficher le site en local, le programme main.go héberge le site sur le port 8080 de l'ordinateur, ce qui permet d'y accéder en entrant l'adresse suivante dans un navigateur internet : http://localhost:8080/.

Le site présente quelques fonctionnalités supplémentaires, comme une barre de recherche pour retrouver plus facilement les données des artistes. Cette barre de recherche fonctionne selon le principe de "ctrl+F", c'est-à-dire qu'elle recherche la syntaxe exacte de ce qui est écrit sur le site, mais uniquement pour les artistes. C'est pourquoi les artistes sont répertoriés dans la sidebar. Il suffit donc de taper le nom exact de l'artiste pour afficher uniquement ses informations.

Il y a également quelques easter-eggs sur le site pour apporter une touche personnelle ;) .