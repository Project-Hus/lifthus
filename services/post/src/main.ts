import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { setupSwagger } from './openapi/swagger';

import envbyjson from 'envbyjson';

import cookieParser from 'cookie-parser';

/* swagger not working in lambda environment */
// import SwaggerUi from 'swagger-ui-express';
// import { OpenapiModule } from './openapi/openapi.module';

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
      'https://www.lifthus.com',
      'https://auth.lifthus.com',
      'https://api.lifthus.com',
      'https://lifthus.com',
    ],
    allowedHeaders: [
      'Origin',
      'Content-Type',
      'Authorization',
      'Accept',
      'X-Requested-With',
      'Access-Control-Allow-Origin',
      'HeaderAccessControlAllowHeaders',
      'HeaderAccessControlAllowMethods',
      'HeaderXRequestedWith',
    ],
    credentials: true,
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    //preflightContinue: false,
    //optionsSuccessStatus: 204,
  });

  app.use(cookieParser());

  /* swagger not working in lambda environment */
  // const openapi = await NestFactory.create(OpenapiModule);
  // const openapiDoc = await setupSwagger(openapi);
  // app.use('/post/openapi', SwaggerUi.serve, SwaggerUi.setup(openapiDoc));

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
  // run the nestjs server if the environment is native
  if (process.env.HUS_ENV === 'native') {
    envbyjson.loadProp('../../env.json', 'Parameters');
    console.log('native nestsjs running');
    run(Number(process.env.POST_PORT));
  }
} catch (e) {
  console.log(e);
}
