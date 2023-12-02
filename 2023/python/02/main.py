# Determine which games would have been possible if the bag had been loaded with only
# 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?


def check_hands(hands: list[str], cubes: dict[str, int]) -> bool:
    for hand in hands:
        selected_cubes = hand.split(", ")

        for cube in selected_cubes:
            cube_data = cube.split(" ")
            # i.e. ['3', 'blue']

            if int(cube_data[0]) > cubes[cube_data[1]]:
                # Impossible, move to the next game
                return False

    return True


def part_one():
    path = "02/input.txt"
    cubes = {"red": 12, "green": 13, "blue": 14}
    total = 0

    with open(path) as file:
        while game := file.readline():
            # Note: get rid of any new line chars
            game_details = game.strip().split(": ")
            hands = game_details[1].split("; ")

            if check_hands(hands, cubes):
                game_num = game_details[0].split(" ")[1]
                total += int(game_num)

    print(total)


def part_two():
    path = "02/input.txt"
    values = []

    with open(path) as file:
        while game := file.readline():
            # Note: get rid of any new line chars
            game_details = game.strip().split(": ")
            hands = game_details[1].split("; ")

            # Determine the fewest cubes needed for the game
            cubes = {"red": 0, "green": 0, "blue": 0}

            for hand in hands:
                selected_cubes = hand.split(", ")

                for cube in selected_cubes:
                    cube_data = cube.split(" ")
                    # i.e. ['3', 'blue']

                    if int(cube_data[0]) > cubes[cube_data[1]]:
                        cubes[cube_data[1]] = int(cube_data[0])

            power_value = 1
            for cube in cubes:
                power_value *= cubes[cube]

            values.append(power_value)

        print(sum(values))


if __name__ == "__main__":
    part_one()
    part_two()

# find the fewest cubes needed for each game
# multiply the values to get the power value for the game
# sum all power values
