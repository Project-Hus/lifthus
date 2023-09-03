DEFAULT_GOAL := build

# OLD MACBOOK BUILD METHOD
# build:
# 	# Remove previous build artifacts
# 	#rm -rf services/post/dist
# 	#rm -rf services/post/dist-bundle

# 	# Transpile and bundle PostService
# 	tsc -p services/post/tsconfig.build.json
# 	cp services/post/package.json services/post/dist
# 	npm --prefix services/post/dist install ./services/post/dist
# 	cp services/post/prisma/schema.prisma services/post/dist-bundle
# 	cp services/post/node_modules/.prisma/client/libquery_engine-rhel-openssl-1.0.x.so.node services/post/dist-bundle
# 	node services/post/esbuild.js

# 	sam build
# .PHONY: build

build:
	sam build
	cp services/post/prisma/schema.prisma .aws-sam/build/DevLifthusPostService
	cp services/post/node_modules/.prisma/client/libquery_engine-rhel-openssl-1.0.x.so.node .aws-sam/build/DevLifthusPostService
.PHONY: build

# make start DEBUG=--debug
start:
	sam local start-api --env-vars env.json -p 9100 $(DEBUG)
.PHONY: start