package constant

var DalleSizeRatios = map[string]map[string]float64{
	"dall-e-2": {
		"256x256":   1,
		"512x512":   1.125,
		"1024x1024": 1.25,
	},
	"dall-e-3": {
		"1024x1024": 1,
		"1024x1792": 2,
		"1792x1024": 2,
	},
	"gpt-image-1": {
		"1024x1024": 1,
		"1024x1792": 2,
		"1792x1024": 2,
	},
	"stable-diffusion-xl": {
		"512x1024":  1,
		"1024x768":  1,
		"1024x1024": 1,
		"576x1024":  1,
		"1024x576":  1,
	},
	"wanx-v1": {
		"1024x1024": 1,
		"720x1280":  1,
		"1280x720":  1,
	},
}

var DalleGenerationImageAmounts = map[string][2]int{
	"dall-e-2":            {1, 10},
	"dall-e-3":            {1, 1}, // OpenAI allows n=1 currently.
	"gpt-image-1":         {1, 1}, // OpenAI allows n=1 currently.
	"stable-diffusion-xl": {1, 4}, // Ali
	"wanx-v1":             {1, 4}, // Ali
}

var DalleImagePromptLengthLimitations = map[string]int{
	"dall-e-2":            1000,
	"dall-e-3":            4000,
	"gpt-image-1":         4000,
	"stable-diffusion-xl": 4000,
	"wanx-v1":             4000,
	"cogview-3":           833,
}
