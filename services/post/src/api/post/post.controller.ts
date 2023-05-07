import { Controller, Get, Header } from '@nestjs/common';
import { PostService, OpenapiService } from './post.service';

@Controller('/post')
export class PostController {
  constructor(
    private readonly appService: PostService,
    private readonly openapiService: OpenapiService,
  ) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @Get()
  @Header('Content-Type', 'text/html')
  getSwagger(): string {
    return this.openapiService.getSwaggerHTML();
  }
}
