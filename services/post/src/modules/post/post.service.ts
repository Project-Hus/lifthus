import { Injectable } from '@nestjs/common';
import { Post, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { CreatePostDto, UpdatePostDto } from './post.dto';
import { slugify } from 'src/common/utils/utils';

@Injectable()
export class PostService {
  constructor(private readonly prisma: PrismaService) {}

  wirtePost(post: CreatePostDto): Promise<Post> {
    let slug: string;
    // get slugEnd by the first '\n'.
    const slugEnd: number = post.content.indexOf('\n');
    // if '\n' not found or the first '\n' is after 30th character, slice the first 30 characters.
    if (slugEnd == -1 || slugEnd > 30) {
      // if automatically takes all if the content is less than 30 characters not throwing error.
      slug = post.content.slice(0, 30);
    } else {
      // if '\n' is found and it is before 31th character, slice the content until '\n'.
      slug = post.content.slice(0, slugEnd);
    }
    // get slug
    slug = slugify(slug);

    // Post create form
    let data: Prisma.PostCreateInput = {
      author: post.author,
      slug,
      content: post.content,
    };

    return this.prisma.post.create({
      data,
    });
  }

  /**
   * @param aid
   * @param pid
   * @returns {count:number} only expected 0 or 1
   */
  updatePost(data: UpdatePostDto): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.post.updateMany({
      data,
      where: { id: data.id },
    });
  }

  /**
   *
   * @param aid
   * @param pid
   * @returns {count:number} only expected 0 or 1
   */
  deletePost(
    aid: number,
    pid: number,
  ): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.post.deleteMany({
      where: { id: pid, author: aid },
    });
  }

  likePost(uid: number, where: Prisma.PostWhereUniqueInput): Promise<Post> {
    return this.prisma.postLike
      .create({
        data: { user: uid, post: { connect: where } },
      })
      .then((res) => {
        return this.prisma.post.update({
          data: { likenum: { increment: 1 } },
          where,
        });
      });
  }

  unlikePost(uid: number, where: Prisma.PostWhereUniqueInput): Promise<Post> {
    return this.prisma.postLike
      .delete({
        where: { postId_user: { user: uid, postId: where.id } },
      })
      .then((res) => {
        return this.prisma.post.update({
          data: { likenum: { decrement: 1 } },
          where,
        });
      });
  }
}
