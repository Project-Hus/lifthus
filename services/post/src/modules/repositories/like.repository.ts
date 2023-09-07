import { Inject, Injectable, Logger } from '@nestjs/common';
import { Like } from '../../domain/aggregates/like/like.model';
import { User } from '../../domain/aggregates/user/user.model';
import { LikeRepository } from './abstract/like.repository';
import { PrismaService } from 'src/prisma/prisma.service';

import { Post } from '../../domain/aggregates/post/post.model';
import { Comment } from '../../domain/aggregates/comment/comment.model';

@Injectable()
export class PrismaPostLikeRepository extends LikeRepository<Post> {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
  ) {
    super();
  }

  async _getLike(uid: bigint, pid: bigint): Promise<Like<Post>> {
    const postLike = await this.prismaService.postLike.findUnique({
      where: {
        postId_user: {
          postId: pid,
          user: uid,
        },
      },
    });
    return Like.create(uid, pid, !!postLike);
  }

  async _getLikesNum(pid: bigint): Promise<number> {
    const n = await this.prismaService.postLike.count({
      where: {
        postId: pid,
      },
    });
    return n;
  }

  async _save(like: Like<Post>): Promise<Like<Post>> {
    try {
      const liker = like.getLiker();
      const target = like.getTarget();
      const liked = like.isLiked();
      const postLike = await this.prismaService.postLike.findUnique({
        where: {
          postId_user: {
            postId: target,
            user: liker,
          },
        },
      });
      if ((!!postLike && liked) || (!postLike && !liked)) {
        // same state, do nothing
      } else if (postLike === null && liked) {
        await this.prismaService.postLike.create({
          data: {
            postId: target,
            user: liker,
          },
        });
      } else if (!!postLike && !liked) {
        await this.prismaService.postLike.delete({
          where: {
            postId_user: {
              postId: target,
              user: liker,
            },
          },
        });
      }
      return like;
    } catch (e) {
      return Promise.reject(e);
    }
  }
}

@Injectable()
export class PrismaCommentLikeRepository extends LikeRepository<Comment> {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
  ) {
    super();
  }

  async _getLike(uid: bigint, cid: bigint): Promise<Like<Comment>> {
    const commentLike = await this.prismaService.commentLike.findUnique({
      where: {
        commentId_user: {
          commentId: cid,
          user: uid,
        },
      },
    });
    return Like.create(uid, cid, !!commentLike);
  }

  async _getLikesNum(cid: bigint): Promise<number> {
    return await this.prismaService.commentLike.count({
      where: {
        commentId: cid,
      },
    });
  }

  async _save(like: Like<Comment>): Promise<Like<Comment>> {
    try {
      const liker = like.getLiker();
      const target = like.getTarget();
      const liked = like.isLiked();
      const postLike = await this.prismaService.commentLike.findUnique({
        where: {
          commentId_user: {
            commentId: target,
            user: liker,
          },
        },
      });
      if ((!!postLike && liked) || (!postLike && !liked)) {
        // same state, do nothing
      } else if (postLike === null && liked) {
        await this.prismaService.commentLike.create({
          data: {
            commentId: target,
            user: liker,
          },
        });
      } else if (!!postLike && !liked) {
        await this.prismaService.commentLike.delete({
          where: {
            commentId_user: {
              commentId: target,
              user: liker,
            },
          },
        });
      }
      return like;
    } catch (e) {
      return Promise.reject(e);
    }
  }
}
