All:
#HUB ?=jaclond
#IMAGE ?=httpserver:0.0.1
HUB ?=jaclond
IMAGE ?=httpserver-lgs:0.0.2
build: All
	./build.sh
push: build
	docker push "${HUB}/${IMAGE}"
run: All
	docker run -d -p 8080:8080 ${HUB}/${IMAGE}


# = 是最基本的赋值
#:= 是覆盖之前的值
#?= 是如果没有被赋值过就赋予等号后面的值
#+= 是添加等号后面的值
# https://www.cnblogs.com/wanqieddy/archive/2011/09/21/2184257.html