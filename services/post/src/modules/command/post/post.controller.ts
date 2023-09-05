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
import { PostService } from './post.service';
import { Post2Service } from 'src/modules/command/post/post2.service';
import { Post as PPost, Prisma } from '@prisma/client';
import { CreatePostDto, UpdatePostDto } from './post.dto';
import { FilesInterceptor } from '@nestjs/platform-express';
import crypto from 'crypto';
import aws from 'aws-sdk';

import multerS3 from 'multer-s3';
import { S3Service } from './s3.service';
import { Post as DPost } from 'src/domain/aggregates/post/post.model';

const s3 = new aws.S3();

/**
 * Mutation Controller
 * @description this controller is responsible for handling all mutation requests for posts.
 * @class MutationController
 */
@Controller('/post/post')
export class PostController {
  constructor(
    @Inject(PostService) private readonly postService: PostService,
    @Inject(Post2Service) private readonly post2Service: Post2Service,
    @Inject(S3Service) private readonly s3Service: S3Service,
  ) {}

  /**
   * generates new post by the form data if the user is signed.
   * @param req
   * @returns
   */
  @UseGuards(UserGuard)
  @Post()
  @UseInterceptors(
    FilesInterceptor('images', 5, {
      storage: multerS3({
        s3: s3,
        bucket: 'lifthus-post-bucket',
        acl: 'public-read',
        contentType: multerS3.AUTO_CONTENT_TYPE,
        key: function (req, file, cb) {
          cb(
            null,
            `post/images/${Date.now()}_${crypto
              .randomBytes(4)
              .toString('hex')}_${file.originalname}`,
          );
        },
      }),
    }),
  )
  createPost(
    @Req() req: Request,
    @Body() post: CreatePostDto,
    @UploadedFiles() images: Array<Express.Multer.File>,
  ): Promise<DPost> {
    this.s3Service.uploadImages(images);
    const uid: bigint = req.uid;
    const author: number = parseInt(post.author);
    if (BigInt(author) !== uid) throw new ForbiddenException();
    return this.post2Service.createPost({
      post,
      imageSrcs: images.map((image) => image.location),
    });
  }

  /**
   * updates the post by the form data if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Put()
  updatePost(
    @Req() req: Request,
    @Body() post: UpdatePostDto,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    const uid: bigint = req.uid;
    const aid: number = post.author;
    // if the author is not signed user, return 403 Forbidden.
    if (uid !== BigInt(aid)) throw new ForbiddenException();
    return this.postService.updatePost(post);
  }

  /**
   * deletes the post by the pid in the body if the user is signed.
   * @param req
   * @param pid
   * @returns
   */
  @UseGuards(UserGuard)
  @Delete('/:pid')
  deletePost(
    @Req() req: Request,
    @Param('pid') pid: any,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.postService.deletePost({
      aid: Number(req.uid),
      pid: Number(pid),
    });
  }

  /**
   * likes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/like/:pid')
  likePost(@Req() req: Request, @Param('pid') pid: any): Promise<number> {
    return this.postService.likePost({
      uid: Number(req.uid),
      pid: Number(pid),
    });
  }

  // /**
  //  * unlikes the post by the pid in the body if the user is signed.
  //  * @param req
  //  * @param post
  //  * @returns
  //  */
  // @UseGuards(UserGuard)
  // @Post('/unlike/:pid')
  // unlikePost(
  //   @Req() req: Request,
  //   @Param('pid') pid: number,
  // ): Promise<[PostLike, PPost]> {
  //   return this.postService.unlikePost(req.uid, pid);
  // }
}
