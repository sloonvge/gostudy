# 
import copy

class WorkExperience():
    def WorkDate(self, workDate):
        self.workDate = workDate
    def Company(self, company):
        self.company = company
    def Clone(self):
        return copy.copy(self)

class Resume():
    
    def __init__(self, name, work=WorkExperience()):
        self.name = name
        self.work = work.Clone()
    
    def SetPersonalInfo(self,sex, age):
        self.sex = sex
        self.age = age
    
    def SetWorkExperience(self, workDate, company):
        self.work.WorkDate(workDate)
        self.work.Company(company)

    def Show(self):
        print("{} {} {}".format(self.name, self.sex, self.age))
        print("工作经历: {} {}".format(self.work.workDate, self.work.company))
    
    def Clone(self):
        _re = Resume(self.name, self.work)
        _re.sex = self.sex
        _re.age = self.age
        return copy.copy(_re)

if __name__ == "__main__":
    m1 = Resume("Tom")
    m1.SetPersonalInfo("Man", 25)
    m1.SetWorkExperience("xxx-xxx", "Com1")

    m2 = m1.Clone()
    m2.SetPersonalInfo("Man", 25)
    m2.SetWorkExperience("xxx-xxx", "Com2")

    m1.Show()
    m2.Show()