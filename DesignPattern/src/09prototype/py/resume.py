# 
import copy

class Resume():
    def __init__(self, name):
        self.name = name
    
    def SetPersonalInfo(self,sex, age):
        self.sex = sex
        self.age = age
    
    def SetWorkExperience(self, timeArea, company):
        self.timeArea = timeArea
        self.company = company

    def Show(self):
        print("{} {} {}".format(self.name, self.sex, self.age))
        print("工作经历: {} {}".format(self.timeArea, self.company))
    
    def Clone(self):
        return copy.copy(self)

if __name__ == "__main__":
    m1 = Resume("Tom")
    m1.SetPersonalInfo("Man", 25)
    m1.SetWorkExperience("xxx-xxx", "Com1")

    m2 = m1.Clone()
    m2.SetPersonalInfo("Man", 30)

    m1.Show()
    m2.Show()