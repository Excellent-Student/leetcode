# 两数之和



"""
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

"""

def twoSum(nums, target):
    """
    # 此方法只能计算相邻元素的和，不能计算非相邻元素的和
    if len(nums) <=1 :
        raise Exception("num is empty list or nums length is less than 1")
    else:
        for num_index in range(len(nums)):
            if int(nums[num_index + 1]) > len(nums):
                return []
            if target - int(nums[num_index]) == int(nums[num_index + 1]):
                return [num_index, num_index + 1]
    """
    map = {}
    for index, num in enumerate(nums):
        complement = target - num
        if complement in map:
            return [map[complement], index]
        map[num] = index


if __name__ == '__main__':
    """
    思路 :当列表为空或列表长度小于1时，抛出异常；否则，遍历列表，对于每个元素，计算 target - nums[i]，如果这个值在列表中存在，返回包含这个值的两个元素的索引
    """
    nums = [1, 2, 3, 4, 5, 6, 7]
    targegt = 13
    print(twoSum(nums=nums,target=targegt))