from math import lcm
import random


print("Welcome to the Brain Games!\nMay I have your name?")
user_name = str(input())
print(f"Hello, {user_name}!")
print("Find the smallest common multiple of given numbers.")
while True:
    numbers = []

    for _ in range(3):
        numbers.append(random.randint(1, 100))

    print("Question: ", numbers[0], numbers[1], numbers[2], "\nYour answer:")
    answer = lcm(numbers[0], numbers[1], numbers[2])
    user_answer = int(input())
    if user_answer == answer:
        print("Correct!")
    else:
        print(user_answer, "is wrong answer ;(. Correct answer was", answer, ".")
        print("Let's try again, ",user_name)
        break
