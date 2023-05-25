// lambda.ts
import { INestApplication } from '@nestjs/common';
import { bootstrap } from './main';
import serverlessExpress from '@vendia/serverless-express';
import { Callback, Context, Handler } from 'aws-lambda';

// Lambda optimization
let server: Handler;
let app: INestApplication;
let expressApp: any;

//
// JSON.stringify doesn't work with BigInt. so it should be treated as Number.
BigInt.prototype.toJSON = function () {
  return Number(this);
};

export const handler = async (
  event: any,
  context: Context,
  callback: Callback,
) => {
  console.log('REQUEST:', event);
  // if app is not initialized, initialize it
  if (!server) {
    app = await bootstrap();
    expressApp = app.getHttpAdapter().getInstance();
    server = server ?? serverlessExpress({ app: expressApp });
  }
  const res = await server(event, context, callback);
  console.log('RESPONSE:', res);
  return res;
};
