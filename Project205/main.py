from Task import Task
from TaskGenerator import TaskGenerator
from TaskRunner import TaskRunner

taskGenerator=TaskGenerator("number384.ini")
taskGenerator.parseFile()
taskList=taskGenerator.getTaskList()

runner=TaskRunner()
runner.setTask(taskList)
try:
    runner.start()
except Exception as e:
    print(e)
