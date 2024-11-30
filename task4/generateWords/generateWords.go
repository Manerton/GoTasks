package generateWords

import (
	"math/rand"
	"strings"
	"time"
)

// Префиксы, корни и суффиксы для генерации осмысленных слов
var prefixes = []string{"pro", "un", "re", "trans", "mis", "dis", "in", "out"}
var roots = []string{"form", "spect", "ject", "duct", "tract", "port", "press", "vert"}
var suffixes = []string{"able", "ible", "ing", "ed", "er", "ion", "ment", "ly", "ity"}

// randomElement выбирает случайный элемент из слайса строк
func randomElement(elements []string) string {
	return elements[rand.Intn(len(elements))]
}

// generateMeaningfulWord генерирует осмысленное слово
func generateMeaningfulWord() string {
	var word strings.Builder

	// Используем префикс, корень и суффикс с некоторой вероятностью
	if rand.Float32() < 0.7 { // Префикс добавляется с вероятностью 70%
		word.WriteString(randomElement(prefixes))
	}
	word.WriteString(randomElement(roots))
	if rand.Float32() < 0.8 { // Суффикс добавляется с вероятностью 80%
		word.WriteString(randomElement(suffixes))
	}

	return word.String()
}

// GenerateWords генерирует список осмысленных слов
func GenerateWords(count int) []string {
	rand.NewSource(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = generateMeaningfulWord()
	}
	return words
}
