import { Inject, Injectable } from '@nestjs/common';
import { Post, PostLike, Prisma } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { CreatePostDto, UpdatePostDto } from './post.dto';
import { slugify } from 'src/common/utils/utils';
import { SLUG_MAX_LENGTH } from 'src/common/constraints';

@Injectable()
export class PostService {
  constructor(@Inject(PrismaService) private readonly prisma: PrismaService) {}

  createPost({
    post,
    imageSrcs,
  }: {
    post: CreatePostDto;
    imageSrcs: string[];
  }): Promise<Post> {
    // first, set the range of slug and get it.
    let slug: string;
    const slugEnd: number = post.content.indexOf('\n');
    if (slugEnd == -1 || slugEnd > SLUG_MAX_LENGTH) {
      slug = post.content.slice(0, SLUG_MAX_LENGTH);
    } else {
      slug = post.content.slice(0, slugEnd);
    }
    // get slug
    slug = slugify(slug);

    let images: Prisma.PostImageCreateWithoutPostInput[] = [];
    imageSrcs.forEach((location, idx) => {
      images.push({
        src: location,
        order: idx,
      });
    });

    // Post create form
    let data: Prisma.PostCreateInput = {
      author: BigInt(post.author),
      slug,
      content: post.content,
      images: {
        create: images,
      },
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
  deletePost({
    aid,
    pid,
  }: {
    aid: number;
    pid: number;
  }): Prisma.PrismaPromise<Prisma.BatchPayload> {
    return this.prisma.post.deleteMany({
      where: { id: pid, author: aid },
    });
  }

  likePost({ uid, pid }: { uid: number; pid: number }): Promise<number> {
    const likePost = this.prisma.postLike.create({
      data: { user: uid, post: { connect: { id: pid } } },
    });
    const increasePostLikeNum = this.prisma.post.update({
      data: { likenum: { increment: 1 } },
      where: { id: pid },
    });
    return this.prisma
      .$transaction([likePost, increasePostLikeNum])
      .then((res) => {
        // if liking is successful, return the new likenum
        return res[1].likenum;
      })
      .catch(async (err) => {
        // if failed, try unliking
        const unlikePost = this.prisma.postLike.delete({
          where: { postId_user: { user: uid, postId: pid } },
        });
        const decreasePostLikeNum = this.prisma.post.update({
          data: { likenum: { decrement: 1 } },
          where: { id: pid },
        });
        return this.prisma
          .$transaction([unlikePost, decreasePostLikeNum])
          .then((res) => {
            // if unliking is successful, return the new likenum
            return res[1].likenum;
          })
          .catch((err) => {
            return Promise.reject(err);
          });
      });
  }
}
