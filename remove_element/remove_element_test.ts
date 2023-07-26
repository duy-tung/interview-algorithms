import { removeElement } from './remove_element';

describe('removeElement', () => {
    it('removes all instances of a value and returns the new length', () => {
        const nums = [3, 2, 2, 3];
        const val = 3;
        const result = removeElement(nums, val);
        expect(result).toBe(2);
        expect(nums.slice(0, result)).toEqual([2, 2]);
    });

    it('returns the same length when the value is not found', () => {
        const nums = [1, 2, 3, 4];
        const val = 5;
        const result = removeElement(nums, val);
        expect(result).toBe(4);
        expect(nums.slice(0, result)).toEqual([1, 2, 3, 4]);
    });

    it('returns 0 when the array is empty', () => {
        const nums: number[] = [];
        const val = 3;
        const result = removeElement(nums, val);
        expect(result).toBe(0);
    });
});
