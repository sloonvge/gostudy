import sys

numbers = [int(i) for i in sys.stdin.readline().split(",")]

# def maxLost(array, s):
#     res = 0
#     n = len(array)
#     for i in range(s, n):
#         for j in range(i+1,n):
#             if array[j] < array[i]:
#                 lost = maxLost(array, j + 1) + array[i] - array[j]
#                 res = max(res, lost)
#     return res
def maxLost(nums):
    res = 0
    n = len(nums)
    for i in range(1, n):
        if nums[i] < nums[i - 1]:
            res += nums[i - 1] - nums[i]

    return res

print(maxLost(numbers))