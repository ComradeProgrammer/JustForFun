import re
from Task import Task

class TaskGenerator:
    def __init__(self,configFileName):
        self.configFileName=configFileName
        self.taskMap={}
    def getTaskMap(self):
        return self.taskMap
    def getTaskList(self):
        res=[]
        for i in self.taskMap.values():
            res.extend(i)
        return res
    def parseFile(self):
        config=""
        ini=open(self.configFileName,"r")
        factorsListMap={}
        while True:
            line = ini.readline()
            if not line:
                break
            #for each config line like [Config random_number384_multi_aodv_2]
            result=re.match(r'\[Config\s*(\S*)\s*\]',line)
            if result!=None:
                config=result.group(1)
                if config=="General":
                    continue
                self.taskMap[config]=[]
                factorsListMap[config]=[]
                continue
            #for something like ${21..40 step 1} or ${1,2,3}
            result2=re.search(r'\$\{(.*)\}',line)
            line=line.strip()
            if result2!=None and line[0]!="#":
                content=result2.group(1)
                if "step" in content:
                    #something like ${21..40 step 1}
                    result3=re.match(r'(\d*)\.\.(\d*)\s*step\s*(\d*)',content)
                    if result3==None:
                        print("Unrecognized pattern"+content)
                        continue
                   
                    start=float(result3.group(1))
                    end=float(result3.group(2))
                    step=float(result3.group(3))
                    num=round((end - start) / step + 1)
                    factorsListMap[config].append(num)
                    
                else:
                    #something like ${1,2,3}
                    args = re.split(r',', content)
                    factorsListMap[config].append(len(args))
        #print(factorsListMap)
        for i in factorsListMap.keys():
            res=1
            for j in factorsListMap[i]:
                res*=j
            for j in range(0,res):
                task=Task(i,j,self.configFileName)
                self.taskMap[i].append(task)

        
                











        




