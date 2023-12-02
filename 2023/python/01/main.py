# Calibration value
# Combine the first digit and the last digit to form a single two digit number


def part_one():
    path = "01/input.txt"
    values = []

    with open(path) as file:
        while line := file.readline():
            for char in line:
                if char.isdecimal():
                    first = char
                    break

            reversed_line = line[::-1]
            for char in reversed_line:
                if char.isdecimal():
                    last = char
                    break

            values.append(first + last)

    total = 0
    for value in values:
        total += int(value)

    print(total)


def part_two():
    path = "01/input.txt"
    written_nums = {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }
    total = 0

    with open(path) as file:
        while line := file.readline():
            # print(line)
            values = []

            for i, char in enumerate(line):
                # Get the written number values
                three = line[i : i + 3]
                four = line[i : i + 4]
                five = line[i : i + 5]

                if three in written_nums:
                    values.append((i, written_nums[three]))

                if four in written_nums:
                    values.append((i, written_nums[four]))

                if five in written_nums:
                    values.append((i, written_nums[five]))

                # Get the digit number values
                if char.isdecimal():
                    values.append((i, char))

            # Sort the values
            sorted_values = sorted(values)

            first = sorted_values[0][1]
            last = sorted_values[-1][1]

            combined_value = first + last
            total += int(combined_value)

    print(total)


if __name__ == "__main__":
    part_one()
    part_two()


# iterate over written nums
# val and index are stored
# then, iterate over the digits
# append all to values list and then sort by the index
# take the first and last elements, concatenate and turn to int
