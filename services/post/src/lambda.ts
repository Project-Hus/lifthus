// lambda.ts
import { bootstrap } from './main';
import serverlessExpress from '@vendia/serverless-express';
import { Callback, Context, Handler } from 'aws-lambda';

let server: Handler;

export const handler = async (
  event: any,
  context: Context,
  callback: Callback,
) => {
  const app = await bootstrap();
  await app.init();
  const expressApp = app.getHttpAdapter().getInstance();

  server = server ?? serverlessExpress({ app: expressApp });
  return server(event, context, callback);
};
