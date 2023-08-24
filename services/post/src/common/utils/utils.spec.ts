import { Test, TestingModule } from '@nestjs/testing';
import { getCookie, slugify } from './utils';

import testdata from './testdata/util.testdata.json';

describe('utils', () => {
  describe('slugify', () => {
    for (const k in testdata) {
      it(`should return "${testdata[k]}"`, () => {
        expect(slugify(k).slice(0, -16)).toBe(testdata[k]);
      });
    }
  });
});
