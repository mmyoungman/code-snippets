import copy

def get_generation(cells, generations):
    oldGrid = copy.deepcopy(cells)
    grid = copy.deepcopy(oldGrid)
    while generations > 0:
        # Increase size of oldGrid
        width = len(grid[0])
        oldGrid.append([0 for i in range(width)])
        oldGrid.insert(0, [0 for i in range(width)])
        for row in oldGrid:
            row.append(0)
            row.insert(0, 0)

        grid = copy.deepcopy(oldGrid)

        # Update grid
        for y, row in enumerate(oldGrid):
            for x, cell in enumerate(row):
                count = get_neighbours(x, y, oldGrid)
                if grid[y][x] == 1:
                    if count < 2:
                        grid[y][x] = 0
                    elif count > 3:
                        grid[y][x] = 0
                else:
                    if count == 3:
                        grid[y][x] = 1

        # Trim grid
        while all( elem == 0 for elem in grid[0] ):
            del grid[0]
        while all( elem == 0 for elem in grid[len(grid)-1] ):
            del grid[len(grid)-1]
        firstElemZero = True
        while firstElemZero == True:
            for row in grid:
                if row[0] == 1:
                   firstElemZero = False
            if firstElemZero == True:
                for row in grid:
                    del row[0]
        lastElemZero = True
        while lastElemZero == True:
            for row in grid:
                if row[len(row)-1] == 1:
                   lastElemZero = False
            if lastElemZero == True:
                for row in grid:
                    del row[len(row)-1]

        oldGrid = grid
        generations -= 1
    return grid
    
def get_neighbours(x, y, cells):
    height = len(cells)
    width = len(cells[0])
    count = 0
    if x - 1 >= 0 and y - 1 >= 0:
        count += cells[y-1][x-1]
    if x - 1 >= 0:
        count += cells[y][x-1]
    if x - 1 >= 0 and y + 1 < height:
        count += cells[y+1][x-1]
    if y + 1 < height:
        count += cells[y+1][x]
    if x + 1 < width and y + 1 < height:
        count += cells[y+1][x+1]
    if x + 1 < width:
        count += cells[y][x+1]
    if x + 1 < width and y - 1 >= 0:
        count += cells[y-1][x+1]
    if y - 1 >= 0:
        count += cells[y-1][x]
    
    return count