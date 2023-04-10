// main.ts
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

export async function bootstrap() {
  const app = await NestFactory.create(AppModule);
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

async function run() {
  const nestApp = await bootstrap();
  await nestApp.listen(3000);
}

run();
