# Building

A to-do item is to move this procedure to github actions, but for the time being here is how to build and deploy locally.

1. Merge PRs into the main branch
1. Commit any local changes
1. Tag the current head in YY.MM.DD format
1. Check for existing folder /deploy/dgraph, if present delete it
1. In the /tools folder, run the `build.sh` script
1. In the /deploy folder, run `make all`. Note you'll need to be logged into the structureorg account on Dockerhub for this step to work

