package main

import (
	"encoding/base64"
	"example16/ai"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

var prompt string = `
Analyze the attached satellite image of a property with these specific steps:

1. Residence identification: Locate the primary residence on the property by looking for:
   - The largest roofed structure 
   - Typical residential features (driveway connection, regular geometry)
   - Distinction from other structures (garages, sheds, pools)
   Describe the residence's location relative to property boundaries and other features.

2. Tree overhang analysis: Examine all trees near the primary residence:
   - Identify any trees whose canopy extends directly over any portion of the roof
   - Estimate the percentage of roof covered by overhanging branches (0-25%, 25-50%, 50-75%, 75-100%)
   - Note particularly dense areas of overhang

3. Fire risk assessment: For any overhanging trees, evaluate:
   - Potential wildfire vulnerability (ember catch points, continuous fuel paths to structure)
   - Proximity to chimneys, vents, or other roof openings if visible
   - Areas where branches create a "bridge" between wildland vegetation and the structure
   
4. Defensible space identification: Assess the property's overall vegetative structure:
   - Identify if trees connect to form a continuous canopy over or near the home
   - Note any obvious fuel ladders (vegetation that can carry fire from ground to tree to roof)

5. Fire risk rating: Based on your analysis, assign a Fire Risk Rating from 1-4:
   - Rating 1 (Low Risk): No tree branches overhanging the roof, good defensible space around the structure
   - Rating 2 (Moderate Risk): Minimal overhang (<25% of roof), some separation between tree canopies
   - Rating 3 (High Risk): Significant overhang (25-50% of roof), connected tree canopies, multiple points of vulnerability
   - Rating 4 (Severe Risk): Extensive overhang (>50% of roof), dense vegetation against structure, numerous ember catch points, limited defensible space

For each item above (1-5), write one sentence summarizing your findings, with your final response being the numeric Fire Risk Rating (1-4) with a brief justification.
`

func main() {

	byteData, err := ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	encodedString := base64.StdEncoding.EncodeToString(byteData)

	// fmt.Println("Base64 Encoded String:", encodedString)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_with_Image_And_message(conversation, prompt, encodedString)
	text, _, _ := ai.Chat(conversation, 1.0, nil, nil)
	fmt.Println("AI Response:", text)
}

func ReadFile(path string) ([]byte, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}
