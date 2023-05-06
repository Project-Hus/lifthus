import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { setupSwagger } from './openapi/swagger';
import envbyjson from 'envbyjson';

import SwaggerUi from 'swagger-ui-express';
import { OpenapiModule } from './openapi/openapi.module';

/**
 * sets up the nestjs app and returns it.
 *
 * @returns nestjs application
 */
export async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  // set CORS
  app.enableCors({
    origin: [
      'http://localhost:3000',
      'https://*.lifthus.com',
      'https://lifthus.com',
    ],
    allowedHeaders: [
      'origin',
      'content-type',
      'authorization',
      'accept',
      'x-requested-with',
      'access-Control-Allow-Origin',
      'headerAccessControlAllowHeaders',
      'headerAccessControlAllowMethods',
      'headerXRequestedWith',
    ],
    credentials: true,
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    //preflightContinue: false,
    //optionsSuccessStatus: 204,
  });

  const openapi = await NestFactory.create(OpenapiModule);
  const openapiDoc = await setupSwagger(openapi);
  app.use('/post/openapi', SwaggerUi.serve, SwaggerUi.setup(openapiDoc));

  await app.init();

  return app;
}

/**
 *  if the environment is native, it will be called instead of the lambda handler.
 */
async function run(port: number) {
  const nestApp = await bootstrap();
  await nestApp.listen(port);
}

try {
  // load the environment variables from the env.json file
  envbyjson.loadProp('../../env.json', 'Parameters');
  // run the nestjs server if the environment is native
  if (process.env.HUS_ENV === 'native') {
    console.log('native nestsjs running');
    run(Number(process.env.POST_PORT));
  }
} catch (e) {
  console.log(e);
}
