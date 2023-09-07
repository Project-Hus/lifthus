import {
  Body,
  Controller,
  Delete,
  ForbiddenException,
  Inject,
  Logger,
  Param,
  Post,
  Put,
  Req,
  UploadedFiles,
  UseGuards,
  UseInterceptors,
} from '@nestjs/common';
import { UserGuard } from 'src/common/guards/post.guard';
import { Request } from 'express';
import { PostService } from 'src/modules/command/post/post.service';
import { FilesInterceptor } from '@nestjs/platform-express';
import { PostDto } from 'src/dto/outbound/post.dto';
import { Uid } from 'src/common/decorators/authParam.decorator';
import { getMulterS3Option } from 'src/common/multerS3/multerS3';
import {
  CreatePostRequestDto,
  CreatePostServiceDto,
  UpdatePostRequestDto,
  UpdatePostServiceDto,
} from 'src/dto/inbound/post.dto';

@Controller('/post/post')
export class PostController {
  constructor(
    @Inject(PostService) private readonly post2Service: PostService,
  ) {}

  @UseGuards(UserGuard)
  @Post()
  @UseInterceptors(FilesInterceptor('images', 5, getMulterS3Option()))
  createPost(
    @Uid() clientId: bigint,
    @Body() postForm: CreatePostRequestDto,
    @UploadedFiles() images: Array<Express.Multer.File>,
  ): Promise<PostDto> {
    const post = new CreatePostServiceDto(postForm, images);
    return this.post2Service.createPost({
      clientId,
      post,
    });
  }

  @UseGuards(UserGuard)
  @Put()
  updatePost(
    @Uid() clientId: bigint,
    @Body() postForm: UpdatePostRequestDto,
  ): Promise<PostDto> {
    const postUpdates = new UpdatePostServiceDto(postForm);
    return this.post2Service.updatePost({ clientId, postUpdates });
  }

  /**
   * deletes the post by the pid in the body if the user is signed.
   * @param req
   * @param pid
   * @returns
   */
  @UseGuards(UserGuard)
  @Delete('/:pid')
  deletePost(@Uid() clientId, @Param('pid') pid: string): Promise<PostDto> {
    return this.post2Service.deletePost({ clientId, pid: BigInt(pid) });
  }
}
