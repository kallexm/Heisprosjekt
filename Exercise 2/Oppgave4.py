import threading
from threading import Thread

i = 0
lock = threading.Lock()

def thread1():
   for j in range(1000000):
        lock.acquire() 
        global i 
        i += 1
        lock.release()

def thread2():
    for j in range(999999):
        lock.acquire()
        global i 
        i -= 1
        lock.release()

def main():
    threadOne = Thread(target = thread1, args = (), )
    threadTwo = Thread(target = thread2, args = (), )
    threadOne.start()
    threadTwo.start()
    threadOne.join()
    threadTwo.join()
    print("The value is: ", i)

main()

