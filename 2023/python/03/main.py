symbols = {
    "!",
    "@",
    "#",
    "$",
    "%",
    "_",
    "+",
    "^",
    "&",
    "*",
    "(",
    ")",
    "=",
    "-",
    ",",
    "[",
    "]",
    "{",
    "}",
    "/",
}


def is_a_part(grid: list[str], r: int, c: int) -> bool:
    row_length = len(grid[0])

    # Check right
    if (c + 1) < row_length:
        if grid[r][c + 1] in symbols:
            return True

    # Check left
    if (c - 1) >= 0:
        if grid[r][c - 1] in symbols:
            return True

    # Check above
    if (r - 1) >= 0:
        if grid[r - 1][c] in symbols:
            return True

        # Check diagnal up left
        if (c - 1) >= 0:
            if grid[r - 1][c - 1] in symbols:
                return True

        # Check diagnal up right
        if (c + 1) < row_length:
            if grid[r - 1][c + 1] in symbols:
                return True

    # Check below
    if (r + 1) < len(grid):
        if grid[r + 1][c] in symbols:
            return True

        # Check diagnal down right
        if (c + 1) < row_length:
            if grid[r + 1][c + 1] in symbols:
                return True

        # Check diagnal down left
        if (c - 1) >= 0:
            if grid[r + 1][c - 1] in symbols:
                return True

    return False


def get_star_location(grid: list[str], r: int, c: int) -> tuple[int, int]:
    row_length = len(grid[0])

    # Check right
    if (c + 1) < row_length:
        if grid[r][c + 1] == "*":
            return (r, c + 1)

    # Check left
    if (c - 1) >= 0:
        if grid[r][c - 1] == "*":
            return (r, c - 1)

    # Check above
    if (r - 1) >= 0:
        if grid[r - 1][c] == "*":
            return (r - 1, c)

        # Check diagnal up left
        if (c - 1) >= 0:
            if grid[r - 1][c - 1] == "*":
                return (r - 1, c - 1)

        # Check diagnal up right
        if (c + 1) < row_length:
            if grid[r - 1][c + 1] == "*":
                return (r - 1, c + 1)

    # Check below
    if (r + 1) < len(grid):
        if grid[r + 1][c] == "*":
            return (r + 1, c)

        # Check diagnal down right
        if (c + 1) < row_length:
            if grid[r + 1][c + 1] == "*":
                return (r + 1, c + 1)

        # Check diagnal down left
        if (c - 1) >= 0:
            if grid[r + 1][c - 1] == "*":
                return (r + 1, c - 1)

    return None


def main():
    path = "03/input.txt"
    # path = "03/test.txt"
    grid: list[str] = []
    total: int = 0
    starred = {}

    with open(path) as file:
        for line in file:
            grid.append(line.strip())

    num: str = ""
    part: bool = False
    star_location = None
    for r, row in enumerate(grid):
        for c, item in enumerate(row):
            if item.isdecimal():
                num += item

                if is_a_part(grid, r, c):
                    part = True
                    star_location = get_star_location(grid, r, c)

            elif num:
                if part:
                    total += int(num)

                if star_location:
                    if star_location in starred:
                        starred[star_location].append(num)
                    else:
                        starred[star_location] = [num]

                    star_location = None

                num = ""
                part = False

            else:
                num = ""
                part = False

    print(total)

    gear_ratios = []
    for star_location, part_nums in starred.items():
        # print(f"{star_location}: {part_nums}")
        if len(part_nums) > 1:
            gear_ratio = 1
            for num in part_nums:
                gear_ratio *= int(num)

            gear_ratios.append(gear_ratio)

    print(sum(gear_ratios))


if __name__ == "__main__":
    main()

# make a grid of the input
# iterate over each row if the grid and find the numbers
# if a number, check the surrounding area for a symbol
# if a symbol is next to it, use the number as a part number
# otherwise, ignore it and move to the next number

# ###
# #1#
# ###
