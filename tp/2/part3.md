# Part III : Base image

Dans cette partie on s'attarder un peu sur **le choix de l'image de base.**

➜ **Déjà : la sécu.** On veut s'assurer que notre image est pas **toute pourrie** (pas remplie de vulns critiques). Pour ce qui est de **la provenance**, dans l'idéal on veut une image conçue par les dévs du truc (notion de confiance).

> Genre si on utilise une image `debian` on aimerait qu'elles soient fournies par les gars de chez Debian. Idem pour n'importe quelle autre image, comme une image `ubuntu` ou `symfony`.

➜ Pour **une bonne maîtrise de l'environnement**, on va souvent préférer prendre une image contenant juste un OS basique, et ajouter nous-mêmes notre langage/framework/libs/autres dépendances. Ca permet aussi de choisir parfois des configs exotiques, mais nécessaires.

> C'pas rare que les entreprises éditent elles-mêmes des images internes. Attention quand ça provient de l'extérieur en tout cas !

➜ Parmi les images officielles les plus utilisées comme base, on trouve notamment `debian` et `alpine`.

> L'image officielle `ubuntu` est pas mal utilisée aussi. Et d'autres. Mais y'a un peu la question de `alpine` versus the world.

➜ **Dans cette partie, on va donc se concentrer sur les premières lignes du `Dockerfile`, en particulier le `FROM`.**

## 1. Provenance

Vous pourrez trouver ces infos en partant de [la WebUI du Docker Hub](https://hub.docker.com/).

🌞 **Donnez le lien vers les `Dockerfile`s opensource qui permettent de produire**

- l'image `debian:latest`
- l'image `alpine:latest`
- l'image que vous aviez utilisé comme `FROM` jusqu'à maintenant

> Z'avez vu ce `FROM scratch` dans les `Dockerfile` de `alpine` ou `debian` ? Clin d'oeil clin d'oeil. 

## 2. Vulnérabilités connues

Vous pensez que l'image `debian:latest` officielle (par exemple) contient combien de vulnérabilités connues ?

Ptite partie pour répondre à cette question, j'vous donne juste le tool, et vous lancez juste pour avoir la réponse !

Le tool que je vous recommande c'est [`trivy`](https://github.com/aquasecurity/trivy), outil efficace et plutôt mature.

> Il existe plein de tools pour ça, certains avec des WebUI sexy. J'aime toujours beaucoup les outils en CLI parce que ça s'intègre à tout et n'importe quoi, des hooks git, des pipelines de CI/CD, ou n'importe quoi d'autres : c'est trop facile à wrapper dans autre chose.

🌞 **Déterminer le nombre et la criticité des vulns connues dans les images**

- `debian:latest`
- `alpine:latest`
- l'image que vous aviez utilisé comme `FROM` jusqu'à maintenant

> Note : la criticité c'est genre "low" "medium" "high" "critical" par exemple. [Des termes définis de façon plus ou moins standard](https://nvd.nist.gov/vuln-metrics).

## 3. Dockerfile writing

➜ **Ecrire un `Dockerfile-alpine`**

- en partant du `Dockerfile` de la partie précédente (il faut donc modifier le début)
- commence par un `FROM alpine:xxx`
  - vous remplacez la `xxx` par la dernière version explicitement ([***version pinning***](https://jonathan.bergknoff.com/journal/always-pin-your-versions/))
  - pas de `latest`
- s'en suit sûrement un ou plusieurs `RUN` pour installer :
  - votre langage
  - votre tooling (genre un gestionnaire de paquet pour télécharger des dépendances)
  - d'autres dépendances si besoin
- la fin du `Dockerfile-alpine` devraient être identique à `Dockerfile`
- pas besoin qu'il soit dans le dépôt de rendu ce fichier

➜ **Ecrire un `Dockerfile-debian`**

- pareil mais avec l'image `debian` officielle
- pas besoin qu'il soit dans le dépôt de rendu ce fichier

## 4. Measure !

🌞 **Build time**

- deux commandes `docker build` chronométrée dans le rendu, utilisées pour build ces deux `Dockerfile`s
- au moment du build, elles doivent être nommées décemment ces deux images svp
- utilisez un truc pour chronométrer le temps du `build` en terminal :

```bash
# Linux et MacOS : y'a la commande time
time docker build ...

# Windows : ça s'fait avec Measure-Command
Measure-Command { docker build ... }
```

🌞 **Comparaison post-build**

- comparez la taille des deux images
  - avec une simple commande `docker`
- comparez la perf des deux images
  - ça va très largement dépendre de votre app ça
  - ça serait bien de timer combien de temps prend l'app à démarrer, à fonctionner, à traiter une requête un peu lourde, ce genre de truc
  - le but est de mettre en évidence (ou pas) la lenteur de `alpine`

## 3. Choose yourself !

➜ **Choisissez ce que vous préférez pour continuer le TP**

- l'un de ces deux `Dockerfile`s devient votre `Dockerfile` pour la suite du TP, j'veux pas entendre parler de l'autre !

> Si on avait pas utilisé de multi-stage build, il aurait fallu éditer deux `Dockerfile` là. Relou. Ou oublier de le faire. Encore plus relou.

---

➜ [**Dernière partie avec quelques tips en vrac**](./part4.md)
