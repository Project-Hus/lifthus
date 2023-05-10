import { Controller, Get, Res } from '@nestjs/common';
import { GetService } from './get.service';
import { Response } from 'express';

@Controller('/post/get')
export class GetController {
  constructor(private readonly getService: GetService) {}

  @Get()
  getHello(): string {
    return this.getService.getHello();
  }

  @Get('/cookie')
  getCookie(@Res() res: Response) {
    res.send(this.getService.getCookie(res));
  }

  @Get('/one/:id')
  getOne(): string {
    return 'Hello World!';
  }
}
