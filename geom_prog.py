import random


print("Welcome to the Brain Games!\nMay I have your name?")
user_name = str(input())
print(f"Hello, {user_name}!")
print("What number is missing in the progression?")

while True:
    length = random.randint(5, 10)
    numbers = [0]*length
    q = random.randint(2, 7)
    print("Question:", end=" ")
    for i in range(length):
        if i == 0:
            numbers[i] = random.randint(1, 7)
            continue
        numbers[i] = numbers[i-1] * q

    hiden = random.randint(1, length-2)

    for i in range(length):
        if i == hiden:
            print("..", end=" ")
            continue
        print(numbers[i], end=" ")
    print()
    print("Your answer:")
    user_answer = int(input())
    if user_answer == numbers[hiden]:
        print("Correct!")
    else:
        print(user_answer, "is wrong answer ;(. Correct answer was", numbers[hiden], ".")
        break
