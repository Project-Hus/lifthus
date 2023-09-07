export class CreatePostRequestDto {
  constructor(public author: string, public content: string) {}
}

export class CreatePostServiceDto {
  author: bigint;
  content: string;
  imageSrcs: string[];
  constructor(
    pj: CreatePostRequestDto,
    imageFiles: Array<Express.Multer.File>,
  ) {
    this.author = BigInt(pj.author);
    this.content = pj.content;
    this.imageSrcs = CreatePostServiceDto.getImageSrcs(imageFiles);
  }
  private static getImageSrcs(
    imageFiles: Array<Express.Multer.File>,
  ): string[] {
    return imageFiles.map((imageFile) => imageFile.location);
  }
}

export class UpdatePostRequestDto {
  constructor(
    public id: string,
    public author: string,
    public content: string,
  ) {}
}

export class UpdatePostServiceDto {
  id: bigint;
  author: bigint;
  content: string;
  constructor(up: UpdatePostRequestDto) {
    this.id = BigInt(up.id);
    this.author = BigInt(up.author);
    this.content = up.content;
  }
}
