# Part IV : En vrac

Quelques derniers axes d'améliorations en vrac, qui avaient pas leur place dans une autre partie, et qui sont trop courts pour que je leur en réserve une.

## Index

- [Part IV : En vrac](#part-iv--en-vrac)
  - [Index](#index)
  - [I. Clean caches](#i-clean-caches)
  - [II. Labels](#ii-labels)
  - [III. No root](#iii-no-root)
    - [1. Intro](#1-intro)
  - [2. Utilisateur applicatif](#2-utilisateur-applicatif)
  - [IV. Going further](#iv-going-further)

## I. Clean caches

> C'est l'heure pour notre image de faire une sèche.

A chaque fois qu'on télécharge un truc avec un outil dédié (comme les gestionnaires de paquets `apt`, `dnf`, `apk`, `pacman`, etc.), y'a généralement du *cache* qui est créé quelque part, ou des fichiers d'index, ce genre de trucs.

**Des choses donc qui ne sont pas utiles au moment où on run notre application** (le *runtime*).

On va donc les dégager pour gagner un peu de place !

> **Ne négligez jamais un gain de 1Mo qui pourrait vous paraître insignifiant.** Ce Mo il est stocké à chaque fois sur votre machine, là ou vous pushez votre image, il transite et va transiter sur le réseau des dizaines de fois, voire beaucoup plus. On est des milliers (millions ?) de dévs. La multiplication fait vite mal au crâne. Puis l'efficience et l'élégance, en informatique, ont toujours une valeur inestimable. Que tu sois le patron qui veut économiser des euros, le tech qui veut optimiser à la microseconde ou l'écolo qui veut économiser de l'énergie, dégage moi le moindre octet superflu stp.

🌞 **Modifiez le `Dockerfile`**

- clean tous les caches !
- renseignez-vous suivant l'image que vous avez
- `apt` et `apk` génèrent du cache par exemple !
- les gestionnaires de libs comme `npm` ou `pip` c'est récurrent aussi qu'ils laissent des machins derrière eux

## II. Labels

On peut stocker des metadatas associées à notre image à l'aide l'instruction `LABEL` dans le `Dockerfile`.

Une image est toujours accompagnée d'une tonne de metadonnées (vous pouvez `docker image inspect <IMAGE>` pour toutes les voir).

Il existe des standards autour de la conteneurisation émis par l'[OCI](https://opencontainers.org/). L'OCI a définit une spécification pour des images, et dans la spec, on trouve notamment [les `LABEL`s standards](https://specs.opencontainers.org/image-spec/annotations/) pour maximiser l'interopérabilité, la traçabilité, les bails quoi

🌞 **Ajoutez une clause `LABEL` à votre `Dockerfile`** qui précise, **avec les `LABEL`s standards** :

- l'auteur de l'image (pseudo, nom, peu importe)
- l'URL vers le `Dockerfile` d'origine et le code (votre dépôt git quoi)
- un vendor : l'entité qui a la propriété intellectuelle sur le truc

> Vous pouvez mettre des fake values partout si vous voulez. Peu importe, j'veux des `LABEL`s standards :d

## III. No root

### 1. Intro

> Je le mets en dernier et ça me fait mal au coeur.

➜ **BON CA VA LA DE TOUT FAIRE TOURNER EN ROOT C'EST FATIGANT A LA FIN**

Mais cette fois, c'est pas votre faute : **Docker fait tout tourner en `root` par défaut** le boug', et c'est pas jonti.

> Je pourrai m'étaler autant que vous voulez sur les dangers énormes que ça implique, just ask. Si chacun d'entre nous pouvait affirmer qu'il n'introduit JAMAIS aucune vulnérabilités dans son code, et que toutes les *supply chain* étaient parfaites, et que les moules avaient des gants, ça n'aurait pas d'impact. C'est pas le cas.

➜ **Concrètement, si on utilise pas la clause `USER` dans nos `Dockerfile`s alors par défaut, le conteneur s'exécutera en `root`.**

Quand on crée un utilisateur qui est dédié à exécuter une seule application, on voit parfois le terme "utilisateur applicatif" pur le désigner.

> C'est la base de la base de la base de la sécu. Si la gestion de droits et de users est OK, on se prémunit de 99% des attaques, exploitations, impact d'un bug, etc.

➜ **Ptit rappel de système en passant** : quand un utilisateur fait un truc qui génère/crée des fichiers, par défaut, ces fichiers lui appartiendront. Autrement dit, c'est pas exclu que tu sois obligé de passer un ou deux ptit `RUN chown -R` dans ton `Dockerfile`. A votre dispo pour des rappels sur ça si besoin !

## 2. Utilisateur applicatif

🌞 **Ajoutez un utilisateur applicatif à votre `Dockerfile`**

- il faut un `RUN` qui lance un `useradd` afin de créer un utilisateur dans l'image
  - perso j'le nomme comme mon app
  - le user `meow` lance l'application `meow`, straightforward quoi
- spécifiez ensuite une ligne `USER` qui indique que tout le reste du `Dockerfile` c'est lui doit l'exécuter (et pas `root`)

> J'insiste : **à partir de la clause `USER` tout est exécuté en tant que cet utilisateur.** Il faut donc évidemment que ce soit avant le `CMD` ou `ENTRYPOINT` pour que le conteneur soit pas lancé en `root` au moment du `docker run`. Il faut aussi que ce soit après les `apt` parce que ça nécessite les droits `root` d'installer des paquets. Placez-le donc judicieusement dans le fichier.

## IV. Going further

On va s'arrêter là pour ce TP, mais il y a beacoup de lecture sur internet pour aller plus loin.

Go chercher comme optimiser/sécuriser/good-practicer vos `Dockerfile`s sur internet, il reste quelques pistes (comme par exemple minimiser le nombre de *layers*), et aussi des pistes spécifiques à vos langages/frameworks.
