import { Controller, Get } from '@nestjs/common';
import { OpenapiService } from './openapi.service';

@Controller('/post/openapi')
export class OpenapiController {
  constructor(private readonly openapiService: OpenapiService) {}

  @Get()
  getSwagger(): string {
    return this.openapiService.getSwaggerHTML();
  }
}
