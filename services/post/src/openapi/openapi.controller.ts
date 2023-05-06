import { Controller, Get, Header } from '@nestjs/common';
import { OpenapiService } from './openapi.service';

@Controller()
export class OpenapiController {
  constructor(private readonly openapiService: OpenapiService) {}

  @Get('/post/openapi')
  getHello(): string {
    console.log('YOYOYO OPENAPI');
    return this.openapiService.getSwaggerHTML();
  }
}
