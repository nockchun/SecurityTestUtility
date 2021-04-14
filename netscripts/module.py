def sum(a, b):
    return a + b

def subtraction(a, b):
    if a > b:
        return a - b
    else:
        return "error"

def multiply(a, b):
    i = 1
    result = 0
    while i < b:
        result = result + a
        i += 1
    
    return result

def divide(a, b):
    if b == 0:
        return "not devided by 0"
    else:
        return float(a) / b