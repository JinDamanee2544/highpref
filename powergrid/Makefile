fileName := rand-40-160

pyrun:
	python3 ./main.py input/${fileName} output/${fileName}.out

build:
	docker build . -t powergrid

dockerRemove:
	docker rm powergrid | true

dockerRun:
	time docker run \
	--name powergrid \
	-v ${PWD}/data/input:/input \
	-v ${PWD}/data/output:/output \
	powergrid \
	prog /input/grid-6-7 /output/grid-6-7

update: build dockerRemove dockerRun