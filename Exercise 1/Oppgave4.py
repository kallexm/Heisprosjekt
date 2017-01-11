from threading import Thread


i = 0

def thread1():
    for j in range(1000000):
        global i 
        i += 1

def thread2():
    for j in range(1000000):
        global i 
        i -= 1

def main():
    threadOne = Thread(target = thread1, args = (), )
    threadTwo = Thread(target = thread2, args = (), )
    threadOne.start()
    threadTwo.start()
    threadOne.join()
    threadTwo.join()
    print("The value is: ", i)

main()

