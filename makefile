.PHONY: build

build:
	rm -rf services/post/dist
	tsc -p services/post/tsconfig.build.json
	cp services/post/package.json services/post/dist
	npm --prefix services/post/dist install ./services/post/dist
	node services/post/esbuild.js
	sam build

build-win:
	rmdir /s services/post/dist
	tsc -p services/post/tsconfig.build.json
	copy services/post/package.json services/post/dist
	npm --prefix services/post/dist install ./services/post/dist
	sam build

start:
	# make start DEBUG=--debug
	sam local start-api --env-vars env.json -p 9091 $(DEBUG)