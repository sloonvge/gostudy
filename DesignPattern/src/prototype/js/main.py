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

if __name__ == "__main__":
    m1 = Resume("Tom")
    m1.SetPersonalInfo("Man", 25)
    m1.SetWorkExperience("xxx-xxx", "Com1")

    m2 = Resume("Tom")
    m2.SetPersonalInfo("Man", 25)
    m2.SetWorkExperience("xxx-xxx", "Com1")

    m3 = Resume("Tom")
    m3.SetPersonalInfo("Man", 25)
    m3.SetWorkExperience("xxx-xxx", "Com1")

    m1.Show()
    m2.Show()
    m3.Show()

    m4 = Resume("Jack")
    m4.SetPersonalInfo("Man", 39)
    m5 = m4
    m4.SetPersonalInfo("Man", 9)
    m5.Show()