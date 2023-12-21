package main

import (
    "fmt"
)

func main() {
    var counter int
    lines := GetInput()

    for i, line := range lines {
       for j, item := range line{
            if item == '.' && i < (len(lines)-1){
                nextLine := lines[i+1]
                if nextLine[j] == 'O' {
                    MoveStone(lines, i, j)
                }
            }
        }
    }

    counter += CalculateStones(lines)

    fmt.Println(counter)
}

func MoveStone (lines []string, i, j int) {

    moveUp := true

    for (moveUp){
        lines[i] = UpdateSlice(lines[i], "O", j)
        lines[i+1] = UpdateSlice(lines[i+1], ".", j)

        if i > 0 {
            i -= 1
        }

        if (lines[i][j] != '.'){
            moveUp = false
        }
    }
}

func UpdateSlice (input string, replacement string, location int) string {
    return input[:location] + string(replacement) + input[location+1:]
}

func CalculateStones(lines []string) (counter int){
    for i, line := range lines {
        for _, item := range line{
            if item == 'O'{
                counter += len(lines) - i
            }
        }
    }
    return
}



func GetInput () (lines []string){
    return [] string{
".O..OO#......O...O...#.O.#....O..#....O..O#..O##O#.....OO##.O.....#.....O.....O.####.##O##.O.O.#.O.O",
".#O#OO.....#O.....#....OO####.OO..O....#.#....#......#O#.....O.O.......O...OO#...O.....#.#.O#..##.#.",
"#...O....OO.#...O.O..O..#...O....#...##......O..#...O.#O.O.....O.......O.O.#..##...#O..#..#.......#O",
"#.O.O.....#O##O.......O..O#O..........O##...#O.#.....O..O.##.#.#...#..#.#O..O.....O........O....#O..",
".....O....#O.........O....O.OOO.#.#...O###.#.#O..O...#.#.#....O...O.O..O.#O#.O.#.O....O#.#...OO#...#",
"O.O.O#.#....##O..#..O.OO.#....O#O.O..OOO..O#O......#...#.#..O......#......O.OOO##.......O#.O.#......",
"O.......OO..O..O.O..OO...#OO.O#....O..##O.O.OO....O.O....#......O..O.OO.O....OO...OOO.OO.OO.#.O...OO",
"OOO..O.#.O.......O....O...O.#O.......O.O.......##.#..OOO#..#...#...O............O.O...O#.O...OO.#O##",
".......##.O.O..#.#O..#O..O.##.#...........##.##O.O#.OO#..O......O..............#...O..O#...OO..O#..#",
".....#.....#....O....O...#.....O#.O.#..O#O.O.O..#.#...O...O..OO...OO..........OO#...##...O..#.O..#O.",
"OO..#O.....O#..#..OO....O.OO..#....#.O.#.#.#....O.....O#O...............#.......O..O#O.....O#......#",
".OO.....O#O...#...##.O.#.O....O...O.#.........O.#.......O.O..#O..#...O.OO..O#O.O#.........O.....O...",
"..O.......O#O.#O..O.#.....#.OO##OO...##...O.O..#..OO.....#.O.#OOO.O#O....OOOOO.O.O....OO#..#...OOOO.",
"...O..O....#.#..####..OO...O...#.O.O#O#O..#.....#.O#........OO.O##...#...O.O..O#O.O...#..O...#....O.",
"#..O...O#...O..O....#.....#OO..O..O...O#.#.OO.O.O....#.O..O....#O....O..OO.O#..O..#.O.O..OO.....#O..",
".#O.#.#...O#....O..##..OO..#.....#..O...O#O#O...##OO..#.O.O......O.#O.......OO..O#.O#O...O..........",
".O.O......#O.#OO.O.....#...........O....OO#O.O.....O...#..O...O...#......#.OO...#.#.#..O.#O......OO#",
"..#O..#.#.O...OOO.O.O..##..OO......O.O......O.##.....O......#OO...O.O.#.#..##.OOOOO.#.#.....OO..O...",
"....O..#.O.O...#..O#..#...............O...O..................#.O.#...#OOOOO#.......#.#...#.O.#..#.#O",
"..#.#O....#..#O.O#.....OO#.........O...O.O...#..O#.#..O.#..O#....#..O..O....#......O.O..O.#O.O.#..O#",
"...#.#O...#O..O#....O#.#......##OO#...O...O.....#.##..#......#O..#.O.......O....O..............#O.OO",
".....O..#..#.......O#.O.O.#.O.#.....O#..#O#OO..##...#.##..........#OOOO..OO#..O...OO#...O.....##..O.",
".##..O..O.O#.O........O.#.O..#O...O.....#............O.O.O.....#.......OO.#.##O........#.OO.##OO..#.",
".#..O#O..###O#.#.OOO........#...#.O....OO...O..O##...#.#..#.O.......OO.#....O.#..#..###O.OOO..#...#.",
"......O...OO.#O.....O..O......OO..###..###...O..##..#.............O..OO#..#.....O..#...#.##..O#O..O.",
"......O.##O..O#..O..O....#..O#..#.....#O##...O..........O....O##...#...O..O....#.O.O..O..OO..O....O.",
"O#.#..O.OO.O...O..#OOOOO#.OOO..#O#O..OO.O#....OO..#.##.##..#.O#............O.....#..#..OO.O....#O.#.",
"..#.O.O.....O.O#..O.#.O#.O.#..O...O..........O#...O....#..O#.O#.OOO...OO..#O..#O.O..#.#O#OO.....O...",
"O..O#O.....#.#O.......#.#OO...O#...O..O.O..O..O.#.O##O.#.....O#.#O....##...##.O..O.O...............O",
"OO...##..O.#......O.....#O...OO.#.O.O....O.#.#.OO..OO.#....OOO...O..#....#.O..#....#O...#......#O.#.",
"..O##O..O.O.O##...O......#..O.O.O###.O..O...#O....O...........##.OO...O.#.###O#.#.O...##.O......#.O#",
"..OO...O.O#..O......OO....#..#...O.O..#.O....#O.#.O.#..OOO#..O.O.O.O##.O....OO.O..O.O###.#O.O...OO.O",
"...........OO.O.#...O..O#..OO.....O.O..O.#.#.OOO#.O..OO#.O....O..O..#.OO#O.#.##O.....#...#..........",
"..O.#..OO.O.O#.....O.O.#.O##O#O.O.O#........#..O#.......O.O..#.##....OO..#.O...O..#.......O.OOOO...O",
"....O.#..OOO#..#.#..#.........#..#.#.O...OO..........O#O...O..OO..#.O..O.....O.....O....#..#.....O..",
"OO#.#O..OO...O..O#....#O...O.##...##..#.#..#.O.##..#..#......O.#.O.O.##..OO........OO...O.OO#....O..",
"O#.#......O.O..O#.#...OO#O......O#OO.OO.....O..O#.......#...O###..#..O#O..#........#O..OOOO#.O#..O#.",
"...O.###.#.........O....OO.O##...##....#.O#O........O........O...#.#........#O......O........O##O..O",
"....O.#.O#.#..O.......##.#.#..O.O.O.......#O..#.#O.O.#....O......O...O....#.....O...#O.....O.#.....O",
"#.....O.#...O.#..O##..O...##..O...#.....O...#....O##OO#O#...O.....#O..O..#...O.OO.O..O.....#.OO...O.",
"...O....O.#..O....O..#...##.....#....#.................#...O....O.O.OO##...O#......O.#.#.#.........O",
".O.O..O...O..O#.O......#O.#O....#..OO.O.O..............#O....O....O....#..#.O.....#OO...O..#.O..O#.O",
"O.#.#..O.OO..O.....O#O.O.#..O..#..OO...OO.....O.........#O...........##.O.....#...O#O.#....OO.O.O.O#",
".O.O.#OO#OO.O.OOOO.#...OO.OO.OO#O.O..O#...O....##.#..#....O.....#..#......O..O.....OO....O.O.O#.OO#.",
"..OO...#..#.O.OO..#......O#.#.O...O.OO..#O..O..#..O.##..O...##.O.....#.#.#.#..###O.......#..##..O.O#",
"...O...O.O...O#...#....O.OOO...OOOO.OO#....O#..O#.##...##O.O...#O.O......#....#O.O..O...#O..OO...OO.",
".##...OO.##..O.........#.......#.#.....O.O.#....O.#....OO..##O..O#.....O#O....OO..O....O....O...O...",
"O........O##...#...#.O....O...#OOOO##..O...OO.....##......#..O.O#.O..O#O..#O...#.......#......O.OO..",
".OOO..#O.O#O.O..O...O..OO.O.....#...#.O#OOO..#.OOO.#...#....OO.OO.O.#...#.O....O.....#..O.O#.O...OO.",
"O.......#...OO..O.#..#O..#...OO.##O..#O.O...#.O..#..O#.....O..#...#OO...##..#............O.....O#...",
"#OO...OO.O.OO...#OO.#......O...#..#....O.#....#O.#O.#.O.##....O..O..O.#O.O.....#..........#....OO.O.",
"..O......O.#.O.OO...O.#...O.O#.##......OO.....#O.#...O.OO..#O..#...#..#.O......O#.#...O....O###O.O..",
".#...OO.#.#.O...#...O...O..#..#..#..#...OO...#.OO......O..O#..O.O.O.O...#.O...#.......OO.....OO...O.",
"O.O.O...O.O..O..O##......#...#...O..O.O..#.O.#...O...O.#....#...O.O.#.##O#..#..#....O#O.#O##.OO#.O..",
".#......O.....O.....#....#..O.....#.....#O..O#.........#OO#...O##.O.....O.#..OO.........O.....#..O..",
"..O....##.O..O..O..#.O#...O......OO.OO....##..#..#...#.##O....O#..O.O....#.O##.....#O#...OO...#O.O..",
"OO...#.O...O#..O..O.....O...#..O..#.#..OO#O.###.#...OO..#.....O....O#.O..#OO..##..#O..O.OO.....##.O.",
".#.........#..#O......O..OO...#...O#..OO.O#....O#........OO.#......O..#..O##O.O.....OOO.OO.........O",
"#..OOO...O..OO#..O......#O......#..O.OOO#..OO...O.OOO##.O...O.O.O..#...O...#...#..O.OOOO..O..O......",
"...#..#O....O.O#O.OOOO..O.O...#.O.#..OO....OOOO....O.O#...............O#......O.........#.#..O..O..#",
"..#.....OO..#.#.O#O..O.O#O...#.....O.#....#.O#O#..#.#.....##..O..O#.OO...OO....O.OO.O.O#O..###....O.",
"...O..O.......#.O.O......O........O......#...O....O....#O#.#.OO...#.#.O.O...##...#O.O...OO...O......",
"###O.##..O##O...........#..##O.....OO..O..#.O....O................O....#.#O#.#.O...#.....O...O##O##.",
".O.....O...O#O....OO#.O#...#......#...O...O..#.#........#.#.##.........O.##.O.OO#.#.#..#.....O#.....",
"OO...........#O.O#.O...#...OO....#.#....#O..OO...OO....O..#O..O.O....#.O.O....O....O#.##O..#O..O#O#O",
"..O##.....#.O.#...O##..O#.O.O.#O#..#O..O.OO#O..#..#...#..#....O.....#.....OOO.#.O....#.O#...#....O##",
"....O.O.#..O.##.#.O.#.O..O....O..O#.#..#...OO....OO#O#....###.#OO..#.....##.....#.###.O...#..O#O#...",
".#...#...O.#OO...#O..O.##.#.O.......OO..#.O..#...#........O.O.......#..O.##...##........O...O...O.O.",
"...#...#O#.O#.......O..O...O#O...O#.OO.........O#.O...O..O....#O.O...#O.#OO...O..O....#O..O......#..",
"....O.....O.....#...#O.O.#OO..#.OO#O...##..O#...##..........O..O.#..OOO#..O...#..#O.O......#OOO.O..O",
".#.O..O.........OOO........#O.OOO.......O.O..O.#O......#..#O#...O...#O....O.##.O....#.OO...O.OO.#...",
"..#..#O..O#O....#..#.##O..#..O...#....#O........O.O.......O..OOOO#...#.O.O..O..O.#..#O.....#..O.O...",
"..#....O.......OO.....O......O.#O###.OOO.##OO.OO...#..O.OO.OO..#...O.......O.O.O..#O...#.#........O.",
"..#...O..#...O..#.##......OO.O..#..O.OOO.##O.OO.O..##.##.OO#.OO...#...#.....#O.O...O...O..#...#...OO",
"#....##OO..#.......O..O.OO.....OO..#...#.O##..O.O.......#.##.OO.O...#..#.#.O#..O.....##...##O...#.O#",
"O....O#....OO.......O.O......O.O.#O.O#...#.O#....O.O....#.#....#.#.##.O........#.#...O..O...O....O.O",
".OO.O.OO......#....O#.O..#..O.#..#O....O..O#........O...#....O....#..#........O..O........O.O....#..",
"...O..OO.O....O..OO#.O#O..OO#....O#.##...#.....O#..#O.#..#..#.#....O#.##...#...#O...###...###.......",
"..O.O#..#.OO.#..O........O..O#...O..#..O.OO...#O..##O...#.#....O.#..#......O.O..#O...#O.#...#O...OO.",
"..O...O#..O#.O.#O.OOO..#.#......OO#.#..#......#O.......O..O.#O....#.##......#....O#..###.O......O..O",
"...#.O....#O..##...#............O.O......##..O..#.O..#.O.O.O...OO.O#....O#OO.##O..#..O....#.O#..#...",
"#..#.#.#...OO....O.###.....O...O..O#.O...O.#....#O#...O...O..#..O........O..O..#.......O..O.....O.#.",
"#......#...O...O.........#..#O.....#..O..#.O...#O#.O#.#O#.#..O..O...#O.##..#O.....#.........#.#...OO",
"###OO.....O##.##.O.O.O..#.....O..#O.###.O.#..O...##O##...#O....#......#.O##...#.O...O#O#..#...O.#O.O",
"....O....#.OOO.#..O..#.O..O....O....#...O.............#O##OO...OO....#.OO...#OO......O..O..OO#.O...#",
"..O.#..##.....O.#O.#O.#...O#O.#O.O#OO#.O.O.#....##.O..#OOO....##OO.##..OO.O#..O....#.#......#......#",
"O...#..OO.O..O#.........O##..O....#....#OO#....#O....O..O..#..O...##.....#.....OO...#.O.#O...#.....#",
"OO.O.#...#.O#.O.O.....#..#...OO.##.O##.OO.....O...#..O.......OO#.O.......#OO.#O..#.O...........#.##.",
".O...O##...O.O.OO......O#..O.###O...........O...#.O....OOO.#.....OOOO.#......O.....O..O.....#.......",
".O..#O.OOO...#.O.#.O##....##.....#....#.O..#.O......O.....O...O..#O#..#O.#.O......O....O..#O.#......",
".#.O...#...O..#..#.....O..O.......#....O.#...O#.#OO..#.O#....O.O...#....O..#OO..#.......O.##...O..O.",
"#.OO.#...#....O#.....#...O.....#.#OO...#..O..OO.#......#.....#.O#...#.O.......#....O.#..#..O...O#..#",
"#...#O..##.#...#...O...##.#...#.#.#....##.O....O.O##..#O...#.OO...#O.....#O..#..#.O##...OO..O.#O..#.",
"#.#....O.......#...O....#......OO.#O..#O#..O.#....O.....#O.O.O..O.#..O.O......O...##..O..##.#....#..",
"....#..#.O.#...O.O.O.#O..O......O.O.....#...O.#O..#....#.O.O.O......OO....##....#.O.O.O....O..#...O.",
".OO....#.....O....##...O.......#...#.OO.O#....OO#.#OOO#.#O#...O.O#....O#..O..O.#O...##.....OO..O....",
"O........#....#OOO...O..###..O...O.....O#...#.O.O..O.........OO....O..#O...OOO#.#.......O...#..O.#..",
".O#..OO.##.O.....O.....#..#..#OO...#.....#O.....#..O...O..OOO...O.#.#...O..O.....O.#O..OO.O##....OO.",
"O..O#..#...#.O...#O.........O.#.........O..#O......#..O.O.O.#..#O.##.#...#.......#...###O#..O...O...",
"##..O#.OO.#.O#O.O#......#.....O......O...OO....O..O..O....O..O.O.O.........O.#.O.....OO...##....#O.O",
    }
}
