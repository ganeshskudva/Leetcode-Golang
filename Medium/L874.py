class Solution:
    def robotSim(self, commands, obstacles):
        obstacle_set = set()
        for obs in obstacles:
            obstacle_set.add(f"{obs[0]} {obs[1]}")

        directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]  # north, east, south, west
        d = 0  # direction index
        x, y = 0, 0  # initial position
        result = 0

        for command in commands:
            if command == -1:  # turn right
                d = (d + 1) % 4
            elif command == -2:  # turn left
                d = (d - 1) % 4
            else:
                for _ in range(command):
                    next_x = x + directions[d][0]
                    next_y = y + directions[d][1]
                    if f"{next_x} {next_y}" not in obstacle_set:
                        x, y = next_x, next_y
                        result = max(result, x * x + y * y)

        return result
