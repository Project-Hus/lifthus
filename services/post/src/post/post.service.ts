import { Injectable } from '@nestjs/common';

import SwaggerUi from 'swagger-ui-express';

@Injectable()
export class PostService {
  getHello(): string {
    return 'Hello World!';
  }
}
