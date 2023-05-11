import { Test, TestingModule } from '@nestjs/testing';
import { QueryController } from 'src/query/query.controller';
import { PostMutationService } from './post.mutation.service';
import { MutationController } from './mutation.controller';

describe('AppController', () => {
  let queryController: QueryController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [MutationController],
      providers: [PostMutationService],
    }).compile();

    queryController = app.get<QueryController>(QueryController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(queryController.getHello()).toBe('Hello World!');
    });
  });
});
