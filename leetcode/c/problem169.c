int majorityElement(int* nums, int numsSize) {
    int m = 0;    
    int res = 0;
    for (int i = 0; i < numsSize; i++) {
        if (m == 0)  res = nums[i];
        m += nums[i] == res ? 1 : -1;
    }
    return res;
}
