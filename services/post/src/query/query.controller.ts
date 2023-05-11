import { Controller, Get, Res, Req, UseGuards } from '@nestjs/common';
import { Request, Response } from 'express';
import { UserGuard } from 'src/common/guards/post.guard';
import { QueryService } from './query.service';

@Controller('/post/query')
export class QueryController {
  constructor(private readonly gueryService: QueryService) {}

  @Get()
  getHello(): string {
    return this.gueryService.getHello();
  }

  @UseGuards(UserGuard)
  @Get('/signed')
  getUid(@Req() req: Request) {
    return `Yo! ${req.uid}`;
  }

  @Get('/cookie')
  getCookie(@Res() res: Response) {
    res.send(this.gueryService.getCookie(res));
  }

  @Get('/one/:id')
  getOne(): string {
    return 'Hello World!';
  }
}
