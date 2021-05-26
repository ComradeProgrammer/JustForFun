import subprocess
import os
class Task:
    def __init__(self,taskName,taskNumber,taskConfigFile):
        self.taskName=taskName
        self.taskNumber=taskNumber
        self.taskConfigFile=taskConfigFile

    def getShellCommand(self):
        command=["opp_run","-r",str(self.taskNumber),"-m","-u","Cmdenv","-c",
        self.taskName,"-n","../../../src:../..:../../../tutorials:../../../showcases",
        "--image-path=../../../images","-l","../../../src/INET","--cmdenv-redirect-output=false","--record-eventlog=false","--scalar-recording=false","--vector-recording=false",self.taskConfigFile]
        return command

    def run(self):
        c=self.getShellCommand()
        #for test
        #c=["./test",str(self.taskNumber)]
        devNull = open(os.devnull, 'w')
        subprocess.run(c)

    