.PHONY: build

build:
    # building nest js app
	rm -rf services/post/dist
	tsc -p services/post/tsconfig.build.json
	cp services/post/package.json services/post/dist
	sam build

start:
	# make start DEBUG=--debug
	sam local start-api --env-vars env.json -p 9091 $(DEBUG)