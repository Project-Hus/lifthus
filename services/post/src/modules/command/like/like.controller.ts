import { Controller, Inject, Param, Post, UseGuards } from '@nestjs/common';
import { Uid } from 'src/common/decorators/authParam.decorator';
import { UserGuard } from 'src/common/guards/post.guard';
import { LikeDto } from 'src/dto/outbound/like.dto';
import { LikeService } from 'src/modules/command/like/like.service';

@Controller('/post/like')
export class LikeController {
  constructor(@Inject(LikeService) private readonly likeService: LikeService) {}

  @UseGuards(UserGuard)
  @Post('post/:pid')
  likePost(@Uid() clientId, @Param('pid') pid: string): Promise<LikeDto> {
    return this.likeService.likePost({ clientId, pid: BigInt(pid) });
  }

  @UseGuards(UserGuard)
  @Post('comment/:cid')
  likeComment(@Uid() clientId, @Param('cid') cid: string): Promise<LikeDto> {
    return this.likeService.likeComment({ clientId, cid: BigInt(cid) });
  }
}
