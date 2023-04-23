const esbuild = require('esbuild');
const path = require('node:path');

const cwd = process.cwd();
const outfile = path.resolve(cwd, 'services/post/dist-bundle/lambda.js');
const entryPoints = [path.resolve(cwd, 'services/post/dist/lambda.js')];
const config = {
    platform: 'node',
    target: ['node18'],
    bundle: true,
    keepNames: true,
    plugins: [],
    entryPoints,
    outfile,
    external: [
        "@nestjs/microservices",
        "@nestjs/websockets/socket-module",
        "cache-manager",
        "class-transformer",
        "class-validator"
    ]
};

(async () => {
    await esbuild.build(config);
})();