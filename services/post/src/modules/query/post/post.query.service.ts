import { Injectable } from '@nestjs/common';
import { Post as PrismaPost } from '@prisma/client';
import { Post } from 'src/modules/domain/aggregates/post/post.model';
import { PostSummary } from 'src/modules/domain/aggregates/post/postSummary.model';
import { User } from 'src/modules/domain/aggregates/user/user.model';
import { PrismaPostRepository } from 'src/modules/repositories/post.repository';
import { UserRepository } from 'src/modules/repositories/user.repository';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class PostQueryService {
  constructor(
    private readonly prismaService: PrismaService,
    private readonly userRepo: UserRepository,
    private readonly postRepo: PrismaPostRepository,
    private readonly userModel: User,
  ) {}
  getHello(): string {
    return 'Hello World!';
  }

  async getUsersPostsV2({
    users,
    skip,
  }: {
    users: number[];
    skip: number;
  }): Promise<PostSummary[]> {
    try {
      const targetUsers: User[] = [];
      users.forEach((uid) =>
        targetUsers.push(this.userRepo.getUser(BigInt(uid))),
      );
      return this.postRepo.getUsersPostSumms(targetUsers, skip);
    } catch (err) {
      throw err;
    }
  }

  getAllPostsV2(skip: number): Promise<PostSummary[]> {
    try {
      return this.postRepo.getAllPostSumms(skip);
    } catch (err) {
      throw err;
    }
  }

  /**
   * Gets post by slug.
   * @param slug
   * @returns
   */
  getPostBySlugV2(slug: string): Promise<Post> {
    try {
      return this.postRepo.getPostBySlug(slug);
    } catch (err) {
      throw err;
    }
  }

  getPostByIdV2(id: number): Promise<Post> {
    return this.postRepo.getPostByID(BigInt(id));
  }

  /**
   * Get users' posts without comments.
   * @param users
   * @param skip
   * @returns
   */
  async getUsersPostsNoComments({
    users,
    skip,
  }: {
    users: number[];
    skip: number;
  }): Promise<PrismaPost[]> {
    try {
      const userPosts = await this.prismaService.post.findMany({
        where: {
          author: {
            in: users,
          },
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
          mentions: {
            select: {
              mentionee: true,
            },
          },
        },
        orderBy: {
          createdAt: 'desc',
        },
        take: 10,
        skip: skip,
      });

      return userPosts;
    } catch (error) {
      throw new Error('Failed to get user posts');
    }
  }

  /**
   * Gets post by slug.
   * @param slug
   * @returns
   */
  getPostBySlug(slug: string): Promise<PrismaPost> {
    return this.prismaService.post.findUnique({
      where: {
        slug: encodeURIComponent(slug),
      },
      include: {
        images: {
          select: {
            order: true,
            id: true,
            src: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
    });
  }

  /**
   * Gets post by id.
   * @param id
   * @returns
   */
  getPostById(id: number): Promise<PrismaPost> {
    return this.prismaService.post.findUnique({
      where: {
        id: id,
      },
      include: {
        images: {
          select: {
            order: true,
            id: true,
            src: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
    });
  }

  /**
   * Gets all posts.
   * @param skip
   * @returns
   */
  getAllPosts(skip: number): Promise<PrismaPost[]> {
    return this.prismaService.post.findMany({
      include: {
        images: {
          select: {
            order: true,
            id: true,
            src: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        comments: {
          select: {
            id: true,
            postId: true,
            author: true,
            createdAt: true,
            updatedAt: true,
            content: true,
            likenum: true,
            mentions: {
              select: {
                mentionee: true,
              },
            },
            replies: {
              select: {
                id: true,
                parentId: true,
                author: true,
                createdAt: true,
                updatedAt: true,
                content: true,
                likenum: true,
                mentions: {
                  select: {
                    mentionee: true,
                  },
                },
              },
            },
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
      orderBy: {
        createdAt: 'desc',
      },
      take: 10,
      skip: skip,
    });
  }

  /**
   * Gets users' ids and returns posts from the users
   * @param param0
   * @returns
   */
  async getUsersPosts({
    users,
    skip,
  }: {
    users: number[];
    skip: number;
  }): Promise<PrismaPost[]> {
    try {
      const userPosts = await this.prismaService.post.findMany({
        where: {
          author: {
            in: users,
          },
        },
        include: {
          images: {
            select: {
              id: true,
              order: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
          comments: {
            select: {
              id: true,
              postId: true,
              author: true,
              createdAt: true,
              updatedAt: true,
              content: true,
              likenum: true,
              mentions: {
                select: {
                  mentionee: true,
                },
              },
              replies: {
                select: {
                  id: true,
                  parentId: true,
                  author: true,
                  createdAt: true,
                  updatedAt: true,
                  content: true,
                  likenum: true,
                  mentions: {
                    select: {
                      mentionee: true,
                    },
                  },
                },
              },
            },
          },
          mentions: {
            select: {
              mentionee: true,
            },
          },
        },
        orderBy: {
          createdAt: 'desc',
        },
        take: 10,
        skip: skip,
      });

      return userPosts;
    } catch (error) {
      throw new Error('Failed to get user posts');
    }
  }

  /**
   * Get posts by user id
   * @param uid
   * @param skip
   * @returns
   */
  getUserPosts(uid: number, skip: number): Promise<PrismaPost[]> {
    return this.prismaService.post.findMany({
      include: {
        images: {
          select: {
            id: true,
            order: true,
            src: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        comments: {
          select: {
            id: true,
            postId: true,
            author: true,
            createdAt: true,
            updatedAt: true,
            content: true,
            likenum: true,
            mentions: {
              select: {
                mentionee: true,
              },
            },
            replies: {
              select: {
                id: true,
                parentId: true,
                author: true,
                createdAt: true,
                updatedAt: true,
                content: true,
                likenum: true,
                mentions: {
                  select: {
                    mentionee: true,
                  },
                },
              },
            },
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
      where: { author: uid },
      orderBy: {
        createdAt: 'desc',
      },
      take: 10,
      skip: skip,
    });
  }
}
