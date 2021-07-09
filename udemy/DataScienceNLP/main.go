package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
	"github.com/cdipaolo/goml/text"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/roca/must"
)

func main() {
	// languageDetectionExample01()
	//languageDetectionExample02()
	//filesExmaples()
	// tokenizationExample02()
	//nerExample01()
	//nlpExample01()
	//rakeExample01()
	//summaryExample()
	// sentimentExample01()
	//sentimentExample02()
	//sentimentExample03()
	//StatsExample01()
	//dataAnalysisExample01()
	//onumExample01()
	// CleanDataExample01() // First step
	//gomlExample01()
	textClassifierExample01()

}

func textClassifierExample01() {
	// Use CSV

	// Normal
	// Stream
	// Chan: channel to communicate between go routines
	stream := make(chan base.TextDatapoint)
	// Error
	errors := make(chan error)

	// Init Model
	model := text.NewNaiveBayes(stream, 2, base.OnlyWordsAndNumbers)

	// Training
	go model.OnlineLearn(errors)
	// Classify a text as either hardware or software
	// Parse our stream of data
	// {X: Features , Y: Label}
	stream <- base.TextDatapoint{X: "He bought a computer with some memory", Y: 1}
	stream <- base.TextDatapoint{X: "It needs more memory", Y: 1}
	stream <- base.TextDatapoint{X: "computer", Y: 1}
	stream <- base.TextDatapoint{X: "code", Y: 0}
	stream <- base.TextDatapoint{X: "memory", Y: 1}
	stream <- base.TextDatapoint{X: "bugs", Y: 0}
	stream <- base.TextDatapoint{X: "I found some bugs in the code", Y: 0}
	stream <- base.TextDatapoint{X: "I swapped the memory", Y: 1}
	stream <- base.TextDatapoint{X: "I tested the code", Y: 0}

	// Close
	close(stream)

	for {
		err, more := <-errors
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Errors:", more)
			break
		}
	}

	// Labels: Hardware 1, software 0
	ex1 := "John bought a new computer" // Hardware 1
	ex2 := "He fixed the bugs in the code" // Software 0


	pred1 := model.Predict(ex1)
	pred2 := model.Predict(ex2)

	fmt.Println("'John bought a new computer'",pred1,"Prediction should be Hardware 1")
	fmt.Println("'He fixed the bugs in the code'",pred2,"Prediction should be Software 0")

	// Save model
	_ = model.PersistToFile("data/hardwareSoftwareClassifier") // Default json
}

func gomlExample01() {
	// Load out train/test datasets
	xTrain, yTrain, err := base.LoadDataFromCSV("data/train.csv")
	if err != nil {
		log.Fatal(err)
	}
	xTest, yTest, err := base.LoadDataFromCSV("data/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Training Dataset")
	fmt.Printf("X Train %T \n", xTrain)
	fmt.Printf("Y Train %T \n", yTrain)
	fmt.Println("Testing Dataset")
	fmt.Printf("X Test %T  \n", xTest)
	fmt.Printf("Y Test %T  \n", yTest)

	// Initialize Model
	//  Optimization Method ()
	//  Learning Rate
	//  Regularization : for overfitting
	//  Dataset (Xfeatures) [][]float64
	//  Class(binar 0 and 1)
	// model := linear.NewLogistic(base.BatchGA, 0.00001, 0, 1000, xTrain, yTrain)

	// Train
	// must.ReturnElseLogFatal(model.Learn)

	// Prediction
	// s1 := xTest[0]   //Should be Negative result of 0
	// s2 := xTest[169] // Should be Positive result of 1

	// mypred1 := must.ReturnElseLogFatal(model.Predict, s1).([]float64)
	// fmt.Println("Prediction 1 expected 0 to equal", math.Round(mypred1[0]))
	// mypred2 := must.ReturnElseLogFatal(model.Predict, s2).([]float64)
	// fmt.Println("Prediction 2 expected 1 to equal", math.Round(mypred2[0]))

	// Save Model
	// fmt.Println("Saved model to file logisticHCVmodel.json")
	// model.PersistToFile("data/logisticHCVmodel.json")

	// Evaluate
	cm := evaluateModel(xTrain, yTrain, xTest, yTest)
	fmt.Println("ConfusionMatrix")
	fmt.Println(cm)
	json, _ := json.MarshalIndent(cm, " ", " ")
	fmt.Println(string(json))
}

type ConfusionMatrix struct {
	Total         int     `json:"Total"`
	TotalTrain    int     `json:"TotalTrain"`
	TotalTest     int     `json:"TotalTest"`
	Positive      int     `json:"Positive"`
	Negative      int     `json:"Negative"`
	TruePositive  int     `json:"TruePositive"`
	TrueNegative  int     `json:"TrueNegative"`
	FalsePositive int     `json:"FalsePositive"`
	FalseNegative int     `json:"FalseNegative"`
	Accuracy      float64 `json:"accuracy"`
	Precision     float64 `json:"precision"`
	Recall        float64 `json:"recall"`
}

func evaluateModel(xTrain [][]float64, yTrain []float64, xTest [][]float64, yTest []float64) ConfusionMatrix {
	model := linear.NewLogistic(base.BatchGA, 0.00001, 0, 1000, xTrain, yTrain)
	must.ReturnElseLogFatal(model.Learn)
	fmt.Println("Finishing Training")
	// Save Model
	fmt.Println("Saved model to file logisticHCVmodel.json")
	model.PersistToFile("data/logisticHCVmodel.json")

	cm := ConfusionMatrix{}
	for _, y := range yTest {
		if y == 1.0 {
			cm.Positive++
		}
		if y == 1 {
			cm.Negative++
		}
	}

	decisionBoundary := 0.5
	for i, v := range xTest {
		prediction := must.ReturnElseLogFatal(model.Predict, v).([]float64)
		y := int(yTest[i])
		Positive := prediction[0] >= decisionBoundary

		if y == 1 && Positive {
			cm.TruePositive++
		}
		if y == 1 && !Positive {
			cm.FalseNegative++
		}
		if y == 0 && Positive {
			cm.FalsePositive++
		}
		if y == 0 && !Positive {
			cm.TrueNegative++
		}
	}
	cm.Accuracy = (float64(cm.TruePositive) + float64(cm.TrueNegative)) / (float64(cm.TruePositive) + float64(cm.TrueNegative) + float64(cm.FalsePositive) + float64(cm.FalseNegative))

	cm.Precision = float64(cm.TruePositive) / (float64(cm.TruePositive) + float64(cm.FalsePositive))
	cm.Recall = float64(cm.TruePositive) / (float64(cm.TruePositive) + float64(cm.FalseNegative))
	cm.TotalTest = len(yTest)
	cm.TotalTrain = len(yTrain)
	cm.Total = len(yTest) + len(yTrain)

	return cm
}

func CleanDataExample01() {

	/* Requirement For Dataset:
	   1) Numerical
	   2) No missing value & Header
	   3) Train/Test CSV
	   4) [features]<target>
	*/
	// Load our dataset
	// 	// Open CSV
	mode := int(0644)
	m := fs.FileMode(mode)
	csvfile := must.ReturnElseLogFatal(os.Open, "data/hcvdat0.csv").(*os.File)
	trainFile := must.ReturnElseLogFatal(os.OpenFile, "data/train.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, m).(*os.File)
	testFile := must.ReturnElseLogFatal(os.OpenFile, "data/test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, m).(*os.File)
	defer csvfile.Close()
	defer trainFile.Close()
	defer testFile.Close()

	// Read CSV
	df := dataframe.ReadCSV(csvfile)
	//fmt.Println(df)
	fmt.Println(df.Names())
	fmt.Println(df.Types())

	sexMap := map[string]int{"f": 0, "m": 1}
	options := map[string]interface{}{"colMap": sexMap}
	df = MutateColInt("Sex", &df, options)

	donarMap := map[string]int{
		"0=Blood Donor":          0,
		"0s=suspect Blood Donor": 1,
		"1=Hepatitis":            1,
		"2=Fibrosis":             1,
		"3=Cirrhosis":            1,
	}
	options["colMap"] = donarMap
	df = NewColInt("Category", "Target", &df, options)

	for _, name := range df.Names() {
		if df.Col(name).HasNaN() {
			df = df.Mutate(ReplacMissingValuesFloat(df.Col(name)))
		}
	}

	// Drop first two Columns
	df = df.Drop([]int{0, 1})
	fmt.Println(df.Subset([]int{0, 1}))

	// Get random sample of train and test data
	train, test := Sample(df, .7, time.Now().UTC().UnixNano())
	fmt.Print("Total: ")
	fmt.Println(df.Dims())

	fmt.Print("Train: ")
	fmt.Println(train.Dims())
	noHeaderOption := dataframe.WriteHeader(false)
	must.ReturnElseLogFatal(train.WriteCSV, trainFile, noHeaderOption)

	fmt.Print("Test: ")
	fmt.Println(test.Dims())
	must.ReturnElseLogFatal(test.WriteCSV, testFile, noHeaderOption)

}
func Sample(df dataframe.DataFrame, percentage float64, seed int64) (dataframe.DataFrame, dataframe.DataFrame) {
	r, _ := df.Dims()

	sampleSize := int(float64(r) * percentage)
	sampleIndicesMap := map[int]int{}
	sampleIndices := []int{}
	rand.Seed(seed)

	for len(sampleIndicesMap) < sampleSize {
		i := rand.Intn(r)
		sampleIndicesMap[i] = i
	}
	for _, v := range sampleIndicesMap {
		sampleIndices = append(sampleIndices, v)
	}
	sort.Ints(sampleIndices)

	testIndices := []int{}
	for i, _ := range df.Maps() {
		found := false
		for _, j := range sampleIndices {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			testIndices = append(testIndices, i)
		}
	}

	return df.Subset(sampleIndices), df.Subset(testIndices)
}
func ReplacMissingValuesFloat(s series.Series) series.Series {
	if !s.HasNaN() {
		return s
	}

	newValues := []float64{}

	for i, isNaN := range s.IsNaN() {
		newValues = append(newValues, s.Elem(i).Float())
		if isNaN {
			newValues[i] = 0.
		}

	}
	newSeries := series.New(newValues, series.Float, s.Name)

	return newSeries
}
func MutateColInt(name string, df *dataframe.DataFrame, options ...map[string]interface{}) dataframe.DataFrame {
	//fmt.Println(df.Col(name))
	colMap := make(map[string]int)

	if len(options) > 0 {
		if v, found := options[0]["colMap"]; found {
			colMap = v.(map[string]int)
		}
	} else {
		for _, v := range df.Col(name).Records() {
			colMap[v] = 0
		}
		i := 0
		for k := range colMap {
			i++
			colMap[k] = i
		}
	}

	//fmt.Println(colMap)
	var newColValues = []int{}
	for _, v := range df.Col(name).Records() {
		newColValues = append(newColValues, colMap[v])
	}
	//fmt.Println(newColValues)

	s := series.New(newColValues, series.Int, name)

	mut := df.Mutate(s)

	return mut

}

func NewColInt(name string, newname string, df *dataframe.DataFrame, options ...map[string]interface{}) dataframe.DataFrame {
	//fmt.Println(df.Col(name))
	colMap := make(map[string]int)

	if len(options) > 0 {
		if v, found := options[0]["colMap"]; found {
			colMap = v.(map[string]int)
		}
	} else {
		for _, v := range df.Col(name).Records() {
			colMap[v] = 0
		}
		i := 0
		for k := range colMap {
			i++
			colMap[k] = i
		}
	}

	//fmt.Println(colMap)
	var newColValues = []int{}
	for _, v := range df.Col(name).Records() {
		newColValues = append(newColValues, colMap[v])
	}
	//fmt.Println(newColValues)

	s := series.New(newColValues, series.Int, newname)

	mut := df.Mutate(s)

	return mut

}

// func gonumExample01() {
// 	// Scalar: A number
// 	var a int = 44
// 	fmt.Println("Scalar:", a)
// 	// Vectors
// 	// 1 Dim
// 	// Row Vector/Column Vector[list of same datatype]
// 	// Method 1
// 	myvector := []float64{1.2, 3.4, 4.5, 3.5, 4.4}
// 	fmt.Println("Vector:", myvector)
// 	fmt.Printf("%T \n", myvector)

// 	// Method 2: Using Gonum
// 	myvectorA := mat.NewVecDense(2, []float64{1.2, 3.4})
// 	myvectorB := mat.NewVecDense(2, []float64{3.2, 4.4})
// 	fmt.Println("Vector A", myvectorA)
// 	fmt.Println("Vector B", myvectorA)

// 	// Dot Product
// 	dp := mat.Dot(myvectorA, myvectorB)
// 	fmt.Println("Dot product of A and B:", dp)

// 	// Matrix
// 	// Creating a Matrix
// 	// Linear wrapped
// 	data := []float64{
// 		1.1, 1.2, 1.3,
// 		2.1, 2.2, 2.3,
// 	}
// 	mymatrix := mat.NewDense(2, 3, data)
// 	//fmt.Println(mymatrix)
// 	// Formatted
// 	fmt.Println("Matrices:\n", mat.Formatted(mymatrix))
// 	fmt.Println("Row 0",mat.Row(nil, 0, mymatrix))
// 	fmt.Println("Col 0",mat.Col(nil, 0, mymatrix))

// 	// Matrix of Zeros
// 	fmt.Println(mat.Formatted(mat.NewDense(3, 3, nil)))
// }

// func dataAnalysisExample01() {
// 	// Data Analysis in Go
// 	// Open CSV
// 	csvfile := must.ReturnElseLogFatal(os.Open, "data/diamonds.csv").(*os.File)
// 	defer csvfile.Close()
// 	// Read CSV
// 	df := dataframe.ReadCSV(csvfile)
// 	// fmt.Println(df)
// 	// EDA

// 	// Shape of Data
// 	row, col := df.Dims()
// 	fmt.Println("rows:", row, ",columns:", col)

// 	// Get only row size
// 	fmt.Println("row size:", df.Nrow())

// 	// Get only col size
// 	fmt.Println("col size:", df.Ncol())

// 	// Get column names
// 	fmt.Println("col names:", df.Names())

// 	// // Get DataTypes
// 	// fmt.Println("DataTypes:", df.Types())

// 	// // Describe/Summary
// 	// fmt.Println("Describe/Summary:", df.Describe())

// 	// Selection of Columns & Rows
// 	// Select columns by Column name
// 	//fmt.Println("Carats", df.Select("carat"))

// 	// Select column by index
// 	// fmt.Println("Carats index 0", df.Select(0))

// 	// Multiple column selection with slice of strings
// 	// fmt.Print(df.Select([]string{"carat","cut"}))

// 	// Selection of single row
// 	// fmt.Println(df.Subset(0))

// 	// Selection of multiple rows
// 	// fmt.Println(df.Subset([]int{0,2,4}))

// 	// Series and apply functions
// 	// ds := df.Col("carat")
// 	// fmt.Printf("%T \n", ds)
// 	// fmt.Println(ds)

// 	// Apply function 'Mean'  to the series
// 	// dsmean := ds.Mean()
// 	// fmt.Println("Mean of carat series:", dsmean)

// 	// Check for missing values
// 	// fmt.Printf("There are %d missing values in this series\n", len(ds.IsNaN()))

// 	// gmean := stat.Mean(ds.Float(),nil)
// 	// fmt.Println("Go 'num' package Mean series",gmean)

// 	//  Apply Conditions/Filter
// 	// type F struct {
// 	// 	Colname string
// 	// 	Comparator series.Comparator
// 	// 	Comparando interface{}
// 	// }
// 	fmt.Println(df.Select("cut"))
// 	isPremium := df.Filter(dataframe.F{
// 		Colname:    "cut",
// 		Comparator: series.Eq,
// 		Comparando: "Premium",
// 	})
// 	fmt.Println(isPremium.Dims())

// }

// func StatsExample01() {
// 	fmt.Println("Statistics in Go")
// 	even := []float64{2, 4, 6, 8, 10, 8, 8}
// 	var odd = []float64{1, 3, 5, 7, 9, 7, 7}

// 	// Basic Math
// 	// Mean

// 	evenmean := must.ReturnElseLogFatal(stats.Mean, even).(float64)
// 	oddmean := must.ReturnElseLogFatal(stats.Mean, odd).(float64)
// 	evenmax := must.ReturnElseLogFatal(stats.Max, even).(float64)
// 	oddmax := must.ReturnElseLogFatal(stats.Max, odd).(float64)
// 	evenmode := must.ReturnElseLogFatal(stats.Mode, even).([]float64)
// 	oddmode := must.ReturnElseLogFatal(stats.Mode, odd).([]float64)
// 	fmt.Println("Even:", even, "Mean:", evenmean, "Max:", evenmax, "Mode:", evenmode)
// 	fmt.Println("Odd:", odd, "Mean:", oddmean, "Max:", oddmax, "Mode:", oddmode)

// 	// Stats package also has Arithmetic, Harm and Geo mean
// 	// std, variance

// }

// type savedDetails struct {
// 	Sentence   string
// 	Label      string
// 	Vaderlabel float64
// }

// func sentimentExample03() {

// 	// Open File
// 	csvfile := must.ReturnElseLogFatal(os.Open, "data/amazondataset.csv").(*os.File)
// 	defer csvfile.Close()

// 	// Method 1: Read our CSV File with 'Gota' dataframe
// 	// df := dataframe.ReadCSV(csvfile)
// 	// fmt.Println(df.Names())
// 	// fmt.Println(df.Select("sentences").String()[0])

// 	// Method 2: Read our CSV File with 'csv'
// 	detailsList := []savedDetails{}

// 	csvr := csv.NewReader(csvfile)
// 	csvLines := must.ReturnElseLogFatal(csvr.ReadAll).([][]string)
// 	for _, line := range csvLines {
// 		sentence := line[0]
// 		label := line[1]
// 		// fmt.Println(sentence, "[Sentiment: {Orig:", label, ",NewLabel", analyze(sentence), "}]")
// 		detailsList = append(detailsList, savedDetails{
// 			Sentence:   sentence,
// 			Label:      label,
// 			Vaderlabel: analyze(sentence),
// 		})
// 	}

// 	// Apply our Fxn

// 	// Results as A DataFrame
// 	//  Create Slice/Struct to store values
// 	//  Struct to Dataframe
// 	df := dataframe.LoadStructs(detailsList)
// 	fmt.Println(df)
// 	//  Save using Gota
// 	f := must.ReturnElseLogFatal(os.Create, "data/newdata.csv").(*os.File)
// 	df.WriteCSV(f)
// }

// func analyze(text string) float64 {
// 	parseText := sentitext.Parse(string(text), lexicon.DefaultLexicon)
// 	results := sentitext.PolarityScore(parseText)
// 	return results.Compound
// }

// func sentimentExample01() {
// 	//content := must.ReturnElseLogFatal(ioutil.ReadFile, "AiHistory.txt").([]byte)
// 	content := "I hate apples and coding"
// 	parseText := sentitext.Parse(string(content), lexicon.DefaultLexicon)
// 	results := sentitext.PolarityScore(parseText)
// 	fmt.Println("SENTIMENT POLARITY SCORE:", results)
// 	fmt.Println("Positive:", results.Positive)
// 	fmt.Println("Negative:", results.Negative)
// 	fmt.Println("Neutral:", results.Neutral)
// 	fmt.Println("Sentiment/Compound:", results.Compound)

// }
// func sentimentExample02() {
// 	content := "I hate apples and coding"
// 	sentimentModel, err := sentiment.Restore()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	results := sentimentModel.SentimentAnalysis(content, sentiment.English)
// 	fmt.Println(results)
// 	//Sentiment for the whole sentence
// 	fmt.Println("Sentiment Score:", results.Score)

// }
// func summaryExample() {

// 	content := must.ReturnElseLogFatal(ioutil.ReadFile, "AiHistory.txt").([]byte)

// 	summary := tldr.New()
// 	results, _ := summary.Summarize(string(content), 1)
// 	// fmt.Println(string(content))
// 	fmt.Println("Summary: ", results)

// }
// func rakeExample01() {
// 	// var myText string = `Natural language processing (NLP) is a subfield of linguistics, computer science, and artificial intelligence concerned with the interactions between computers and human language, in particular how to program computers to process and analyze large amounts of natural language data. The result is a computer capable of "understanding" the contents of documents, including the contextual nuances of the language within them. The technology can then accurately extract information and insights contained in the documents as well as categorize and organize the documents themselves.

// 	// Challenges in natural language processing frequently involve speech recognition, natural language understanding, and natural-language generation.
// 	// Natural language processing (NLP) is a subfield of linguistics, computer science, and artificial intelligence concerned with the interactions between computers and human language, in particular how to program computers to process and analyze large amounts of natural language data. The result is a computer capable of "understanding" the contents of documents, including the contextual nuances of the language within them. The technology can then accurately extract information and insights contained in the documents as well as categorize and organize the documents themselves.

// 	// Challenges in natural language processing frequently involve speech recognition, natural language understanding, and natural-language generation.
// 	// A major drawback of statistical methods is that they require elaborate feature engineering. Since the early 2010s,[16] the field has thus largely abandoned statistical methods and shifted to neural networks for machine learning. Popular techniques include the use of word embeddings to capture semantic properties of words, and an increase in end-to-end learning of a higher-level task (e.g., question answering) instead of relying on a pipeline of separate intermediate tasks (e.g., part-of-speech tagging and dependency parsing). In some areas, this shift has entailed substantial changes in how NLP systems are designed, such that deep neural network-based approaches may be viewed as a new paradigm distinct from statistical natural language processing. For instance, the term neural machine translation (NMT) emphasizes the fact that deep learning-based approaches to machine translation directly learn sequence-to-sequence transformations, obviating the need for intermediate steps such as word alignment and language modeling that was used in statistical machine translation (SMT).
// 	// `
// 	// words := rake.RunRake(myText)
// 	// keyWordMap := make(map[string]float64)
// 	// for _,word := range words {
// 	// 	// fmt.Println("Keyword: ",word.Key,"Score: ",word.Value)
// 	// 	keyWordMap[word.Key] = word.Value
// 	// }
// 	// fmt.Println(keyWordMap)

// 	// Method: 1
// 	fmt.Println("Method1: Using 'ioutil' package")
// 	content, err := ioutil.ReadFile("exampletext.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(content))
// 	words := rake.RunRake(string(content))
// 	// keyWordMap := make(map[string]float64)
// 	for _, word := range words {
// 		fmt.Println("Keyword: ", word.Key, "Score: ", word.Value)
// 		// keyWordMap[word.Key] = word.Value
// 	}
// 	// fmt.Println(keyWordMap)

// }

// func nlpExample01() {
// 	var myText string = "Hello world this is Golang"

// 	// NLP Document Struct
// 	doc := must.ReturnElseLogFatal(prose.NewDocument, myText).(*prose.Document)
// 	fmt.Printf("%T \n", doc)

// 	for i, tok := range doc.Tokens() {
// 		fmt.Println("Index:", i, "Tokens:", tok.Text, "Tag:", tok.Tag, "Label:", tok.Label)
// 	}

// 	// Reading from a TextFile
// 	content := must.ReturnElseLogFatal(ioutil.ReadFile, "example.txt").([]byte)
// 	//content, _ := ioutil.ReadFile("example.txt")
// 	fmt.Println(string(content))
// 	doc2 := must.ReturnElseLogFatal(prose.NewDocument, string(content)).(*prose.Document)

// 	// Word Tokens
// 	for i, tok := range doc2.Tokens() {
// 		fmt.Println("Index:", i, "Tokens:", tok.Text, "Tag:", tok.Tag, "Label:", tok.Label)
// 	}

// 	// Sentence Tokens
// 	// for i, sentence := range doc2.Sentences(){
// 	// 	fmt.Println("Index:", i, "Sentence:", sentence.Text)
// 	// }

// }

// func nerExample01() {
// 	// NER
// 	// Entity: Person/People/Org/Location/etc

// 	var myText string = "John Mark works in London as a Go developer"

// 	// NLP Document Struct
// 	doc := must.ReturnElseLogFatal(prose.NewDocument, myText).(*prose.Document)

// 	for index, entity := range doc.Entities() {
// 		fmt.Println(index, entity.Text, entity.Label)
// 	}

// 	for index, token := range doc.Tokens() {
// 		fmt.Println(index, token.Text, token.Label)
// 	}

// }

// func tokenizationExample01() { // From Scratch

// 	// Method 1: strings.Split
// 	var myText string = "Paul wasn't coding at all"
// 	tokens := strings.Split(myText, " ")
// 	fmt.Println(tokens)

// 	// Method 2: Rule Based (Regex)
// 	r := regexp.MustCompile(`\w+`)
// 	tokens2 := r.FindAllString(myText, -1)
// 	fmt.Println(tokens2)

// 	// Method 3: Regex + Split
// 	r2 := regexp.MustCompile(`\s+`)
// 	tokens3 := r2.Split(myText, -1)
// 	fmt.Println(tokens3)

// 	// Method 4: Using pros
// 	tokenizer := tokenize.NewTreebankWordTokenizer()
// 	tokens4 := tokenizer.Tokenize(myText)
// 	for _, tok := range tokens4 {
// 		fmt.Println(tok)
// 	}
// 	fmt.Println(tokens4)
// }

// func tokenizationExample02() { // Using prose
// 	myText := "Jesse was going to fish a fish at the bank in London"

// 	// Tokens
// 	tokenizer := tokenize.NewTreebankWordTokenizer()
// 	tokens := tokenizer.Tokenize(myText)
// 	fmt.Println(tokens)

// 	// Tags
// 	postagger := tag.NewPerceptronTagger()
// 	tags := postagger.Tag(tokens)
// 	for _, token := range tags {
// 		fmt.Println(token.Text, token.Tag)
// 	}

// 	fmt.Println("Noun Chunks::", getChunks(myText, "NN"))
// 	fmt.Println("Verb Chunks::", getChunks(myText, "V"))

// 	regex := chunk.TreebankNamedEntities
// 	// Loop: tag + reg == Named Entity Chunks
// 	for _, entity := range chunk.Chunk(postagger.Tag(tokens), regex) {
// 		fmt.Println(entity)
// 	}

// }

// func getChunks(text string, tagName string) []string {
// 	// Tokenize
// 	tokens := tokenize.NewTreebankWordTokenizer().Tokenize(text)

// 	// Tags
// 	tags := tag.NewPerceptronTagger().Tag(tokens)

// 	// if tag ==  requested tagName
// 	chunks := []string{}
// 	for _, token := range tags {
// 		if strings.HasPrefix(token.Tag, tagName) {
// 			chunks = append(chunks, token.Text)
// 		}
// 	}
// 	return chunks
// }

// func languageDetectionExample02() { // Using github.com/abadojack/whatlanggo
// 	var mydocx string = "Hello world of Go"
// 	lang := whatlanggo.Detect(mydocx)
// 	fmt.Println("Text:", mydocx)
// 	fmt.Println("whatlango: ", lang.Lang.String()) // Language name
// 	fmt.Println("whatlango: ", lang.Confidence)    // Confidence/Accuracy of prediction
// }

// func languageDetectionExample01() { // Using github.com/rylans/getlang
// 	var mystr string = "Hello world of Go"
// 	// var mystrfr string = "Bonjour a tous"

// 	lang := getlang.FromString(mystr)
// 	// lang2 := getlang.FromString(mystrfr)

// 	fmt.Println("Text:", mystr)
// 	fmt.Println("getlang: ", lang.LanguageCode()) // Language code
// 	fmt.Println("getlang: ", lang.Confidence())   // Confidence/Accuracy of prediction

// 	// fmt.Println("Text:", mystrfr)
// 	// fmt.Println("getlang: ",lang2.LanguageCode()) // Language code
// 	// fmt.Println("getlang: ",lang2.Confidence())   // Confidence/Accuracy of prediction
// }

// func textCleaningExamples01() { // Using github.com/mingrammer/commonregex
// 	// Textcleaning using Regex & more
// 	// var mystr string = "Hello GoDev my email is jharis@gmail.com"
// 	// Multi line large text string literal ``
// 	// 	var docx string = `
// 	// 	Golang was designed at Google by Robert Griesemer, Rob Pike,
// 	//  and Ken Thompson. Ken called Rob on 519-555-7765 which was redirected to +44 22 777 555.
// 	// Jesse sent an email to jc.@gmail.com which he found on the website http://jcharistech.com.
// 	// Golang was publicly announced in November 2009 and version 1.0 was released in March 2012.
// 	// Go is widely used in production at Google USA and in many other organizations and open-source projects.
// 	// In November 2016, the Go and Go Mono fonts were released by type designers Charles Bigelow and Kris Holmes specifically for use by the Go project. Go is a humanist sans-serif which resembles Lucida Grande and Go Mono is monospaced. Each of the fonts adhere to the WGL4 character set and were designed to be legible with a large x-height and distinct letterforms. Both Go and Go Mono adhere to the DIN 1450 standard by having a slashed zero, lowercase l with a tail, and an uppercase I with serifs.
// 	// I have been coding since 4:00 AM this morning.Accra is big but not bigger as London.
// 	// john.smith@yahoo.com
// 	// 	`

// 	/*
// 		Reading text from a file
// 		os, ioutil, bufio
// 	*/
// 	content, err := ioutil.ReadFile("example.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	docx2 := string(content)
// 	fmt.Println("Sample:::", docx2)

// 	// Text Preprocessing
// 	// Normalizing: uniform case, removing unicode chars
// 	// fmt.Println(strings.ToLower(mystr))
// 	// fmt.Println(docx)
// 	// Remove noise [ special chars, eamils, phone #s]
// 	// Lemma/Stemming
// 	// Tokenization

// 	/*
// 	  Extra Emails
// 	  Method 1: Standard Library Regexp
// 	*/
// 	// Pattern
// 	// p := regexp.MustCompile(`GoDev`)
// 	// Find/Replace
// 	// fmt.Println(p.FindAllString(mystr, 1))
// 	// fmt.Println(p.ReplaceAllString(mystr, "REPLACED"))

// 	/*
// 	  Extra Emails
// 	  Method 2: External Library commonregex similar
// 	  to python library neatregex
// 	*/
// 	// fmt.Println(cregex.Emails(mystr))

// 	// Exercise 1.
// 	fmt.Println("Emails: ", cregex.Emails(docx2))

// 	// Remove/Replace : Document Redaction/Text Cleaning
// 	p := cregex.EmailRegex
// 	fmt.Println("ALTERED::: ", p.ReplaceAllString(docx2, "REPLACED"))

// }
// func stringExamples02() {
// 	var mystr string = "hello Go"
// 	fmt.Printf("Value: '%s'\n", mystr)
// 	fmt.Printf("Type: %T\n", mystr)
// 	fmt.Printf("Length: %d\n", len(mystr))
// 	fmt.Printf("Uppercase: %s\n", strings.ToUpper(mystr))
// 	fmt.Printf("Lowercase: %s\n", strings.ToLower(mystr))
// 	fmt.Printf("Titlecase: %s\n", strings.Title(mystr))
// 	fmt.Printf("Count 'l' occurrences: %d\n", strings.Count(mystr, "l"))
// 	fmt.Printf("Contains 'Go': %t\n", strings.Contains(mystr, "Go"))
// 	fmt.Printf("Split on ' ': %q\n", strings.Split(mystr, " "))
// 	fmt.Printf("Split after 'hel': %q\n", strings.SplitAfter(mystr, "hel"))
// 	fmt.Printf("Replace 'hello': %s\n", strings.ReplaceAll(mystr, "hello", "I love"))

// 	s := strings.Split(strings.ReplaceAll(mystr, "hello", "N.L.P programing"), " ")
// 	ss := strings.Join(s, " using ")
// 	fmt.Printf("Split and Join : %s\n", ss)

// }

// func stringExamples01() {
// 	fmt.Println("Method1: Create Characters with type ")
// 	var char byte = 'A'
// 	var char2 rune = 'A'
// 	fmt.Printf("Char as Byte %d %T\n", char, char)
// 	fmt.Printf("Char as Rune %d %T\n", char2, char2)

// 	fmt.Println("Method2: Create Characters with Method")
// 	charA := byte('A')
// 	charB := rune('A')
// 	fmt.Printf("Char as Byte:Fxn %d %T\n", charA, charA)
// 	fmt.Printf("Char as Rune:Fxn %d %T\n", charB, charB)

// 	fmt.Println("Actual Representation")
// 	str1 := string(char)
// 	str2 := string(char2)
// 	fmt.Printf("Char as String %s %T\n", str1, str1)
// 	fmt.Printf("Char as String %s %T\n", str2, str2)

// 	fmt.Println("Representation method 2 using Printf")
// 	fmt.Printf("Char as Byte %c %T\n", char, char)
// 	fmt.Printf("Char as Rune %c %T\n", char2, char2)
// }
// func filesExmaples() {
// 	args := os.Args[1:]
// 	if len(args) == 2 {
// 		switch args[1] {
// 		case "csv":
// 			files.OpenCSVFile(args[0])
// 		case "pdf":
// 			files.OpenPDFFile(args[0])
// 		case "txt":
// 			files.OpenTextFile(args[0])
// 		default:
// 			fmt.Println("Usage: file type[csv|txt|pdf]")
// 		}
// 	} else {
// 		fmt.Print("Usage: file type[csv|txt|pdf]")
// 	}
// }
