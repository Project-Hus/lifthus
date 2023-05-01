import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { setupSwagger } from './util/swagger';
import envbyjson from 'envbyjson';

/**
 * sets up the nestjs app and returns it.
 *
 * @returns nestjs application
 */
export async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  setupSwagger(app);

  // set CORS
  app.enableCors({
    origin: [
      'http://localhost:3000',
      'https://*.lifthus.com',
      'https://lifthus.com',
    ],
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    preflightContinue: false,
    optionsSuccessStatus: 204,
  });
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
