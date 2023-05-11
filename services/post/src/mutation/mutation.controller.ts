import { Controller, Post, Req, UseGuards } from '@nestjs/common';
import { UserGuard } from 'src/common/guards/post.guard';
import { Request } from 'express';
import { MutationService } from './mutation.service';

@Controller('/post/mutation')
export class MutationController {
  constructor(private readonly appService: MutationService) {}

  @UseGuards(UserGuard)
  @Post()
  post(@Req() req: Request): string {
    return ''; //this.appService.post();
  }
}
