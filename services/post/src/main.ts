import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { setupSwagger } from './util/swagger';

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

// for local native development
async function run() {
  const nestApp = await bootstrap();
  await nestApp.listen(9092);
}

//run();
