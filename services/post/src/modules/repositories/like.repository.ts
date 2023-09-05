import { Inject, Injectable, Logger } from '@nestjs/common';
import { Like } from '../../domain/aggregates/like/like.model';
import { User } from '../../domain/aggregates/user/user.model';
import { LikeRepository } from '../../domain/repositories/like.repository';
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

  async _getLike(u: User, t: Post): Promise<Like<Post>> {
    const postLike = await this.prismaService.postLike.findUnique({
      where: {
        postId_user: {
          postId: t.getID(),
          user: u.getID(),
        },
      },
    });
    return Like.create(u, t, !!postLike);
  }

  async _getLikesNum(pid: bigint): Promise<number> {
    const n = await this.prismaService.postLike.count({
      where: {
        postId: pid,
      },
    });
    return n;
  }

  async _save(likes: Set<Like<Post>>): Promise<void> {
    try {
      const likesList = Array.from(likes);
      this.prismaService.$transaction(
        likesList.map((like) => {
          return like.isLiked()
            ? this.prismaService.postLike.delete({
                where: {
                  postId_user: {
                    postId: like.getTarget().getID(),
                    user: like.getLiker().getID(),
                  },
                },
              })
            : this.prismaService.postLike.create({
                data: {
                  postId: like.getTarget().getID(),
                  user: like.getLiker().getID(),
                },
              });
        }),
      );
      return;
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

  async _getLike(u: User, t: Comment): Promise<Like<Comment>> {
    const commentLike = await this.prismaService.commentLike.findUnique({
      where: {
        commentId_user: {
          commentId: t.getID(),
          user: u.getID(),
        },
      },
    });
    return Like.create(u, t, !!commentLike);
  }

  async _getLikesNum(cid: bigint): Promise<number> {
    return await this.prismaService.commentLike.count({
      where: {
        commentId: cid,
      },
    });
  }

  async _save(likes: Set<Like<Comment>>): Promise<void> {
    try {
      const likesList = Array.from(likes);
      this.prismaService.$transaction(
        likesList.map((like) => {
          return like.isLiked()
            ? this.prismaService.commentLike.delete({
                where: {
                  commentId_user: {
                    commentId: like.getTarget().getID(),
                    user: like.getLiker().getID(),
                  },
                },
              })
            : this.prismaService.commentLike.create({
                data: {
                  commentId: like.getTarget().getID(),
                  user: like.getLiker().getID(),
                },
              });
        }),
      );
      return;
    } catch (e) {
      return Promise.reject(e);
    }
  }
}
