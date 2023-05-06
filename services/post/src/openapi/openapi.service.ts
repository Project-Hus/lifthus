import { Injectable } from '@nestjs/common';
import SwaggerUi from 'swagger-ui-express';

@Injectable()
export class OpenapiService {
  getSwaggerHTML() {
    return SwaggerUi.generateHTML();
  }
}
