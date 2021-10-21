hub=${jaclond}
image=${httpserver-lgs:0.0.2}
docker build . -t "${hub}/${image}"