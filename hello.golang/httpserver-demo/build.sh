hub=${HUB-jaclon}
image=${IMAGE-httpserver:0.0.1}
docker build . -t "${hub}/${image}"