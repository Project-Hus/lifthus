export class PostSummary {
  constructor(
    public readonly author: bigint,
    public readonly id: bigint,
    public readonly slug: string,
    public readonly imageSrcs: string[],
    public readonly createdAt: Date,
    public readonly updatedAt: Date,
  ) {}
}
