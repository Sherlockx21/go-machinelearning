package CSV

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func WithDataFrame() {
	irisFile, err := os.Open("./Data/DataCollection/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	var irisDF dataframe.DataFrame
	opt := dataframe.HasHeader(false)
	irisDF = dataframe.ReadCSV(irisFile, opt)
	err = irisDF.SetNames("sepal_length", "sepal_width", "petal_length", "petal_width", "species")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(irisDF)

	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
	}

	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Error() != nil {
		log.Fatal(versicolorDF.Error())
	}

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(versicolorDF)

}
