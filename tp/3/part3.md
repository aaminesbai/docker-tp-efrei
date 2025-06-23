# Part 3 : Deploy the world

Troisième partie : **on met en place la *CD*** pour que notre application soit déployée automatiquement lorsque nos premiers *stages* de *CI* sont exécutés avec succès (ce qu'on a fait en [part 2](./part2.md))

> *CD* pour *Continuous Deployment* (ou *Déploiement Continu*) : le fait de déployer le code automatiquement sur une machine à la fin des tests automatisés.

Dans le TP, on va réutiliser votre ~~magnifique~~ compte Azure pour pop une VM avec une IP publique. **Ce sera notre environnement de *"production"*.**

![Azure](./img/azure.jpg)

## Index

- [Part 3 : Deploy the world](#part-3--deploy-the-world)
  - [Index](#index)
  - [1. Prepare Azure VM](#1-prepare-azure-vm)
  - [2. Manual deploy](#2-manual-deploy)
  - [3. Auto-deploy](#3-auto-deploy)
  - [La suiiite](#la-suiiite)

## 1. Prepare Azure VM

➜ **Créez une VM dans Azure**

- OS de votre choix
- assurez-vous de pouvoir vous y connecter en SSH sur son IP publique avant de continuer

> **TOUS** les OS possèdent désormais la commande `ssh` (y compris Windows). Donc ouvrez juste un terminal sur votre PC pour cette étape !

---

Désolééé on peut pas y couper on va faire un poil de conf système. On limite au max ! Je serai ravi de (ré)expliquer plein de trucs sur le sujet si vous demandez, version courte et efficace promis. Sinon ChatGepetto lui aussi sera ravi de l'écrire cette conf, allez droit au but si besoin, pas la partie qui vous intéresse.

➜ **Configurez la VM Azure**

- installation de Docker (suivez la doc officielle)
  - `start` et `enable` le service `docker.service`
- créez un utilisateur `deploy`
  - ajoutez-le au groupe `docker`
  - il a tous les droits `sudo` en `NOPASSWD`

> L'utilisateur `deploy` sera utilisé par notre *pipeline* GitLab pour se connecter à notre machine de production. Une fois connecté sur l'utilisateur `deploy`, GitLab pourra déployer automatiquement des machins en tapant des commandes.

➜ **Connexion SSH pour l'utilisateur `deploy`**

- créez une paire de clés SSH (depuis votre machine)
- déposez la clé publique sur le compte de l'utilisateur `deploy` sur la VM Azure
- gardez la clé privée au chaud, on devra la fournir à GitLab pour qu'il puisse se connecter à la VM Azure

🌞 **Pour le compte-rendu**

- je veux voir une commande `ssh deploy@IP_PUBLIQUE_VM_AZURE` qui fonctionne **depuis VOTRE PC vers la machine Azure**
  - sans demander de password
- suivi d'un `sudo -l` qui permet d'afficher les droits `sudo`
  - on devrait voir la conf avec un `NOPASSWD`
- et enfin un `groups deploy`
  - on devrait voir que l'utilisateur `deploy` est dans le groupe `docker`
  - c'est nécessaire pour pouvoir utiliser la commande `docker`

## 2. Manual deploy

🌞 **Déployer l'image qui été build à la partie précédente en une seule commande SSH**

- une commande SSH depuis votre PC vers la machine Azure
- elle fait un `docker run` de l'image qui a été push sur le registre Gitlab
- par exemple, un truc du genre :

```bash
# oui oui on peut passer une commande directement sur la ligne ssh comme ça
ssh deploy@IP_PUBLIQUE_VM_AZURE docker run NOM_DE_LIMAGE
```

➜ **Je me doute que l'app est pétée si elle est pas lancée proprement avec un `docker compose`, on fait juste un ptit test pour éviter d'accumuler mille problèmes potentiels au moment du déploiement automatisé.**

## 3. Auto-deploy

🌞 **Modifiez votre fichier `.gitlab-ci.yml`**

- il doit comporter un *stage* `deploy`
  - le *job* se connecte à la machine Azure en SSH
  - clone votre dépôt git
  - lance votre app avec `docker-compose.yml` dans sa nouvelle version (il faut ptet couper/relancer)

🌞 **Vérifiez que ça a été correctement déployé !**

- mettez moi un `curl IP_PUBLIQUE_VM_AZURE` dans le compte-rendu

![Continuously](./img/deliver_continuously.png)

## La suiiite

👉 On continue à approfondir le sujet, et on va aborder le terme [***shift left* dans la quatrième partie.**](./part4.md)

