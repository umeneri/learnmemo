#!/bin/bash

# in api/
# cd ../frontend
# yarn build
# cd ../api
docker build . -f docker/production/api/Dockerfile -t learnmemo
docker tag learnmemo registry.heroku.com/learnmemo/web
docker push registry.heroku.com/learnmemo/web
heroku container:release web

