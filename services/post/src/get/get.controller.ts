import { Controller, Get, Res, Req, UseGuards } from '@nestjs/common';
import { GetService } from './get.service';
import { Request, Response } from 'express';
import { UserGuard } from 'src/common/guards/post.guard';

@Controller('/post/get')
export class GetController {
  constructor(private readonly getService: GetService) {}

  @Get()
  getHello(): string {
    return this.getService.getHello();
  }

  @UseGuards(UserGuard)
  @Get('/signed')
  getUid(@Req() req: Request) {
    return `Yo! {req.uid}`;
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
