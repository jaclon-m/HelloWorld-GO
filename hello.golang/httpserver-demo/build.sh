hub=${HUB-jaclond}
image=${IMAGE-httpserver-lgs:0.0.2}
#hub=${HUB-idefav}
#image=${IMAGE-httpserver:0.0.1}
docker build . -t "${hub}/${image}"