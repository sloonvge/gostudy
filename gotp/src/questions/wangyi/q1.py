
num1 = 123
num2 = 456
def multiply(num1, num2):
    pass
def helper(numbers, b):
    result = [i for in range(len(numbers) + 1)]
    for n in numbers:
        c = n * b
        if c < 10:
            result.append(c)
        else:
            result.append(c % 10)
            result.append()
multiply(num1, num2)