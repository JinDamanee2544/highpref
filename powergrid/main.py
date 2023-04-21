# usage example: time python3 powerplant.py /input/grid-100-180 /output/grid-100-180.o

import sys  # sys.argv
from ortools.linear_solver import pywraplp

# config
inputPath = ""
outputPath = ""

solver = pywraplp.Solver.CreateSolver('SCIP_MIXED_INTEGER_PROGRAMMING')


def IntegerProg():
    print("Reading...")
    cityCount, _, adjList = read(inputPath)

    print("Solving...")
    # cityList is just index of city
    cityList = [solver.BoolVar(str(i)) for i in range(cityCount)]
    for cid in range(cityCount):
        city = solver.IntVar(1, cityCount, 'city'+str(cid))
        constraint = solver.Constraint(0, 0, 'constraint'+str(cid))
        constraint.SetCoefficient(city, -1)
        constraint.SetCoefficient(cityList[cid], 1)
        for adjCity in adjList[cid]:
            constraint.SetCoefficient(cityList[adjCity], 1)
    # minimal number of powerplants to cover all cities
    # min(sum of cities[i])
    obj = solver.Objective()
    obj.SetMinimization()
    for city in cityList:
        obj.SetCoefficient(city, 1)
    solver.Solve()
    minVal = int(obj.Value())
    isPowerPlant = ''.join([str(int(city.solution_value()))
                           for city in cityList])
    print("Writing...")
    write(outputPath, str(minVal), isPowerPlant)

    print("Done")


def write(outputPath, minVal, isPowerPlant):
    with open(outputPath, 'w') as f:
        f.write(minVal+':'+isPowerPlant)


def read(inputPath):
    # read input file
    with open(inputPath, 'r') as f:
        n = int(f.readline())
        m = int(f.readline())
        adjList = [[] for i in range(n)]
        for line in f:
            a, b = line.strip().split()
            a = int(a)
            b = int(b)
            adjList[a].append(b)
            adjList[b].append(a)
    return n, m, adjList


if __name__ == '__main__':
    # check args
    print("Argc \t: ", len(sys.argv))
    print("Argv \t:", sys.argv)
    if len(sys.argv) != 3:
        print("usage: prog <inputPath> <outputPath>")
    else:
        # args[0] is always prog
        inputPath = sys.argv[1]
        outputPath = sys.argv[2]
        IntegerProg()
