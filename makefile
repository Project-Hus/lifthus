.PHONY: build

build:
    # building nest js app
	tsc -p services/lifthus-post/tsconfig.build.json
	sam build

start:
	sam local start-api --env-vars env.json -p 9091

start --debug:
	sam local start-api --env-vars env.json -p 9091 --debug