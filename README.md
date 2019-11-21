```
██╗███████╗███████╗██╗  ██╗ █████╗ ██╗
██║██╔════╝██╔════╝██║ ██╔╝██╔══██╗██║
██║███████╗█████╗  █████╔╝ ███████║██║
██║╚════██║██╔══╝  ██╔═██╗ ██╔══██║██║
██║███████║███████╗██║  ██╗██║  ██║██║
╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝
```
# bifrost

## purpose
Provides a headless executable for running Weave programs like EarthRock. Useful for servers and development.

## log
 - [X] Static HTTP server
 - [X] Single Page Application Support
 - [X] Prevent directories from showing up via FileServer
 - [ ] Event Source Channel
 - [ ] POST Channel
 - [ ] Host off //bifrost.earthrock.run
    - Keep using CDNs for earthrock.run and isekai.run
    - Might have to get off github pages? Need more info

## future
 - [ ] Run server side weaves

This is meant to be the glue that holds the isekai clients together. Should only implement basic nodes and only run validation on views. Ultimately the games are meant to be ran and played on the clients but the bifrost can provide validation to ensure the integrity of games being played. 
