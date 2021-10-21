hub=${HUB-jaclond}
image=${IMAGE-httpserver-lgs:1.0.1}
docker build . -t "${hub}/${image}"