import xlwt
import xlsxwriter
import json2xls

file = xlwt.Workbook()

def json2xlsx():
    pass

def write_xlsx(path, fname, value):
    wb = xlsxwriter.Workbook(fname)
    ws = wb.add_worksheet()
    ws.write(0, 0, 1)
    ws.write(0, 1, 2)
    wb.close()

if __name__ == "__main__":
    write_xlsx(".", "demp.xlsx", 0)

table = file.add_sheet('data')

data = {
        "1":["张三",150,120,100],
        "2":["李四",90,99,95],
        "3":["王五",60,66,68]
}

ldata = []
num = [a for a in data]
num.sort()

for x in num:
    t = [int(x)]
    for a in data[x]:
        t.append(a)
    ldata.append(t)

for i, p in enumerate(ldata):
    for j, q in enumerate(p):
        table.write(i, j, q)
file.save('data.xls')