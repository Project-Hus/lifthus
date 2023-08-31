import { Injectable } from '@nestjs/common';
import { Like } from '../domain/aggregates/like/like.model';
import { User } from '../domain/aggregates/user/user.model';
import { LikeRepository } from '../domain/repositories/like.repository';
import { PrismaService } from 'src/prisma/prisma.service';

import { Post } from '../domain/aggregates/post/post.model';

@Injectable()
export abstract class PrismaPostLikeRepository extends LikeRepository<Post> {
  constructor(private readonly prismaService: PrismaService) {
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
    return new Like(u, t, !!postLike);
  }

  async _getLikeNum(t: Post): Promise<number> {
    return this.prismaService.postLike.count({
      where: {
        postId: t.getID(),
      },
    });
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
                    postId: like.target.getID(),
                    user: like.liker.getID(),
                  },
                },
              })
            : this.prismaService.postLike.create({
                data: {
                  postId: like.target.getID(),
                  user: like.liker.getID(),
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
