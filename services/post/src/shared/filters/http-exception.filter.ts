import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpException,
  UnauthorizedException,
  Logger,
} from '@nestjs/common';
import { Request, Response } from 'express';

@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    const status = exception.getStatus();

    if (status === 401) {
      response.status(status).send('expired_token');
    } else {
      response.status(status).json({
        status: status,
        timestamp: new Date().toISOString(),
        path: request.url,
        message: exception.message,
      });
    }
  }
}
