.PHONY: build

build:
    # building nest js app
	rm -rf services/post/dist
	tsc -p services/post/tsconfig.build.json
	# rename 's/\.js$$/.mjs/' ./services/post/dist/*.js
	cp services/post/package.json services/post/dist
	sam build

build-win:
	rmdir /s services/post/dist
	tsc -p services/post/tsconfig.build.json
	copy services/post/package.json services/post/dist
	sam build

start:
	# make start DEBUG=--debug
	sam local start-api --env-vars env.json -p 9091 $(DEBUG)