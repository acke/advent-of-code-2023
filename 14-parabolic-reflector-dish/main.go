package main

import (
    "fmt"
)

func main() {
    var counter int
    lines := GetInput()
    PrintLines(lines)

    n:=0
    for n < 1000 {

        RotateBoard(lines, "north")
        RotateBoard(lines, "west")
        RotateBoard(lines, "south")
        RotateBoard(lines, "east")

        n++
    }

    PrintLines(lines)
    counter += CalculateStones(lines)
    fmt.Println(counter)
}

func RotateBoard(lines []string, direction string){
    rows := (len(lines)-1)
    collums := (len(lines[0])-1)

    for i, line := range lines {
        for j, _ := range line{
                switch direction{
                case "north":
                    MoveStoneNorth(lines, i, j, rows)
                case "west":
                    MoveStoneWest(lines, i, j, collums)
                case "south":
                    MoveStoneSouth(lines, (rows-i), j, rows)
                case "east":
                    MoveStoneEast(lines, i, j, collums)
                }
        }
    }
}

func MoveStoneEast (lines []string, i, j, collums int){
    moveMore, checkNextStone := true, true
    targetStone := j+1

    if j >= collums {return}

    for (moveMore){
        if lines[i][j] == 'O' {
            for (checkNextStone){
                if lines[i][targetStone] == '.' {
                    lines[i] = UpdateSlice(lines[i], ".", j)
                    lines[i] = UpdateSlice(lines[i], "O", targetStone)
                    checkNextStone = false
                } else if lines[i][targetStone] == '#' || targetStone >= collums {
                    moveMore = false
                    checkNextStone = false
                } else{
                   targetStone++
                }
            }
        } else {
            moveMore = false
        }
    }
}

func MoveStoneWest (lines []string, i, j, collums int){
    moveMore := true

    if j >= collums {return}

    for (moveMore){
        if lines[i][j+1] == 'O' && lines[i][j] == '.' {
            lines[i] = UpdateSlice(lines[i], "O", j)
            lines[i] = UpdateSlice(lines[i], ".", j+1)

            if j > 0 {
                j -= 1
            }

        }else{
            moveMore = false
        }
    }
}

func MoveStoneNorth (lines []string, i, j, rows int){
    moveMore := true
    if i >= rows {return}

    for (moveMore){
        if lines[i+1][j] == 'O' && lines[i][j] == '.' && i < rows {
            lines[i] = UpdateSlice(lines[i], "O", j)
            lines[i+1] = UpdateSlice(lines[i+1], ".", j)

            if i > 0 {
                i -= 1
            }

        }else{
            moveMore = false
        }
    }
}

func MoveStoneSouth (lines []string, i, j, rows int){
    moveMore := true

    if i >= rows {return}

    for (moveMore){
        if lines[i][j] == 'O' && lines [i+1][j] == '.'{
            lines[i] = UpdateSlice(lines[i], ".", j)
            lines[i+1] = UpdateSlice(lines[i+1], "O", j)

            if i < (rows-1){
                i += 1
            }

        }else{
            moveMore = false
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

func PrintLines (lines []string){

    for _, line := range lines {
        fmt.Println(line)
    }
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
