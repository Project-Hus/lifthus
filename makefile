.PHONY: build

build:
	# Remove previous build artifacts
	rm -rf services/post/dist
	rm -rf services/post/dist-bundle

	# Transpile and bundle PostService
	tsc -p services/post/tsconfig.build.json
	cp services/post/package.json services/post/dist
	npm --prefix services/post/dist install ./services/post/dist
	node services/post/esbuild.js

	sam build

build-win:
	tsc -p services/post/tsconfig.build.json
	copy "services/post/package.json" "services/post/dist"
	npm install --prefix ./services/post/dist ./services/post/dist
	node services/post/esbuild.js

	sam build

start:
	# make start DEBUG=--debug
	sam local start-api --env-vars env.json -p 9091 $(DEBUG)