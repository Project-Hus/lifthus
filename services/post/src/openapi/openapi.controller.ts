import { Controller, Get, Header } from '@nestjs/common';
import { OpenapiService } from './openapi.service';

@Controller('nonono/post/openapi')
export class OpenapiController {
  constructor(private readonly openapiService: OpenapiService) {}

  @Get()
  getSwagger(): string {
    return this.openapiService.getSwaggerHTML();
  }
}
