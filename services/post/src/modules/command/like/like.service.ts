import { Inject, Injectable } from '@nestjs/common';
import { LikeDto } from 'src/dto/outbound/like.dto';
import {
  CommentLikeRepository,
  PostLikeRepository,
} from 'src/modules/repositories/abstract/like.repository';
import { UserRepository } from 'src/modules/repositories/abstract/user.repository';

@Injectable()
export class LikeService {
  constructor(
    @Inject(UserRepository)
    private readonly userRepo: UserRepository,
    @Inject(PostLikeRepository)
    private readonly postLikeRepo: PostLikeRepository,
    @Inject(CommentLikeRepository)
    private readonly commentLikeRepo: CommentLikeRepository,
  ) {}

  async likePost({
    clientId,
    pid,
  }: {
    clientId: bigint;
    pid: bigint;
  }): Promise<LikeDto> {
    const liker = this.userRepo.getUser(clientId);
    const postLike = await this.postLikeRepo.getLike(liker.getID(), pid);
    if (postLike.isLiked()) liker.unlikePost(postLike);
    else liker.likePost(postLike);
    const savedPostLike = await this.postLikeRepo.save(postLike);
    return new LikeDto(
      savedPostLike.getLiker(),
      savedPostLike.getTarget(),
      savedPostLike.isLiked(),
    );
  }

  async likeComment({
    clientId,
    cid,
  }: {
    clientId: bigint;
    cid: bigint;
  }): Promise<LikeDto> {
    const liker = this.userRepo.getUser(clientId);
    const commentLike = await this.commentLikeRepo.getLike(liker.getID(), cid);
    if (commentLike.isLiked()) liker.unlikeComment(commentLike);
    else liker.likePost(commentLike);
    const savedCommentLike = await this.commentLikeRepo.save(commentLike);
    return new LikeDto(
      savedCommentLike.getLiker(),
      savedCommentLike.getTarget(),
      savedCommentLike.isLiked(),
    );
  }
}
