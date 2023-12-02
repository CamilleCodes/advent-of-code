def get_game_record(game: str) -> tuple[int, list[str]]:
    """Splits the game record information to retrieve the game ID number
    and the subsets of cubes that were revealed from the bag.

    Args:
        game (str): a string representation of the game record

    Returns:
        A tuple consisting of the game ID (int) and the subsets of cubes (list[str])

    Example:
    >>> get_game_details('Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green')
    (1, ['3 blue, 4 red', '1 red, 2 green, 6 blue', '2 green'])
    """

    # Note: using strip() to get rid of any new line chars
    game_details: list[str] = game.strip().split(": ")

    game_id: int = int(game_details[0].split(" ")[1])
    subsets_of_cubes: list[str] = game_details[1].split("; ")

    return game_id, subsets_of_cubes


def get_cube_data(subsets_of_cubes: list[str]) -> list[list[str]]:
    """Splits the subsets of cubes for a game into a usable format

    Args:
        subsets_of_cubes (list[str]): the subsets of cubes revealed for each round of the game

    Returns:
        A list of lists containing the cube color and the number of cubes
        of that color that were revealed during the round

    Example:
    >>> get_cube_data(['3 blue, 4 red', '1 red, 2 green, 6 blue', '2 green'])
    [['3', 'blue'], ['4', 'red'], ['1', 'red'], ['2', 'green'], ['6', 'blue'], ['2', 'green']]
    """

    cube_data = []

    for subset in subsets_of_cubes:
        selected_cubes = subset.split(", ")

        for cube in selected_cubes:
            cube_data.append(cube.split(" "))

    return cube_data


def check_cube_subsets(subsets_of_cubes: list[str], cubes: dict[str, int]) -> bool:
    """Check that the cube subsets for the game are valid for part one of the problem set"""

    for cube_data in get_cube_data(subsets_of_cubes):
        if int(cube_data[0]) > cubes[cube_data[1]]:
            # Invalid subset, move to the next game
            return False

    return True


def part_one() -> None:
    """Determine which games would have been possible if the bag had been loaded with only
    12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
    """

    path = "02/input.txt"
    cubes = {"red": 12, "green": 13, "blue": 14}
    sum_of_ids = 0

    with open(path) as file:
        while game := file.readline():
            game_id, subsets_of_cubes = get_game_record(game)

            if check_cube_subsets(subsets_of_cubes, cubes):
                sum_of_ids += game_id

    print(sum_of_ids)


def determine_game_cubes(subsets_of_cubes: list[str]) -> dict[str, int]:
    """Determine the fewest number of cubes of each color for part two of the problem set"""

    cubes = {"red": 0, "green": 0, "blue": 0}

    for cube_data in get_cube_data(subsets_of_cubes):
        if int(cube_data[0]) > cubes[cube_data[1]]:
            cubes[cube_data[1]] = int(cube_data[0])

    return cubes


def calculate_game_power(game_cubes: dict[str, int]) -> int:
    """Calculate the power value for the game based on the determined cubes
    returned from determine_game_cubes()"""

    power_value = 1
    for cube in game_cubes:
        power_value *= game_cubes[cube]

    return power_value


def part_two() -> None:
    """For each game, find the minimum set of cubes that must have been present.
    What is the sum of the power of these sets?"""

    path = "02/input.txt"
    values = []

    with open(path) as file:
        while game := file.readline():
            _, subsets_of_cubes = get_game_record(game)
            game_cubes = determine_game_cubes(subsets_of_cubes)
            power_value = calculate_game_power(game_cubes)
            values.append(power_value)

    print(sum(values))


if __name__ == "__main__":
    part_one()
    part_two()
