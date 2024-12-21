package printer

import (
	"fmt"
	"strings"
)

func printAny(mode string, msg []interface{}) {
	if len(msg) == 0 {
		return
	}

	switch msg[0].(type) {
	case string:
		allString := true

		for _, e := range msg {
			if fmt.Sprintf("%T", e) == "string" {
				continue
			} else {
				allString = false
				break
			}
		}

		if allString {
			stringArray := make([]string, 0)
			for _, item := range msg {
				stringArray = append(stringArray, item.(string))
			}
			fmt.Println(mode, strings.Join(stringArray, " "), "\u001b[0m")
		} else {
			i := fmt.Sprintf("%v", msg)
			fmt.Println(mode, i[1:len(i)-1], "\u001b[0m")
		}
	default:
		i := fmt.Sprintf("%v", msg)
		fmt.Println(mode, i[1:len(i)-1], "\u001b[0m")
	}
}

// PrintBlack чёрный цвет знаков
func PrintBlack(msg ...interface{}) {
	printAny("\033[0;30m", msg)
}

// PrintRed красный цвет знаков
func PrintRed(msg ...interface{}) {
	printAny("\033[0;31m", msg)
}

// PrintGreen зелёный цвет знаков
func PrintGreen(msg ...interface{}) {
	printAny("\033[0;32m", msg)
}

// PrintYellow желтый цвет знаков
func PrintYellow(msg ...interface{}) {
	printAny("\033[0;33m", msg)
}

// PrintBlue синий цвет знаков
func PrintBlue(msg ...interface{}) {
	printAny("\033[0;34m", msg)
}

// PrintMagenta фиолетовый цвет знако
func PrintMagenta(msg ...interface{}) {
	printAny("\033[0;35m", msg)
}

// PrintCyan цвет морской волны знаков
func PrintCyan(msg ...interface{}) {
	printAny("\033[0;36m", msg)
}

// PrintGrey серый цвет знаков
func PrintGrey(msg ...interface{}) {
	printAny("\033[0;37m", msg)
}

// PrintBlackBold чёрный цвет знаков (жирный)
func PrintBlackBold(msg ...interface{}) {
	printAny("\033[1;30m", msg)
}

// PrintRedBold красный цвет знаков (жирный)
func PrintRedBold(msg ...interface{}) {
	printAny("\033[1;31m", msg)
}

// PrintGreenBold зелёный цвет знаков (жирный)
func PrintGreenBold(msg ...interface{}) {
	printAny("\033[1;32m", msg)
}

// PrintYellowBold желтый цвет знаков (жирный)
func PrintYellowBold(msg ...interface{}) {
	printAny("\033[1;33m", msg)
}

// PrintBlueBold синий цвет знаков (жирный)
func PrintBlueBold(msg ...interface{}) {
	printAny("\033[1;34m", msg)
}

// PrintMagentaBold фиолетовый цвет знаков (жирный)
func PrintMagentaBold(msg ...interface{}) {
	printAny("\033[1;35m", msg)
}

// PrintCyanBold цвет морской волны знаков (жирный)
func PrintCyanBold(msg ...interface{}) {
	printAny("\033[1;36m", msg)
}

// PrintGreyBold серый цвет знаков (жирный)
func PrintGreyBold(msg ...interface{}) {
	printAny("\033[1;37m", msg)
}

// PrintBlackBackground чёрный цвет фона
func PrintBlackBackground(msg ...interface{}) {
	printAny("\033[0;40m", msg)
}

// PrintRedBackground красный цвет фона
func PrintRedBackground(msg ...interface{}) {
	printAny("\033[0;41m", msg)
}

// PrintGreenBackground зелёный цвет фона
func PrintGreenBackground(msg ...interface{}) {
	printAny("\033[0;42m", msg)
}

// PrintYellowBackground желтый цвет фона
func PrintYellowBackground(msg ...interface{}) {
	printAny("\033[0;43m", msg)
}

// PrintBlueBackground синий цвет фона
func PrintBlueBackground(msg ...interface{}) {
	printAny("\033[0;44m", msg)
}

// PrintMagentaBackground фиолетовый цвет фона
func PrintMagentaBackground(msg ...interface{}) {
	printAny("\033[0;45m", msg)
}

// PrintCyanBackground цвет морской волны фона
func PrintCyanBackground(msg ...interface{}) {
	printAny("\033[0;46m", msg)
}

// PrintGreyBackground серый цвет фона
func PrintGreyBackground(msg ...interface{}) {
	printAny("\033[0;47m", msg)
}

// PrintBlackBackgroundBold чёрный цвет фона (жирный)
func PrintBlackBackgroundBold(msg ...interface{}) {
	printAny("\033[1;40m", msg)
}

// PrintRedBackgroundBold красный цвет фона (жирный)
func PrintRedBackgroundBold(msg ...interface{}) {
	printAny("\033[1;41m", msg)
}

// PrintGreenBackgroundBold зелёный цвет фона (жирный)
func PrintGreenBackgroundBold(msg ...interface{}) {
	printAny("\033[1;42m", msg)
}

// PrintYellowBackgroundBold желтый цвет фона (жирный)
func PrintYellowBackgroundBold(msg ...interface{}) {
	printAny("\033[1;43m", msg)
}

// PrintBlueBackgroundBold синий цвет фона (жирный)
func PrintBlueBackgroundBold(msg ...interface{}) {
	printAny("\033[1;44m", msg)
}

// PrintMagentaBackgroundBold фиолетовый цвет фона (жирный)
func PrintMagentaBackgroundBold(msg ...interface{}) {
	printAny("\033[1;45m", msg)
}

// PrintCyanBackgroundBold цвет морской волны фона (жирный)
func PrintCyanBackgroundBold(msg ...interface{}) {
	printAny("\033[1;46m", msg)
}

// PrintGreyBackgroundBold серый цвет фона (жирный)
func PrintGreyBackgroundBold(msg ...interface{}) {
	printAny("\033[1;47m", msg)
}

// PrintBlackBackgroundUnderscore чёрный цвет фона с подчеркиванием
func PrintBlackBackgroundUnderscore(msg ...interface{}) {
	printAny("\033[4;40m", msg)
}

// PrintRedBackgroundUnderscore красный цвет фона с подчеркиванием
func PrintRedBackgroundUnderscore(msg ...interface{}) {
	printAny("\033[4;41m", msg)
}

// PrintGreenBackgroundUnderscore зеленый цвет фона с подчеркиванием
func PrintGreenBackgroundUnderscore(msg ...interface{}) {
	printAny("\033[4;42m", msg)
}

// PrintYellowBackgroundUnderscore желтый цвет фона с подчеркиванием
func PrintYellowBackgroundUnderscore(msg ...interface{}) {
	printAny("\033[4;43m", msg)
}

// PrintBlueBackgroundUnderscore синий цвет фона с подчеркиванием
func PrintBlueBackgroundUnderscore(msg ...interface{}) {
	printAny("\033[4;44m", msg)
}

// PrintMagentaBackgroundUnderscore фиолетовый цвет фона с подчеркиванием
func PrintMagentaBackgroundUnderscore(msg ...interface{}) {
	printAny("\033[4;45m", msg)
}

// PrintCyanBackgroundUnderscore цвет морской волны фона с подчеркиванием
func PrintCyanBackgroundUnderscore(msg ...interface{}) {
	printAny("\033[4;46m", msg)
}

// PrintGreyBackgroundUnderscore серый цвет фона с подчеркиванием
func PrintGreyBackgroundUnderscore(msg ...interface{}) {
	printAny("\033[4;47m", msg)
}

// PrintBlackUnderscore чёрный цвет знаков
func PrintBlackUnderscore(msg ...interface{}) {
	printAny("\033[4;30m", msg)
}

// PrintRedUnderscore красный цвет знаков
func PrintRedUnderscore(msg ...interface{}) {
	printAny("\033[4;31m", msg)
}

// PrintGreenUnderscore зелёный цвет знаков
func PrintGreenUnderscore(msg ...interface{}) {
	printAny("\033[4;32m", msg)
}

// PrintYellowUnderscore желтый цвет знаков
func PrintYellowUnderscore(msg ...interface{}) {
	printAny("\033[4;33m", msg)
}

// PrintBlueUnderscore синий цвет знаков
func PrintBlueUnderscore(msg ...interface{}) {
	printAny("\033[4;34m", msg)
}

// PrintMagentaUnderscore фиолетовый цвет знако
func PrintMagentaUnderscore(msg ...interface{}) {
	printAny("\033[4;35m", msg)
}

// PrintCyanUnderscore цвет морской волны знаков
func PrintCyanUnderscore(msg ...interface{}) {
	printAny("\033[4;36m", msg)
}

// PrintGreyUnderscore серый цвет знаков
func PrintGreyUnderscore(msg ...interface{}) {
	printAny("\033[4;37m", msg)
}
