// lambda.ts
import { INestApplication } from '@nestjs/common';
import { bootstrap } from './main';
import serverlessExpress from '@vendia/serverless-express';
import { Callback, Context, Handler } from 'aws-lambda';

// Lambda optimization
let server: Handler;
let app: INestApplication;
let expressApp: any;

export const handler = async (
  event: any,
  context: Context,
  callback: Callback, // this call back is used to
) => {
  // if app is not initialized, initialize it
  if (!server || !app || !expressApp) {
    app = await bootstrap();
    await app.init();
    expressApp = app.getHttpAdapter().getInstance();
    server = server ?? serverlessExpress({ app: expressApp });
  }

  return server(event, context, callback);
};
