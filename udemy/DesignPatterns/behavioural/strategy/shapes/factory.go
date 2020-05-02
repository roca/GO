package shapes

import (
	"fmt"
	"os"

	"example.com/strategy"
)

const (
	CONSOLE_STRATEGY = "console"
	IMAGE_STRATEGY   = "image"
)

func CreateShapePrintStrategy(strategyType string) (strategy.IPrintStrategy, error) {

	switch strategyType {
	case CONSOLE_STRATEGY:
		return &ConsoleSquare{
			PrintOutPut: strategy.PrintOutPut{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			PrintOutPut: strategy.PrintOutPut{
				LogWriter: os.Stdout,
			},
			DestinationFilePath: "./image.jpg",
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", strategyType)
	}

}
