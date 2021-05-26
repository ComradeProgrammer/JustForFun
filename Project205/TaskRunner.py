from Task import Task
from multiprocessing import cpu_count
import threading
from time import sleep
class TaskRunner:
    def __init__(self):
        self.taskList=[]
        self.taskNum=0
        self.coreNum=cpu_count()
        self.threadList=[]
        self.threadNum=0
        self.lock=threading.Lock()
    def setTask(self,taskList):
        self.taskList=taskList
        self.taskNum=len(taskList)
    def setCoreNum(self,coreNum):
        self.coreNum=coreNum

    #每个工作线程所执行的函数
    def run(self):
        self.lock.acquire()
        self.threadNum+=1
        self.lock.release()
        while True:
            self.lock.acquire()
            if len(self.taskList)==0:
                self.lock.release()
                break
            task=self.taskList.pop(0)
            self.lock.release()
            task.run()
        self.lock.acquire()
        self.threadNum-=1
        self.lock.release()
        return

    def daemon(self):
        postfix=["|","/","-","\\"]
        index=0
        while True:
            self.lock.acquire()
            unfinished=len(self.taskList)
            finished=self.taskNum-unfinished
            total=self.taskNum
            running=self.threadNum
            self.lock.release()
           
            percentage=finished*100//total
            s="#"*(percentage//2)
            print("[%-50s %-2d%%%s](%d/%d),%dtasks running\r"
            %(s,percentage,postfix[index],finished,total,running),
            end="")
            index=(index+1)%4
            if unfinished==0:
                print()
                return
            sleep(1)

    def start(self):
        print("using %d processes"%(self.coreNum))
        daemonThread=threading.Thread(target=self.daemon,args=())
        daemonThread.daemon=True
        daemonThread.start()
        for i in range(self.coreNum):
            t=threading.Thread(target=self.run,args=())
            self.threadList.append(t)
            t.daemon=True
            t.start()
        for i in self.threadList:
            i.join()
        daemonThread.join()
