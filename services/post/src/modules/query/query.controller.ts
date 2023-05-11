import { Controller, Get, Res, Req, UseGuards } from '@nestjs/common';
import { Request, Response } from 'express';
import { UserGuard } from 'src/common/guards/post.guard';
import { PostQueryService } from './post.query.service';

@Controller('/post/query')
export class QueryController {
  constructor(private readonly postQueryService: PostQueryService) {}

  @Get()
  getHello(): string {
    return this.postQueryService.getHello();
  }

  @UseGuards(UserGuard)
  @Get('/signed')
  getUid(@Req() req: Request) {
    return `Yo! ${req.uid}`;
  }

  @Get('/cookie')
  getCookie(@Res() res: Response) {
    res.send(this.postQueryService.getCookie(res));
  }

  @Get('/one/:id')
  getOne(): string {
    return 'Hello World!';
  }
}
