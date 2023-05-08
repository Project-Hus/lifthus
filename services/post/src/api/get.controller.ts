import { Controller, Get } from '@nestjs/common';
import { GetService } from './get.service';

@Controller('/post/get')
export class GetController {
  constructor(private readonly getService: GetService) {}

  @Get()
  getHello(): string {
    return this.getService.getHello();
  }

  @Get('/one/:id')
  getOne(): string {
    return 'Hello World!';
  }
}
