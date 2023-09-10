import { PostIds } from 'src/domain/aggregates/post/post.vo';
import { Timestamps } from 'src/domain/vo';

export type PostSumm = {
  id: bigint;
  author: bigint;
  createdAt: Date;
  updatedAt: Date;
  imageSrcs: string[];
  slug: string;
};

export class PostSummary {
  private constructor(
    public readonly author: bigint,
    public readonly id: bigint,
    public readonly slug: string,
    public readonly imageSrcs: string[],
    public readonly createdAt: Date,
    public readonly updatedAt: Date,
  ) {}

  static create(
    author: bigint,
    posdIds: PostIds,
    imageSrcs: string[],
    timestamps: Timestamps,
  ): PostSummary {
    return new PostSummary(
      author,
      posdIds.id,
      posdIds.slug,
      [...imageSrcs],
      timestamps.createdAt,
      timestamps.updatedAt,
    );
  }
}
