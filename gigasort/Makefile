size := 10000000
sizeName := data1B

trace:
	go tool trace ./profile/trace.out

build:
	go build .

timeRun:
	time ./gigasort \
	/input/${sizeName}.txt \
	/output/${sizeName}.txt \

all: build timeRun

buildDocker:
	docker build -t sorting .

dockerRun:
	time docker run \
	--name sorting \
	-v ${PWD}/data/input:/input \
	-v ${PWD}/data/output:/output \
	sorting \
	prog /input/${sizeName}.txt /output/${sizeName}.txt

dockerRemove:
	docker rm sorting | true

update: buildDocker dockerRemove dockerRun

genTest:
	python3 input/gen-sort.py rand ${size} > input/${sizeName}.txt