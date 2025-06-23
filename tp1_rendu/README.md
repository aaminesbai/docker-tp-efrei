# TP1 : Docker ez oupa - Rendu

Ce dossier contient mes réponses aux trois parties du TP1.

## Part I : Init

1. **Installation de Docker**
   - Suivi des instructions officielles.
2. **Utilisation sans sudo**
   ```bash
   sudo usermod -aG docker $USER
   newgrp docker
   docker ps
   ```
3. **Premier conteneur**
   ```bash
   docker run -p 8000:8000 it4lik/meow-api
   docker ps
   docker logs <id>
   docker inspect <id>
   ```
4. **Mode détaché et logs**
   ```bash
   docker run -d -p 8000:8000 it4lik/meow-api
   docker logs <id>
   ```
5. **Remplacement du code par volume**
   ```bash
   docker run -v $(pwd)/part1/custom_app.py:/app/app.py -p 8000:8000 it4lik/meow-api
   curl http://localhost:8000/
    # Hello from custom app!
   ```
6. **Variable d'environnement**
   ```bash
   docker run -d -e LISTEN_PORT=16789 -p 16789:16789 it4lik/meow-api
   curl http://localhost:16789/
    # {"message":"Available routes","routes":{"list_all_users":"http://localhost:16789/users","get_user_by_id":"http://localhost:16789/user/1"}}
   ```

## Part II : Images

1. **Récupération d'images**
   ```bash
   docker pull python:3.11
   docker pull mysql:8.0.42
   docker pull wordpress:latest
   docker pull linuxserver/wikijs:latest
   docker image ls
   ```
2. **Conteneur Python**
   ```bash
   docker run -it python:3.11 bash
   python --version
   ```
3. **Build de meow-api**
   ```bash
   cd tp/1/app
   docker build . -t meow-api
   docker image ls
   docker run -p 8000:8000 meow-api
   ```
4. **Application emoji**
   ```bash
   cd part2/emoji_app
   docker build . -t python_app:version_de_ouf
   docker image ls
   docker run python_app:version_de_ouf
   ```
5. **Mon application perso**
   ```bash
   cd part2/my_flask_app
   docker build . -t my_flask_app:latest
   docker run -p 5000:5000 my_flask_app:latest
   ```
   Publication sur Docker Hub :
   ```bash
   docker tag my_flask_app:latest <monuser>/my_flask_app:latest
   docker push <monuser>/my_flask_app:latest
   ```

## Part III : Compose

1. **compose\_test**
   ```bash
   mkdir compose_test
   cd compose_test
   # docker-compose.yml contenant deux conteneurs debian
   docker compose up -d
   docker compose ps
   ```
   Depuis `conteneur_nul` : installation de `iputils-ping` puis `ping conteneur_flopesque`.
2. **Meow compose**
   Dossier `meow_compose/` (voir contenu dans ce repo).
   ```bash
   cd meow_compose
   docker compose up --build -d
    # Starting db...
    # Starting meow-api...

   curl http://localhost:${LISTEN_PORT}/users
    # [{"id":1,"name":"Alice","favorite_insult":"you fool"}, {"id":2,"name":"Bob","favorite_insult":"clown"}, {"id":3,"name":"Charlie","favorite_insult":"dummy"}, {"id":4,"name":"Diana","favorite_insult":"nerd"}, {"id":5,"name":"Eve","favorite_insult":"baka"}]
   curl http://localhost:${LISTEN_PORT}/user/3
    # {"id":3,"name":"Charlie","favorite_insult":"dummy"}
   ```
