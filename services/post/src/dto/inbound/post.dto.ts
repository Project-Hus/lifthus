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

export class UpdatePostDto {
  id: bigint;
  content: string;
  constructor(id: string, content: string) {
    this.id = BigInt(id);
    this.content = content;
  }
}
