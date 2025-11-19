package prompt_evaluator

import (
	"encoding/json"
	"example12/ai"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
)

var system_prompt_test_cases = anthropic.TextBlockParam{
	Text: `
You are a test case creator specializing in designing evaluation scenarios.
	`,
}
var system_prompt_ideas = anthropic.TextBlockParam{
	Text: `
You are a test scenario designer specialized in creating diverse, unique testing scenarios.
	`,
}

type Idea string

type PromptEvaluator struct {
}

func (pe *PromptEvaluator) RunTestCase(test_case TestCase) (string, string) {
	prompt := fmt.Sprintf(`Generate a one-day meal plan for an athlete that meets their dietary restrictions.

	<athlete_information>
	- Height: %s
	- Weight: %s
	- Goal: %s
	- Dietary restrictions: %s
	</athlete_information>

	Guidelines:
	1. Include accurate daily calorie amount
	2. Show protein, fat, and carb amouts
	3. Specify when to eat each meal
	4. Use only foods that fit restrictions
	5. List all portioin sizes in graams
	6. Keep budget-frendly if mentioned

Here is an example of poor input and output:

<example_poor_input>
- Height: 215
- Weight: 110
- Goal: muscle gain
- Dietary restrictions: vegan
<example_poor_input>

<example_poor_output>
Athlete Meal Plan Analysis and Design

Calorie Calculation:
- Basal Metabolic Rate (BMR): ~2400 calories
- Activity Level Adjustment: +500-600 calories
- Total Daily Caloric Need: ~3000 calories

Macronutrient Breakdown:
- Protein: 225g (30%)
- Carbohydrates: 375g (50%)
- Fats: 83g (20%)

MEAL PLAN

Breakfast (7:00 AM):
- Tofu Scramble
  * Extra firm tofu (250g)
  * Nutritional yeast (30g)
  * Spinach (100g)
  * Bell peppers (80g)
- Protein smoothie
  * Pea protein powder (50g)
  * Banana (150g)
  * Almond milk (250ml)
Calories: 750
Protein: 65g

Mid-Morning Snack (10:00 AM):
- Protein energy balls
  * Dates (50g)
  * Chia seeds (25g)
  * Hemp protein (40g)
Calories: 350
Protein: 35g

Lunch (1:00 PM):
- Quinoa black bean bowl
  * Quinoa (200g)
  * Black beans (150g)
  * Tempeh (150g)
  * Avocado (100g)
  * Mixed vegetables (150g)
Calories: 825
Protein: 60g

Post-Workout (4:00 PM):
- Vegan protein shake
  * Soy protein isolate (60g)
  * Oat milk (250ml)
  * Frozen berries (100g)
Calories: 400
Protein: 50g

Dinner (7:00 PM):
- Lentil and seitan stir-fry
  * Seitan (200g)
  * Red lentils (150g)
  * Brown rice (150g)
  * Broccoli (120g)
  * Coconut aminos (30ml)
Calories: 675
Protein: 65g

Total Daily Intake:
- Calories: 3000
- Protein: 275g
- Carbohydrates: 375g
- Fats: 85g

Key Considerations:
- 100% Vegan
- High-protein sources
- Balanced macronutrients
- Muscle gain focused
- Budget-friendly ingredients

Recommendations:
- Supplement with B12
- Stay hydrated
- Adjust portions as needed
</example_poor_output>

While the solution provides a structured meal plan with balanced nutrition, it fails to demonstrate how the plan was tailored to the specific athlete's parameters. The solution does not utilize the input map of height, weight, goal, and dietary restrictions, which are critical for creating a truly personalized nutrition strategy. The meal plan appears generic and not dynamically generated based on individual athlete needs.

Here is an example of good input and output:
		
<example_good_input> 
- Height: 192
- Weight: 118
- Goal: Muscle mass gain and high-performance strength training
- Dietary restrictions: No dairy, gluten-free diet
</example_good_input>

<example_good_output>
I'll create a comprehensive, personalized meal plan for the athlete based on the provided specifications. Here's a detailed approach:

ATHLETE PROFILE ANALYSIS:
- Height: 192 cm
- Weight: 118 kg
- Goal: Muscle mass gain and high-performance strength training
- Dietary Restrictions: No dairy, gluten-free diet

CALORIE & MACRONUTRIENT CALCULATION:
1. Basal Metabolic Rate (BMR) Calculation:
   - Using Mifflin-St Jeor Equation
   - BMR = (10 × weight) + (6.25 × height) - (5 × age) + 5
   - Assuming age 25: BMR ≈ 2,350 calories

2. Activity Level Adjustment:
   - Strength training, high-intensity: Multiply BMR by 1.7
   - Total Daily Energy Expenditure (TDEE): 3,995 calories
   - Muscle Gain Goal: Add 300-500 calories
   - Target Daily Intake: 4,300 calories

3. Macronutrient Breakdown:
   - Protein: 2.2g per kg body weight = 259g (30%)
   - Carbohydrates: 5.5g per kg body weight = 649g (60%)
   - Fats: 1.2g per kg body weight = 142g (10%)

MEAL PLAN:

Breakfast (7:00 AM) - 1,000 calories:
- Gluten-free oats (100g)
- Egg whites (200g)
- Almond butter (50g)
- Banana (150g)
- Gluten-free protein powder (50g)
- Coconut milk (250ml)
Macros: 
- Protein: 75g
- Carbs: 120g
- Fats: 35g

Mid-Morning Snack (10:00 AM) - 600 calories:
- Chicken breast (150g)
- Sweet potato (200g)
- Avocado (100g)
- Mixed nuts (50g)
Macros:
- Protein: 50g
- Carbs: 80g
- Fats: 30g

Lunch (1:00 PM) - 1,100 calories:
- Grass-fed beef (200g)
- Quinoa (150g)
- Roasted vegetables (200g)
- Olive oil (30ml)
- Almonds (50g)
Macros:
- Protein: 80g
- Carbs: 150g
- Fats: 40g

Post-Workout (4:00 PM) - 800 calories:
- Gluten-free protein shake (60g)
- Rice cakes (100g)
- Lean turkey (150g)
- Coconut water (250ml)
Macros:
- Protein: 70g
- Carbs: 100g
- Fats: 15g

Dinner (7:00 PM) - 800 calories:
- Wild-caught salmon (200g)
- Cauliflower rice (150g)
- Asparagus (100g)
- Almond-based sauce (50ml)
Macros:
- Protein: 65g
- Carbs: 80g
- Fats: 35g

BUDGET-FRIENDLY CONSIDERATIONS:
- Use frozen vegetables
- Buy proteins in bulk
- Choose seasonal produce
- Utilize cost-effective protein sources like eggs and chicken

ADDITIONAL RECOMMENDATIONS:
- Hydration: 4-5 liters of water daily
- Supplement with gluten-free, dairy-free multivitamin
- Consider creatine monohydrate for muscle gain

TOTAL DAILY INTAKE:
- Calories: 4,300
- Protein: 340g
- Carbohydrates: 530g
- Fats: 155g

This meal plan is precisely tailored to the athlete's specific height, weight, goal, and dietary restrictions, ensuring optimal nutrition for muscle mass gain and high-performance strength training.
</example_good_output>

The solution provides a highly detailed, scientifically-grounded meal plan that addresses the core requirements of creating a one-day nutritional strategy for an athlete. It systematically breaks down caloric needs, macronutrient distribution, and meal composition while respecting specified dietary constraints. The plan demonstrates technical nutritional expertise and personalization.

	%s
	`,
		test_case.PromptInputs["height"],
		test_case.PromptInputs["weight"],
		test_case.PromptInputs["goal"],
		test_case.PromptInputs["restrictions"],
		strings.Join(test_case.SolutionCriteria, ".\n\t"),
	)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(prompt)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	// conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	// stop_sequences = append(stop_sequences, "```")
	text, err := ai.Chat(conversation, 0.7, stop_sequences, system_prompt_test_cases)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	return prompt, text
}

func (pe *PromptEvaluator) GenerateTestCase(
	task_description string,
	idea Idea,
	prompt_inputs_spec map[string]string,
) (TestCase, error) {

	prompt := generateTestCasePrompt(task_description, idea, prompt_inputs_spec)
	// fmt.Println(prompt)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")
	text, err := ai.Chat(conversation, 0.7, stop_sequences, system_prompt_test_cases)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	var test_case TestCase

	err = json.Unmarshal([]byte(text), &test_case)
	if err != nil {
		return TestCase{}, err
	}

	return test_case, nil
}

func (pe *PromptEvaluator) GenerateUniqueIdeas(
	task_description string,
	prompt_inputs_spec map[string]string,
	output_file string,
	num_cases int,
) ([]Idea, error) {

	prompt := generateIdeasPrompt(task_description, prompt_inputs_spec, num_cases)
	// fmt.Println(prompt)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")

	text, err := ai.Chat(conversation, 1.0, stop_sequences, system_prompt_ideas)
	if err != nil {
		return []Idea{}, err
	}

	err = os.WriteFile(output_file, []byte(text), 0644) // 0644 sets file permissions
	if err != nil {
		return []Idea{}, err
	}
	log.Println("File written successfully.")

	// read from dataset.json

	bytes, err := os.ReadFile(output_file)
	if err != nil {
		return []Idea{}, err
	}

	var ideas []Idea

	err = json.Unmarshal(bytes, &ideas)
	if err != nil {
		return []Idea{}, err
	}

	return ideas, nil
}

func generateIdeasPrompt(
	task_description string,
	prompt_inputs_spec map[string]string,
	num_cases int,
) string {

	prompt_inputs := []string{}

	for k, v := range prompt_inputs_spec {
		prompt_inputs = append(prompt_inputs, fmt.Sprintf(`"%s": "%s"`, k, v))
	}

	prompt := fmt.Sprintf(`
				Generate %d unique, diverse ideas for testing a prompt that accomplishes this task:
        
        <task_description>
        %s
        </task_description>

        The prompt will receive the following inputs
        <prompt_inputs>
        %s
        </prompt_inputs>
        
        Each idea should represent a distinct scenario or example that tests different aspects of the task.
        
        Output Format:
        Provide your response as a structured JSON array where each item is a brief description of the idea.
        
        Example:
        %sjson
        [
            "Testing with technical computer science terminology",
            "Testing with medical research findings",
            "Testing with complex mathematical concepts",
            ...
        ]
        %s
        
        Ensure each idea is:
        - Clearly distinct from the others
        - Relevant to the task description
        - Specific enough to guide generation of a full test case
        - Quick to solve without requiring extensive computation or multi-step processing
        - Solvable with no more than 400 tokens of output

        Remember, only generate %d unique ideas
`, num_cases, task_description, strings.Join(prompt_inputs, "\n\t"), "```", "```", num_cases)

	return prompt
}
func generateTestCasePrompt(
	task_description string,
	idea Idea,
	prompt_inputs_spec map[string]string,
) string {

	allowed_keys := []string{}
	example_prompt_inputs := []string{}

	for k, v := range prompt_inputs_spec {
		allowed_keys = append(allowed_keys, fmt.Sprintf(`"%s"`, k))
		example_prompt_inputs = append(example_prompt_inputs, fmt.Sprintf(`"%s": "EXAMPLE_VALUE", // %s`, k, v))
	}

	prompt := fmt.Sprintf(` 
        Generate a single detailed test case for a prompt evaluation based on:
        
        <task_description>
        %s
        </task_description>
        
        <specific_idea>
        %s
        </specific_idea>
        
        <allowed_input_keys>
        %s
        </allowed_input_keys>
        
        Output Format:
        %s
        %sjson
        {{
            "prompt_inputs": {{
            		%s
            }},
            "solution_criteria": ["criterion 1", "criterion 2", ...] // Concise list of criteria for evaluating the solution, 1 to 4 items
        }}
        %s
        
        IMPORTANT REQUIREMENTS:
        - You MUST ONLY use these exact input keys in your prompt_inputs: {allowed_keys}        
        - Do NOT add any additional keys to prompt_inputs
        - All keys listed in allowed_input_keys must be included in your response
        - Make the test case realistic and practically useful
        - Include measurable, concise solution criteria
        - The solution criteria should ONLY address the direct requirements of the task description and the generated prompt_inputs
        - Avoid over-specifying criteria with requirements that go beyond the core task
        - Keep solution criteria simple, focused, and directly tied to the fundamental task
        - The test case should be tailored to the specific idea provided
        - Quick to solve without requiring extensive computation or multi-step processing
        - Solvable with no more than 400 tokens of output
        - DO NOT include any fields beyond those specified in the output format

        Here's an example of a sample input with an ideal output:
        <sample_input>
        <sample_task_description>
        Extract topics out of a passage of text
        </sample_task_description>
        <sample_specific_idea>
        Testing with a text that contains multiple nested topics and subtopics (e.g., a passage about renewable energy that covers solar power economics, wind turbine technology, and policy implications simultaneously)
        </sample_specific_idea>

        <sample_allowed_input_keys>
        "content"
        </sample_allowed_input_keys>
        </sample_input>
        <ideal_output>
        %sjson
        {
            "prompt_inputs": {
                "content": "The transition to renewable energy encompasses numerous interdependent dimensions. Solar photovoltaic technology has seen dramatic cost reductions, with panel efficiency improving 24 percent since 2010 while manufacturing costs declined by 89 percent, making it economically competitive with fossil fuels in many markets. Concurrently, wind energy has evolved through innovative turbine designs featuring carbon-fiber composite blades and advanced control systems that increase energy capture by 35 percent in low-wind conditions."
            },
            "solution_criteria": [
                "Includes all topics mentioned"   
            ]
        }
        %s
        </ideal_output>
        This is ideal output because the solution criteria is concise and doesn't ask for anything outside of the scope of the task description.
        `, task_description, idea, strings.Join(allowed_keys, ", "), strings.Join(example_prompt_inputs, "\n\t"), "```", "ppp", "```", "```", "```")

	return prompt
}

type Evaluation struct {
	Strengths  []string `json:"strengths"`
	Weaknesses []string `json:"weaknesses"`
	Reasoning  string   `json:"reasoning"`
	Score      float64  `json:"score"`
	Idea
	Output   string
	TestCase TestCase
	Prompt   string
}
type Task struct {
	Task             string            `json:"task"`
	Format           string            `json:"format"`
	SolutionCriteria string            `json:"solution_criteria"`
	PromptInputSpecs map[string]string `json:"prompt_input_specs"`
}

type TestCase struct {
	Task
	PromptInputs     map[string]any `json:"prompt_inputs"`
	SolutionCriteria []string       `json:"solution_criteria"`
}

// gradeByModel function    Calls the model to grade the output
func (pe *PromptEvaluator) GradeOutput(test_case Task, output string, extra_criteria string) (Evaluation, error) {

	// 	prompt_inputs := []string{}
	// 	example_prompt_inputs := []string{}
	// 	solution_criteria := test_case.SolutionCriteria

	eval_prompt := fmt.Sprintf(`
        Your task is to evaluate the following AI-generated solution with EXTREME RIGOR.

        Original task description:
        <task_description>
        %s
        </task_description>

        Original task inputs:
        <task_inputs>
        %s
        </task_inputs>

        Solution to Evaluate:
        <solution>
        %s
        </solution>

        Criteria you should use to evaluate the solution:
        <criteria>
        %s
        </criteria>

        %s

        Scoring Guidelines:
        * Score 1-3: Solution fails to meet one or more MANDATORY requirements
        * Score 4-6: Solution meets all mandatory requirements but has significant deficiencies in secondary criteria
        * Score 7-8: Solution meets all mandatory requirements and most secondary criteria, with minor issues
        * Score 9-10: Solution meets all mandatory and secondary criteria

        IMPORTANT SCORING INSTRUCTIONS:
        * Grade the output based ONLY on the listed criteria. Do not add your own extra requirements.
        * If a solution meets all of the mandatory and secondary criteria give it a 10
        * Don't complain that the solution "only" meets the mandatory and secondary criteria. Solutions shouldn't go above and beyond - they should meet the exact listed criteria.
        * ANY violation of a mandatory requirement MUST result in a score of 3 or lower
        * The full 1-10 scale should be utilized - don't hesitate to give low scores when warranted

        Output Format
        Provide your evaluation as a structured JSON object with the following fields, in this specific order:
        - "strengths": An array of 1-3 key strengths
        - "weaknesses": An array of 1-3 key areas for improvement
        - "reasoning": A concise explanation of your overall assessment
        - "score": A number between 1-10

        Respond with %s. Keep your response concise and direct.
        Example response shape:
        {{
            "strengths": string[],
            "weaknesses": string[],
            "reasoning": string,
            "score": number
        }
`, test_case.Task, test_case.PromptInputSpecs, output, test_case.SolutionCriteria, extra_criteria, test_case.Format)

	var conversation []anthropic.MessageParam
	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")

	conversation = ai.Add_user_message(conversation, eval_prompt)
	conversation = ai.Add_assistant_message(conversation, "```json")
	text, err := ai.Chat(conversation, 0.0, stop_sequences)
	if err != nil {
		return Evaluation{}, err
	}

	var evaluation Evaluation

	err = json.Unmarshal([]byte(text), &evaluation)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return Evaluation{}, err
	}

	fmt.Println(text)

	return evaluation, nil
}
