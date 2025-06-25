# TP3 : Docker et CI/CD

Dans ce TP on va voir **un des cas d'utilisation majeur de Docker : les *pipelines de CI/CD*.**

L'idée est de déclencher des actions automatiques à chaque fois qu'on fait un `git push` sur un dépôt Git.

On pense par exemple à :

- test de ***lint***
- lancement de **tests unitaires/intégration/etc**
- tests liés à la **sécurité**
- si tous les tests passent...
  - **on effectue un *build*** final (`docker build` pour nous)
  - **déploiement automatique** de la nouvelle version sur un serveur distant

> On finit vite par avoir 10000 trucs déclenchés automatiaquement au moindre *push* sur un dépôt Git de code : vérifier que le code est correct en terme de syntaxe, correspond aux bonnes pratiques, ne comporte pas de vulnérabilités connues, de secrets en clair, etc. :d

![DevSecOps 2024 requirements](./img/2024_requirements.jpeg)

## Sommaire

- [TP3 : Docker et CI/CD](#tp3--docker-et-cicd)
  - [Sommaire](#sommaire)
  - [Prérequis](#prérequis)
  - [Part 1 : Première pipeline](#part-1--première-pipeline)
  - [Part 2 : Test then Build](#part-2--test-then-build)
  - [Part 3 : Deploy the world](#part-3--deploy-the-world)

## Prérequis

➜ **Les commandes `git` et `docker` dispos sur votre PC**

➜ **Compte Azure**

- avec vos identifiants de l'école et l'offre Azure for Student activée
- prêt à créer une machine

## [Part 1 : Première pipeline](./part1.md)

Dans cette première partie, on met en place une première pipeline de CI.

Au menu : on va utiliser **Gitlab** pour **lancer des tâches automatiquement** à chaque push sur un repo.

Pour bosser, **vous continuez sur le dépôt git du TP2**, y'a tout ce qu'il nous faut dessus.

> Juste poussez le sur Gitlab s'il était sur une autre plateforme avant. Si tu te sens chaud, j'ai aucun soucis à ce que tu fasses tout avec Github Actions, ou une autre plateforme de CI/CD, mais je donnerai les instructions spécifiques pour Gitlab.

[**Part 1 : Première pipeline**](./part1.md)

![Gitlab](./img/logo_gitlab.png)

## [Part 2 : Test then Build](./part2.md)

On continue avec **les premiers morceaux d'une chaîne d'automatisation clean**. Une partie dédiée à **la CI**.

> *CI* pour *Continuous Integration* (ou *Intégration Continue*) : le fait d'effectuer des actions automatiques à chaque push sur un dépôt git.

On effectue des **tests automatiquement** sur le code, et on **refuse le push si les tests ne sont pas validés**.

Si les tests sont validés, on déclenche un **build automatique du code**. Ici, on déclenchera le build d'une image Docker.

[**Part 2 : Test then Build**](./part2.md)

## [Part 3 : Deploy the world](./part3.md)

Troisième partie : **on met en place la CD.**

> *CD* pour *Continuous Deployment* (ou *Déploiement Continu*) : le fait de déployer le code automatiquement sur une machine à la fin des tests automatisés.

Dans le TP, on va réutiliser votre ~magnifique~ compte Azure pour pop une VM avec une IP publique. **Ce sera notre environnement de *"production"*.**

[**Part 3 : Deploy the world**](./part3.md)
