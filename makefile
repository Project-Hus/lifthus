DEFAULT_GOAL := build

build:
	# Remove previous build artifacts
	#rm -rf services/post/dist
	#rm -rf services/post/dist-bundle

	# Transpile and bundle PostService
	tsc -p services/post/tsconfig.build.json
	cp services/post/package.json services/post/dist
	npm --prefix services/post/dist install ./services/post/dist
	node services/post/esbuild.js

	sam build
.PHONY: build

build-win:
	tsc -p services/post/tsconfig.build.json
	copy "services/post/package.json" "services/post/dist"
	npm install --prefix ./services/post/dist ./services/post/dist
	node services/post/esbuild.js

	sam build
.PHONY: build-win

# make start DEBUG=--debug
start:
	sam local start-api --env-vars env.json -p 9100 $(DEBUG)
.PHONY: start