from collections import deque

SCREEN_WIDTH = 50
SCREEN_HEIGHT = 6

def create_grid():
    return [[0] * SCREEN_WIDTH for row in xrange(SCREEN_HEIGHT)]

def rotate(l, n):
    return l[-n:] + l[:-n]

def rect(grid, x, y):
    for j in xrange(y):
        for i in xrange(x):
            grid[j][i] = 1
    return grid

if __name__ == "__main__":
    with open("input", 'r') as f:
        instructions = [line.rstrip("\n") for line in f]
    grid = create_grid()
    for instr in instructions:
        parsed = instr.split(" ")
        if parsed[0] == "rect":
            dimensions = parsed[1].split("x")
            grid = rect(grid, int(dimensions[0]), int(dimensions[1]))
        elif parsed[0] == "rotate":
            idx = int(parsed[2].split("=")[-1])
            amount = int(parsed[4])
            if parsed[1] == "row":
                grid[idx] = rotate(grid[idx], amount)
            elif parsed[1] == "column":
                l = [grid[row][idx] for row in xrange(SCREEN_HEIGHT)]
                new_list = rotate(l, amount)
                for row in xrange(SCREEN_HEIGHT):
                    grid[row][idx] = new_list[row]
            else:
                raise Exception("Unknown rotation direction")
        else:
            raise Exception("Unknown instruction")
    print grid
    #print sum([ sum( x > 0 for x in grid[j] ) for j in xrange(SCREEN_HEIGHT) ])
