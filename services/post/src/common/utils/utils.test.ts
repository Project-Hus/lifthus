import { Test, TestingModule } from '@nestjs/testing';
import { getCookie, generateSlug } from './utils';

import testdata from './testdata/util.testdata.json';

describe('utils', () => {
  describe('generateSlug', () => {
    for (const k in testdata) {
      it(`should return "${testdata[k]}"`, () => {
        expect(generateSlug(k)).toBe(testdata[k]);
      });
    }
  });
});
