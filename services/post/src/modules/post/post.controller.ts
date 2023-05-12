import {
  Body,
  Controller,
  Delete,
  Param,
  Post,
  Put,
  Req,
  UseGuards,
} from '@nestjs/common';
import { UserGuard } from 'src/common/guards/post.guard';
import { Request } from 'express';
import { PostService } from './post.service';

/**
 * Mutation Controller
 * @description this controller is responsible for handling all mutation requests for posts.
 * @class MutationController
 */
@Controller('/post/post')
export class PostController {
  constructor(private readonly postService: PostService) {}

  /**
   * generates new post by the form data if the user is signed.
   * @param req
   * @returns
   */
  @UseGuards(UserGuard)
  @Post()
  wirtePost(@Req() req: Request, @Body() post: any): any {
    const uid: number = req.uid;
    return; //this.appService.post(uid, post);
  }

  /**
   * updates the post by the form data if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Put()
  updatePost(@Req() req: Request, @Body() post: any): any {
    // also check if the post belongs to the user
    const uid: number = req.uid;
    return; //this.appService.put(uid, post);
  }

  /**
   * deletes the post by the pid in the body if the user is signed.
   * @param req
   * @param pid
   * @returns
   */
  @UseGuards(UserGuard)
  @Delete()
  deletePost(@Req() req: Request, @Param('pid') pid: string): any {
    // also check if the post belongs to the user
    const uid: number = req.uid;
    return; //this.appService.delete(uid, pid);
  }

  /**
   * likes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/like')
  likePost(@Req() req: Request, @Body() post: any): any {
    // also check if the post belongs to the user
    const uid: number = req.uid;
    return; //this.appService.like(uid, post);
  }

  /**
   * unlikes the post by the pid in the body if the user is signed.
   * @param req
   * @param post
   * @returns
   */
  @UseGuards(UserGuard)
  @Post('/unlike')
  unlikePost(@Req() req: Request, @Body() post: any): any {
    // also check if the post belongs to the user
    const uid: number = req.uid;
    return; //this.appService.unlike(uid, post);
  }
}
