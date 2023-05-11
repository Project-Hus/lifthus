import { Injectable } from '@nestjs/common';

import SwaggerUi from 'swagger-ui-express';

@Injectable()
export class PostMutationService {
  getHello(): string {
    return 'Hello World!';
  }
}
